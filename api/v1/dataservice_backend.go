package v1

import (
	"d3os-operator/resource/dataservice"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/utils/pointer"
	ctrl "sigs.k8s.io/controller-runtime"
	"strings"
)

// DataServiceBackend 这是核心的结构对象,维护了各个服务的资源
type DataServiceBackend struct {
	// 中间件 可选连接
	Mysql *MiddlewareBackend `json:"mysql"`
	Uuc   *MiddlewareBackend `json:"uuc"`
	// 服务
	ApiManager    *ServiceBackend `json:"api-manager"`
	Auth          *ServiceBackend `json:"auth"`
	DsAdapter     *ServiceBackend `json:"ds-adapter"`
	EsAdapter     *ServiceBackend `json:"es-adapter"`
	Eureka        *ServiceBackend `json:"eureka"`
	TrdAdapter    *ServiceBackend `json:"trd-adapter"`
	GatewayMaster *ServiceBackend `json:"gateway-master"`
	GatewayWeb    *ServiceBackend `json:"gateway-web"`
	Proxy         *ServiceBackend `json:"proxy"`
	// 前台, 一般是nginx
	Web *ServiceBackend `json:"web"`
}

type ServiceBackend struct {
	Deployment *appsv1.Deployment           `json:"deployment"`
	Service    *corev1.Service              `json:"service"`
	ConfigMap  map[string]*corev1.ConfigMap `json:"configMap,omitempty"` // key是cm-name，如"api-manager-py"
}

type MiddlewareBackend struct {
	Kind        MiddlewareKind      `json:"kind"`
	Deployment  *appsv1.Deployment  `json:"deployment,omitempty"`
	StatefulSet *appsv1.StatefulSet `json:"statefulSet,omitempty"`
	DaemonSet   *appsv1.DaemonSet   `json:"daemonSet,omitempty"`
	ConfigMap   *corev1.ConfigMap   `json:"configMap,omitempty"`
	Service     *corev1.Service     `json:"service"`
	// 如果中间件采用真机部署，请使用thirdParty
	ThirdParty *MiddlewareThirdParty `json:"thirdParty,omitempty"`
}

// TransSvcToBackend 根据Svc生成SvcBackend结构；这里我想过很多取巧方案，最后还是老老实实写struct吧
func (svc *Service) TransSvcToBackend(req ctrl.Request, uuc, mysql *Middleware) *ServiceBackend {
	svcBackend := ServiceBackend{}

	// deployment
	deployment := svc.NewDeployment(req)
	// service
	service := svc.NewService(req)
	// configMap
	cmMap := svc.NewCMMap(req)
	// configMap replace
	CMReplace(req, cmMap, uuc, mysql)

	svcBackend.Deployment = deployment
	svcBackend.Service = service
	svcBackend.ConfigMap = cmMap
	return &svcBackend
}

func (svc *Service) NewDeployment(req ctrl.Request) *appsv1.Deployment {
	var env []corev1.EnvVar
	for k, v := range svc.Env {
		env = append(env, corev1.EnvVar{
			Name:  k,
			Value: v,
		})
	}

	// 通过svc.name获取默认配置
	reqMem := dataservice.ResourceMap[svc.Name].ReqMem
	reqCpu := dataservice.ResourceMap[svc.Name].ReqCpu
	limitMem := dataservice.ResourceMap[svc.Name].LimitMem
	limitCpu := dataservice.ResourceMap[svc.Name].LimitCpu

	volumeMountList := []corev1.VolumeMount{}
	volumeList := []corev1.Volume{}
	for _, cmData := range dataservice.CMDataMap[svc.Name] {
		volumeMountList = append(volumeMountList, corev1.VolumeMount{
			Name:      cmData.VolumeMountName,
			MountPath: cmData.MountPath,
			SubPath:   cmData.SubPath,
		})
		volumeList = append(volumeList, corev1.Volume{
			Name: cmData.VolumeMountName,
			VolumeSource: corev1.VolumeSource{
				ConfigMap: &corev1.ConfigMapVolumeSource{
					LocalObjectReference: corev1.LocalObjectReference{
						Name: cmData.Name,
					},
				},
			},
		})
	}

	// deployment
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: req.Namespace,
			Name:      svc.Name.String(),
			Labels: map[string]string{
				"dataservice": req.Name,
				"app":         svc.Name.String(),
			},
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: pointer.Int32Ptr(svc.Replica),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": svc.Name.String(),
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": svc.Name.String(),
					},
				},
				Spec: corev1.PodSpec{
					Volumes: volumeList,
					Containers: []corev1.Container{
						{
							Name:            svc.Name.String(),
							Image:           svc.Image,
							ImagePullPolicy: "IfNotPresent",
							Ports: []corev1.ContainerPort{
								{
									Name:          "http",
									Protocol:      corev1.ProtocolTCP,
									ContainerPort: svc.Port,
								},
							},
							Env:          env,
							VolumeMounts: volumeMountList,
							Resources: corev1.ResourceRequirements{
								Requests: corev1.ResourceList{
									"cpu":    resource.MustParse(reqCpu),
									"memory": resource.MustParse(reqMem),
								},
								Limits: corev1.ResourceList{
									"cpu":    resource.MustParse(limitCpu),
									"memory": resource.MustParse(limitMem),
								},
							},
						},
					},
				},
			},
		},
	}
	return deployment
}

