# d3os-operator研发指南

### 前言

Operator 模式是 Kubernetes 的扩展软件，它利用 **定制资源** 管理应用及其组件。

其主要由两部分组成，分别为crd以及相应的controller；

controller应当监听其cr资源的 **增删改** 过程，并进行相应的操作。

为了简化Operator开发，我们可以选用一些已有的开源工具，kubebuilder就是其中之一；本次就使用此工具进行d3os-operator的研发工作。

### 初始化

由于kubebuilder目前仅支持darwin和linux版本，故我在服务器10.206.73.136上进行项目的初始化工作。

##### 1.确定版本

由于我们的k8s使用1.20.4版本，那么经过调研，确定kubuilder版本为3.1.0，其可以使用go1.16.10版本

|         | go      | k8s    | kubebuilder |
| ------- | ------- | ------ | ----------- |
| version | 1.16.10 | 1.20.4 | 3.1.0       |

```
[root@k8s-master ~]# kubebuilder version
Version: main.version{KubeBuilderVersion:"3.1.0", KubernetesVendor:"1.19.2", GitCommit:"92e0349ca7334a0a8e5e499da4fb077eb524e94a", BuildDate:"2021-05-27T17:54:28Z", GoOs:"linux", GoArch:"amd64"}
```

##### 2.安装golang

