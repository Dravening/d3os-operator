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

##### 1.需求分析

我们的目标是，自动创建并维护一个XX项目，比如dataservice；

我们需要提供给dataservice operator以下内容：

1. mysql是否启用默认部署，如果不使用默认部署，请提供Url、验证方法、用户名、密码

- api-manager信息：image、port
- auth信息：image、port
- ds-adapter信息：image、port
- es-adapter信息：image、port
- gateway-master信息：image、port
- gateway-web信息：image、port
- proxy信息：image、port
- trd-adapter信息：image、port

还要提供中间件信息：

- eureka信息：image、port
- mysql信息：image、port

还要提供前台信息：

- dataservice-web：image、port、Nodeport

在operator获得这些信息内容之后，它应当自动为我创建