func (svc *Service) NewService(req ctrl.Request) *corev1.Service {
	// service
	service := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: req.Namespace,
			Name:      svc.Name.String(),
			Labels: map[string]string{
				"dataservice": req.Name,
				"app":         svc.Name.String(),
			},
		},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{
				{
					Name:       "http",
					Protocol:   corev1.ProtocolTCP,
					Port:       svc.Port,
					TargetPort: intstr.FromInt(int(svc.Port)),
				},
			},
			Selector: map[string]string{
				"app": svc.Name.String(),
			},
			Type: corev1.ServiceTypeClusterIP,
		},
	}
	if svc.NodePort != 0 {
		service.Spec.Ports[0].NodePort = svc.NodePort
		service.Spec.Type = corev1.ServiceTypeNodePort
	}
	return service
}

func (svc *Service) NewCMMap(req ctrl.Request) map[string]*corev1.ConfigMap {
	cmMap := map[string]*corev1.ConfigMap{}
	for _, cmData := range dataservice.CMDataMap[svc.Name] {
		cmMap[cmData.Name] = &corev1.ConfigMap{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: req.Namespace,
				Name:      cmData.Name,
				Labels: map[string]string{
					"dataservice": req.Name,
					"app":         svc.Name.String(),
				},
			},
			Data: cmData.Data,
		}
	}
	return cmMap
}

// TransMidToBackend 只会生成自建类型的中间件
func (mid *Middleware) TransMidToBackend(req ctrl.Request) *MiddlewareBackend {
	midBackend := &MiddlewareBackend{}
	midBackend.Kind = mid.Kind
	midBackend.ThirdParty = mid.ThirdParty
	if mid.Kind == ThirdParty {
		return midBackend
	}

	// 通过svc.name获取默认配置
	reqMem := dataservice.ResourceMap[mid.Name].ReqMem
	reqCpu := dataservice.ResourceMap[mid.Name].ReqCpu
	limitMem := dataservice.ResourceMap[mid.Name].LimitMem
	limitCpu := dataservice.ResourceMap[mid.Name].LimitCpu

	switch mid.Name {
	case dataservice.Mysql:
		statefulSet := &appsv1.StatefulSet{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: req.Namespace,
				Name:      mid.Name.String(),
				Labels: map[string]string{
					"dataservice": req.Name,
					"app":         mid.Name.String(),
				},
			},
			Spec: appsv1.StatefulSetSpec{
				Replicas: pointer.Int32Ptr(1),
				Selector: &metav1.LabelSelector{
					MatchLabels: map[string]string{
						"app": mid.Name.String(),
					},
				},
				Template: corev1.PodTemplateSpec{
					ObjectMeta: metav1.ObjectMeta{
						Labels: map[string]string{
							"app": mid.Name.String(),
						},
					},
					Spec: corev1.PodSpec{
						Containers: []corev1.Container{
							{
								Name:            mid.Name.String(),
								Image:           "registry-edge.cosmoplat.com/dataservice/mysql-dataservice:v20230516",
								ImagePullPolicy: "IfNotPresent",
								Ports: []corev1.ContainerPort{
									{
										Name:          "http",
										Protocol:      corev1.ProtocolTCP,
										ContainerPort: 3306,
									},
								},
								Env: []corev1.EnvVar{
									{Name: "TZ", Value: "Asia/Shanghai"},
									{Name: "MYSQL_ROOT_PASSWORD", Value: "d3os@cosmoplat"},
									{Name: "character-set-server", Value: "utf8mb4"},
									{Name: "collation-server", Value: "utf8mb4_general_ci"},
									{Name: "MYSQL_DATABASE", Value: "sdh_auth"},
								},
								Resources: corev1.ResourceRequirements{
									Requests: corev1.ResourceList{
										"cpu":    resource.MustParse(reqCpu),
										"memory": resource.MustParse(reqMem),
									},
									Limits: corev1.ResourceList{
										"cpu":    resource.MustParse(limitCpu),
										"memory": resource.MustParse(limitMem),
									},
								},
							},
						},
					},
				},
			},
		}
		midBackend.StatefulSet = statefulSet

		// service
		service := &corev1.Service{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: req.Namespace,
				Name:      mid.Name.String(),
				Labels: map[string]string{
					"dataservice": req.Name,
					"app":         mid.Name.String(),
				},
			},
			Spec: corev1.ServiceSpec{
				Ports: []corev1.ServicePort{
					{
						Name:     "http",
						Protocol: corev1.ProtocolTCP,
						Port:     3306,
					},
				},
				Selector: map[string]string{
					"app": mid.Name.String(),
				},
				Type: corev1.ServiceTypeClusterIP,
			},
		}
		midBackend.Service = service

	case dataservice.UUC:
		// 目前都是三方, do nothing
	}

	return midBackend
}