在[go语言中文网](https://studygolang.com/dl)下载对应版本的go语言安装包([点击下载go1.16.10.linux-amd64.tar.gz](https://studygolang.com/dl/golang/go1.16.10.linux-amd64.tar.gz))

```
[root@k8s-master ~]# tar -zxvf go1.16.10.linux-amd64.tar.gz
```

```
cp -r go /usr/local/src
```

设置golang环境变量

```
[root@k8s-master ~]# echo 'export PATH=$PATH:/usr/local/src/go/bin' >> /etc/profile
[root@k8s-master ~]# echo 'export GOPATH=/root/draven/go_work' >> /etc/profile
[root@k8s-master ~]# source /etc/profile
```

检查安装情况

```
[root@k8s-master ~]# go version
go version go1.16.10 linux/amd64
```

##### 3.安装kubebuilder

在[github官方地址](https://github.com/kubernetes-sigs/kubebuilder/releases?page=2)查看合适的版本，[点击下载kubebuilder_linux_amd64](https://github.com/kubernetes-sigs/kubebuilder/releases/download/v3.1.0/kubebuilder_linux_amd64)

```
[root@k8s-master ~]# chmod +x kubebuilder_linux_amd64
```

```
[root@k8s-master ~]# cp kubebuilder_linux_amd64 /usr/local/bin/kubebuilder
```

##### 4.执行init操作

```
[root@k8s-master ~]# mkdir -p $GOPATH/src/d3os-operator
[root@k8s-master d3os-operator]# cd $GOPATH/src/d3os
```

执行项目初始化

```
[root@k8s-master d3os-operator]# kubebuilder init --domain com.d3os
```

##### 5.创建API

创建资源相关内容，group/version/kind这三部分可以确定资源的唯一身份

```
[root@k8s-master d3os-operator]# kubebuilder create api --group d3os-product --version v1 --kind DataService
```

##### 6.生成部署文件【选做】

这一步会生成crd的yaml文件和role.yaml文件，调整后还会再次生成的

```
[root@k8s-master d3os-operator]# make install
```

##### 7.上传git

```
git pull http://xxxxx/xxxxx/d3os-operator.git
git add --all
git commit -m "init: kubebuilder 3.1.0"
git push
```

### 研发工作

#### 1.需求分析

我们的目标是，自动创建并维护一个XX项目，比如dataservice；

我们需要提供给dataservice operator以下内容：

中间件信息：

1. mysql是否启用默认部署，如果不使用默认部署，需要提供Url、验证方法、用户名、密码

2. uuc是否启用默认部署，如果不使用默认部署，需要提供Url、验证方法


服务信息：

- api-manager信息：image、port、replica、nodeport（如需要）、env（如需要）
- auth信息：：image、port、replica、nodeport（如需要）、env（如需要）
- ds-adapter信息：：image、port、replica、nodeport（如需要）、env（如需要）
- es-adapter信息：：image、port、replica、nodeport（如需要）、env（如需要）
- eureka信息：：image、port、replica、nodeport（如需要）、env（如需要）
- gateway-master信息：：image、port、replica、nodeport（如需要）、env（如需要）
- gateway-web信息：：image、port、replica、nodeport（如需要）、env（如需要）
- proxy信息：：image、port、replica、nodeport（如需要）、env（如需要）
- trd-adapter信息：：image、port、replica、nodeport（如需要）、env（如需要）

还要提供前台信息：

- dataservice-web：image、port、Nodeport

在operator获得这些信息内容之后，它应当自动为我创建一个dataservice服务。

#### 2.结构体设计

此operator需要根据dataservice crd来创建dataservice服务（包含多个deployment及其相关的service和configmap），则此operator需要维护一个后台数据结构dsBackend，包含相关的所有k8s内置资源。

而dataservice crd与此dsBackend数据结构，存在映射关系。这样就可以实现创建crd后就创建对应服务的功能。

下方是dataservice crd的数据结构

```
// DataServiceSpec defines the desired state of DataService
type DataServiceSpec struct {
	// 中间件 可选连接
	Mysql *Middleware `json:"mysql"`
	Uuc   *Middleware `json:"uuc"`
	// 服务
	ApiManager    *Service `json:"api-manager"`
	Auth          *Service `json:"auth"`
	DsAdapter     *Service `json:"ds-adapter"`
	Eureka        *Service `json:"eureka"`
	EsAdapter     *Service `json:"es-adapter"`
	TrdAdapter    *Service `json:"trd-adapter"`
	GatewayMaster *Service `json:"gateway-master"`
	GatewayWeb    *Service `json:"gateway-web"`
	Proxy         *Service `json:"proxy"`
	// 前台, 一般是nginx
	Web *Service `json:"web"`
}

type Middleware struct {
	Name       dataservice.DSName    `json:"name"`
	Kind       MiddlewareKind        `json:"kind"`
	ThirdParty *MiddlewareThirdParty `json:"thirdParty,omitempty"`
}

type Service struct {
	Name     dataservice.DSName `json:"name"`
	Port     int32              `json:"port"` // todo: 这应该是数组类型
	NodePort int32              `json:"nodePort,omitempty"`
	Image    string             `json:"image"`
	Replica  int32              `json:"replica"`

	Env map[string]string `json:"env,omitempty"`
}
```

下方是DataServiceBackend的数据结构

```
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
```

#### 3.书写逻辑

调谐逻辑如下

```
func (r *DataServiceReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	rLog := log.FromContext(ctx)
	rLog.Info("start reconcile logic")

	// 实例化数据结构
	dsInstance := &d3osproductv1.DataService{}

	// 通过客户端工具查询，查询条件是否有dsInstance存在
	err := r.Get(ctx, req.NamespacedName, dsInstance)
	if err != nil {
		// 如果没有实例，就返回空结果，这样外部就不再立即调用Reconcile方法了
		if errors.IsNotFound(err) {
			rLog.Info("dataservice instance %s.%s not found, maybe removed", req.Name, req.Namespace)
			return reconcile.Result{}, nil
		}
		rLog.Error(err, "error getting dataservice instance %s.%s", req.Name, req.Namespace)
		// 返回错误信息给外部
		return ctrl.Result{}, err
	}

	// 调谐开始，首先要生成新的dataServiceBackend
	// oldDSBackend := r.DsBackend
	r.DsBackend = dsInstance.Spec.NewDSBackend(req)
	// todo: 比较新DsBackend与oldDSBackend

	// 1.查找中间件实例是否存在	Mysql Uuc
	if CheckExistsOrCreateMidBackend(ctx, r, req, r.DsBackend.Mysql) != nil {
		return ctrl.Result{}, err
	}
	if CheckExistsOrCreateMidBackend(ctx, r, req, r.DsBackend.Uuc) != nil {
		return ctrl.Result{}, err
	}

	// 2.调谐服务 ApiManager Auth DsAdapter EsAdapter Eureka TrdAdapter GatewayMaster GatewayWeb Proxy
	// ApiManager
	if CheckExistsOrCreateSvcBackend(ctx, r, req, r.DsBackend.ApiManager) != nil {
		return ctrl.Result{}, err
	}
	// Auth
	if CheckExistsOrCreateSvcBackend(ctx, r, req, r.DsBackend.Auth) != nil {
		return ctrl.Result{}, err
	}
	// DsAdapter
	if CheckExistsOrCreateSvcBackend(ctx, r, req, r.DsBackend.DsAdapter) != nil {
		return ctrl.Result{}, err
	}
	// EsAdapter
	if CheckExistsOrCreateSvcBackend(ctx, r, req, r.DsBackend.EsAdapter) != nil {
		return ctrl.Result{}, err
	}
	// TrdAdapter
	if CheckExistsOrCreateSvcBackend(ctx, r, req, r.DsBackend.TrdAdapter) != nil {
		return ctrl.Result{}, err
	}
	// GatewayMaster
	if CheckExistsOrCreateSvcBackend(ctx, r, req, r.DsBackend.GatewayMaster) != nil {
		return ctrl.Result{}, err
	}
	// GatewayWeb
	if CheckExistsOrCreateSvcBackend(ctx, r, req, r.DsBackend.GatewayWeb) != nil {
		return ctrl.Result{}, err
	}
	// Proxy
	if CheckExistsOrCreateSvcBackend(ctx, r, req, r.DsBackend.Proxy) != nil {
		return ctrl.Result{}, err
	}

	// 3.查找并部署web服务
	if CheckExistsOrCreateSvcBackend(ctx, r, req, r.DsBackend.Web) != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}
```

### 部署调试

```
// 将crd部署到kubernetes
[root@k8s-master d3os-operator]# make install
/root/draven/go_work/src/d3os-operator/bin/controller-gen "crd:trivialVersions=true,preserveUnknownFields=false" rbac:roleName=manager-role webhook paths="./..." output:crd:artifacts:config=config/crd/bases
/root/draven/go_work/src/d3os-operator/bin/kustomize build config/crd | kubectl apply -f -
customresourcedefinition.apiextensions.k8s.io/dataservices.d3os-product.com.d3os created
```

```
// 查询一下
[root@k8s-master d3os-operator]# kubectl api-versions | grep d3os
d3os-product.com.d3os/v1
```

```
// 删除crd
[root@k8s-master d3os-operator]# make uninstall
/root/draven/go_work/src/d3os-operator/bin/controller-gen "crd:trivialVersions=true,preserveUnknownFields=false" rbac:roleName=manager-role webhook paths="./..." output:crd:artifacts:config=config/crd/bases
/root/draven/go_work/src/d3os-operator/bin/kustomize build config/crd | kubectl delete -f -
customresourcedefinition.apiextensions.k8s.io "dataservices.d3os-product.com.d3os" deleted
```

打包并推送镜像

```
make docker-build docker-push
```

部署
```
make deploy
```

