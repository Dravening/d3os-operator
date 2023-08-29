package controllers

import (
	"context"
	"fmt"
	"time"

	d3osoperatorv1 "d3os-operator/api/v1"
	"d3os-operator/resource/dataservice"

	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// CheckExistsOrCreateMidBackend 验证或新建MiddlewareBackend
func CheckExistsOrCreateMidBackend(ctx context.Context, r *DataServiceReconciler, middlewareBackend *d3osoperatorv1.MiddlewareBackend, dsInstance *d3osoperatorv1.DataService) error {
	var err error
	switch middlewareBackend.Kind {
	case d3osoperatorv1.Deployment:
		err = createDeploymentIfNotExists(ctx, r, middlewareBackend.Deployment, dsInstance)
	case d3osoperatorv1.StatefulSet:
		err = createStatefulSetIfNotExists(ctx, r, middlewareBackend.StatefulSet, dsInstance)
	case d3osoperatorv1.DaemonSet:
		err = createDaemonSetIfNotExists(ctx, r, middlewareBackend.DaemonSet, dsInstance)
	case d3osoperatorv1.ThirdParty:
		// 要查询checkCMD
		if err = cmdCall(middlewareBackend.ThirdParty.CheckCMD); err != nil {
			return fmt.Errorf("unable to connect thirdparty middleware %s, err: %s", middlewareBackend.ThirdParty.Url, err)
		}
	}
	if err != nil {
		return err
	}
	if middlewareBackend.Service != nil {
		err = createServiceIfNotExists(ctx, r, middlewareBackend.Service, dsInstance)
		if err != nil {
			return err
		}
	}
	if middlewareBackend.ConfigMap != nil {
		for _, cm := range middlewareBackend.ConfigMap {
			err = createConfigMapIfNotExists(ctx, r, cm, dsInstance)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func updateStatus(ctx context.Context, r *DataServiceReconciler, dsBackendName dataservice.DSName, dsInstance *d3osoperatorv1.DataService) error {
	// 不太优雅
	time.Sleep(1 * time.Second)
	// 查询dsBackendName的deployment是否是1/1
	objKey := types.NamespacedName{
		Name:      dsBackendName.String(),
		Namespace: dsInstance.Namespace,
	}
	var objTarget client.Object
	updateFlag := false
	phase := d3osoperatorv1.MysqlDone
	message := d3osoperatorv1.MessageNone

	// 分三种情况:
	// 1.中间件, 三方, 直接通过验证
	// 2.中间件, 部署
	// 3.服务
	switch dsBackendName {
	case dataservice.Mysql:
		phase = d3osoperatorv1.MysqlDone
		message = d3osoperatorv1.MessageMysql
		if dsInstance.Spec.Mysql.Kind == d3osoperatorv1.ThirdParty {
			updateFlag = true
		} else {
			objTarget = &appsv1.StatefulSet{}
		}
	case dataservice.UUC:
		phase = d3osoperatorv1.UucDone
		message = d3osoperatorv1.MessageUuc
		if dsInstance.Spec.Uuc.Kind == d3osoperatorv1.ThirdParty {
			updateFlag = true
		} else {
			objTarget = &appsv1.Deployment{}
		}
	case dataservice.Eureka:
		phase = d3osoperatorv1.EurekaDone
		message = d3osoperatorv1.MessageEureka
		if dsInstance.Spec.Eureka.Kind == d3osoperatorv1.ThirdParty {
			updateFlag = true
		} else {
			objTarget = &appsv1.Deployment{}
		}
	case dataservice.ApiManager:
		phase = d3osoperatorv1.ApiManagerDone
		message = d3osoperatorv1.MessageApiManager
		objTarget = &appsv1.Deployment{}
	case dataservice.Auth:
		phase = d3osoperatorv1.AuthDone
		message = d3osoperatorv1.MessageAuth
		objTarget = &appsv1.Deployment{}
	case dataservice.DsAdapter:
		phase = d3osoperatorv1.DsAdapterDone
		message = d3osoperatorv1.MessageDsAdapter
		objTarget = &appsv1.Deployment{}
	case dataservice.EsAdapter:
		phase = d3osoperatorv1.EsAdapterDone
		message = d3osoperatorv1.MessageEsAdapter
		objTarget = &appsv1.Deployment{}
	case dataservice.GatewayMaster:
		phase = d3osoperatorv1.GatewayMasterDone
		message = d3osoperatorv1.MessageGatewayMaster
		objTarget = &appsv1.Deployment{}
	case dataservice.GatewayWeb:
		phase = d3osoperatorv1.GatewayWebDone
		message = d3osoperatorv1.MessageGatewayWeb
		objTarget = &appsv1.Deployment{}
	case dataservice.Proxy:
		phase = d3osoperatorv1.ProxyDone
		message = d3osoperatorv1.MessageProxy
		objTarget = &appsv1.Deployment{}
	case dataservice.TrdAdapter:
		phase = d3osoperatorv1.TrdAdapterDone
		message = d3osoperatorv1.MessageTrdAdapter
		objTarget = &appsv1.Deployment{}
	case dataservice.Web:
		phase = d3osoperatorv1.WebDone
		message = d3osoperatorv1.MessageWeb
		objTarget = &appsv1.Deployment{}
	}
	f := func() error {
		// 直接更新dsInstance的状态
		dsInstance.Status.Phase = phase
		dsInstance.Status.Message = message
		if err := r.Update(ctx, dsInstance); err != nil {
			return err
		}
		return nil
	}
	if updateFlag {
		if err := f(); err != nil {
			return err
		}
		return nil
	}

	err := r.Get(ctx, objKey, objTarget)
	if err != nil {
		if !errors.IsNotFound(err) {
			return err
		}
		// 如果是没找到，也直接返回失败，等待下周期
		return err
	}
	if dsBackendName == dataservice.Mysql {
		newStateful := objTarget.(*appsv1.StatefulSet)
		if newStateful.Status.ReadyReplicas >= 1 {
			updateFlag = true
		}
	} else {
		newDeploy := objTarget.(*appsv1.Deployment)
		if newDeploy.Status.AvailableReplicas >= 1 {
			updateFlag = true
		}
	}
	if updateFlag {
		if err = f(); err != nil {
			return err
		}
		return nil
	}
	return nil
}

// CheckExistsOrCreateSvcBackend 验证或新建DeployWithSvcCM, deployment和service都是必填项, 如果nil直接报错
func CheckExistsOrCreateSvcBackend(ctx context.Context, r *DataServiceReconciler, svcBackend *d3osoperatorv1.ServiceBackend, dsInstance *d3osoperatorv1.DataService) error {
	var err error
	err = createDeploymentIfNotExists(ctx, r, svcBackend.Deployment, dsInstance)
	if err != nil {
		return err
	}
	err = createServiceIfNotExists(ctx, r, svcBackend.Service, dsInstance)
	if err != nil {
		return err
	}
	if svcBackend.ConfigMap != nil {
		for _, cm := range svcBackend.ConfigMap {
			err = createConfigMapIfNotExists(ctx, r, cm, dsInstance)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