func CMReplace(req ctrl.Request, cmMap map[string]*corev1.ConfigMap, uuc, mysql *Middleware) {
	for _, cm := range cmMap {
		// 替换configMap中的uuc、mysql和eureka
		for k, v := range cm.Data {
			replaced := v
			// mysqlUrl
			if mysql.Kind != ThirdParty {
				replaced = strings.Replace(replaced, dataservice.MysqlUrlTrans, dataservice.MysqlInternalUrl, -1)
			} else {
				replaced = strings.Replace(replaced, dataservice.MysqlUrlTrans, mysql.ThirdParty.Url, -1)
			}
			// uucUrl
			if uuc.Kind != ThirdParty {
				// 目前uuc都是thirdParty
			} else {
				replaced = strings.Replace(replaced, dataservice.UucUrlTrans, uuc.ThirdParty.Url, -1)
			}
			// eurekaUrl
			replaced = strings.Replace(replaced, dataservice.EurekaUrlTrans, dataservice.EurekaInternalUrl, -1)

			// 统一替换{{Namespace}}
			replaced = strings.Replace(replaced, dataservice.NamespaceTrans, req.Namespace, -1)
			cm.Data[k] = replaced
		}
	}
}

func (dsSpec *DataServiceSpec) NewDSBackend(req ctrl.Request) *DataServiceBackend {
	dsBackend := &DataServiceBackend{
		Mysql:         dsSpec.Mysql.TransMidToBackend(req),
		Uuc:           dsSpec.Uuc.TransMidToBackend(req),
		ApiManager:    dsSpec.ApiManager.TransSvcToBackend(req, dsSpec.Uuc, dsSpec.Mysql),
		Auth:          dsSpec.Auth.TransSvcToBackend(req, dsSpec.Uuc, dsSpec.Mysql),
		DsAdapter:     dsSpec.DsAdapter.TransSvcToBackend(req, dsSpec.Uuc, dsSpec.Mysql),
		Eureka:        dsSpec.Eureka.TransSvcToBackend(req, dsSpec.Uuc, dsSpec.Mysql),
		EsAdapter:     dsSpec.EsAdapter.TransSvcToBackend(req, dsSpec.Uuc, dsSpec.Mysql),
		TrdAdapter:    dsSpec.TrdAdapter.TransSvcToBackend(req, dsSpec.Uuc, dsSpec.Mysql),
		GatewayMaster: dsSpec.GatewayMaster.TransSvcToBackend(req, dsSpec.Uuc, dsSpec.Mysql),
		GatewayWeb:    dsSpec.GatewayWeb.TransSvcToBackend(req, dsSpec.Uuc, dsSpec.Mysql),
		Proxy:         dsSpec.Proxy.TransSvcToBackend(req, dsSpec.Uuc, dsSpec.Mysql),
		Web:           dsSpec.Web.TransSvcToBackend(req, dsSpec.Uuc, dsSpec.Mysql),
	}
	return dsBackend
}
