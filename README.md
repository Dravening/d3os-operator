# d3os-operator for Kubernetes

### 前言

目前d3os产品均使用helm部署方案进行部署，已经起到了很好的效果；但是部分客户希望可以自行掌握产品部署，这要求产品侧应给出一种更"浅显易懂"的部署方案。经过调研，暂定选用operator。

### 目标说明

以dataservice为例，目前dataservice业务，整体依赖三个中间件（mysql、uuc、eureka），且包含9个微服务。由于微服务间存在依赖关系，故拟定部署的先后顺序如下：

```
mysql->uuc->eureka->api-manager->auth->da-adapter->es-adapter->trd-adapter->gateway-master->gateway-web->proxy->web
```

对于中间件，不同客户有不同部署要求，我们默认mysql使用statefulSet，uuc使用第三方服务，eureka使用deployment，这些中间件配置是可以调整的。

### 技术点梳理

1. 简：剥离用户不关注的配置，而仅保留必要的配置，删繁就简，为用户提供一张单一的可变更的配置文件
2. 流：按顺序部署，前置服务部署完毕后，才开始部署下一服务；如前置服务没有部署完毕，等待10s后再次查看其状态
3. 控：对于中间件部署方案，如果客户有调整，则整个dataservice应当联动调整

### 配置文件实例

```
apiVersion: d3os-product.com.d3os/v1
kind: DataService
metadata:
  name: dataservice-sample
spec:
  # 中间件 
  mysql:
    name: mysql-dataservice
    # 使用的类别, 0123分别为deployment、statefulSet、daemonSet、thirdParty
    kind: 1
    thirdParty:
  uuc:
    name: uuc
    kind: 3
    thirdParty:
      # checkCMD用来检查第三方服务是否存活
      checkCMD: 'curl --output /dev/null --silent --head --fail https://id.cosmoplat.com/duc'
      url: https://id.cosmoplat.com/duc
  eureka:
    name: eureka
    kind: 0
    thirdParty:
  
  # 微服务
  api-manager:
    image: registry-edge.cosmoplat.com/dataservice/api-manager:v1.4
    name: api-manager
    nodePort:
    port: 9085
    replica: 1
    env:
      writer: "draven"
  auth:
    image: registry-edge.cosmoplat.com/dataservice/auth:v1.4
    name: auth
    nodePort:
    port: 9081
    replica: 1
  ds-adapter:
    image: registry-edge.cosmoplat.com/dataservice/ds-adapter:v1.4
    name: ds-adapter
    nodePort:
    port: 9086
    replica: 1
  es-adapter:
    image: registry-edge.cosmoplat.com/dataservice/es-adapter:v1.4
    name: es-adapter
    nodePort:
    port: 9087
    replica: 1
  gateway-master:
    image: registry-edge.cosmoplat.com/dataservice/gateway-master:v1.4
    name: gateway-master
    nodePort:
    port: 9184
    replica: 1
  gateway-web:
    image: registry-edge.cosmoplat.com/dataservice/gateway-web:v1.4
    name: gateway-web
    nodePort:
    port: 9084
    replica: 1
  proxy:
    image: registry-edge.cosmoplat.com/dataservice/proxy:v1.4
    name: proxy
    nodePort:
    port: 9094
    replica: 1
  trd-adapter:
    image: registry-edge.cosmoplat.com/dataservice/trd-adapter:v1.4
    name: trd-adapter
    nodePort:
    port: 9088
    replica: 1
  web:
    image: registry-edge.cosmoplat.com/dataservice/dataservice-web:v1.4
    name: web
    nodePort: 30180
    port: 80
    replica: 1
```

部署效果

```
[root@k8s-master ~]# kubectl get dataservice
NAME                 PHASE   MESSAGE   AGE
dataservice-sample                     3s
[root@k8s-master ~]# kubectl get dataservice
NAME                 PHASE      MESSAGE                                                                                                                 AGE
dataservice-sample   mysql已部署   [mysql]->uuc->eureka->api-manager->auth->da-adapter->es-adapter->trd-adapter->gateway-master->gateway-web->proxy->web   5s
[root@k8s-master ~]# kubectl get dataservice
NAME                 PHASE    MESSAGE                                                                                                                 AGE
dataservice-sample   uuc已部署   mysql->[uuc]->eureka->api-manager->auth->da-adapter->es-adapter->trd-adapter->gateway-master->gateway-web->proxy->web   13s
[root@k8s-master ~]# kubectl get dataservice
NAME                 PHASE       MESSAGE                                                                                                                 AGE
dataservice-sample   eureka已部署   mysql->uuc->[eureka]->api-manager->auth->da-adapter->es-adapter->trd-adapter->gateway-master->gateway-web->proxy->web   19s
[root@k8s-master ~]# kubectl get dataservice
NAME                 PHASE            MESSAGE                                                                                                                 AGE
dataservice-sample   api-manager已部署   mysql->uuc->eureka->[api-manager]->auth->da-adapter->es-adapter->trd-adapter->gateway-master->gateway-web->proxy->web   21s
[root@k8s-master ~]# kubectl get dataservice
NAME                 PHASE     MESSAGE                                                                                                                 AGE
dataservice-sample   auth已部署   mysql->uuc->eureka->api-manager->[auth]->da-adapter->es-adapter->trd-adapter->gateway-master->gateway-web->proxy->web   27s
[root@k8s-master ~]# kubectl get dataservice
NAME                 PHASE           MESSAGE                                                                                                                 AGE
dataservice-sample   da-adapter已部署   mysql->uuc->eureka->api-manager->auth->[da-adapter]->es-adapter->trd-adapter->gateway-master->gateway-web->proxy->web   32s
[root@k8s-master ~]# kubectl get dataservice
NAME                 PHASE           MESSAGE                                                                                                                 AGE
dataservice-sample   es-adapter已部署   mysql->uuc->eureka->api-manager->auth->da-adapter->[es-adapter]->trd-adapter->gateway-master->gateway-web->proxy->web   38s
[root@k8s-master ~]# kubectl get dataservice
NAME                 PHASE            MESSAGE                                                                                                                 AGE
dataservice-sample   trd-adapter已部署   mysql->uuc->eureka->api-manager->auth->da-adapter->es-adapter->[trd-adapter]->gateway-master->gateway-web->proxy->web   43s
[root@k8s-master ~]# kubectl get dataservice
NAME                 PHASE               MESSAGE                                                                                                                 AGE
dataservice-sample   gateway-master已部署   mysql->uuc->eureka->api-manager->auth->da-adapter->es-adapter->trd-adapter->[gateway-master]->gateway-web->proxy->web   48s
[root@k8s-master ~]# kubectl get dataservice
NAME                 PHASE            MESSAGE                                                                                                                 AGE
dataservice-sample   gateway-web已部署   mysql->uuc->eureka->api-manager->auth->da-adapter->es-adapter->trd-adapter->gateway-master->[gateway-web]->proxy->web   53s
[root@k8s-master ~]# kubectl get dataservice
NAME                 PHASE      MESSAGE                                                                                                                 AGE
dataservice-sample   proxy已部署   mysql->uuc->eureka->api-manager->auth->da-adapter->es-adapter->trd-adapter->gateway-master->gateway-web->[proxy]->web   58s
[root@k8s-master ~]# kubectl get dataservice
NAME                 PHASE    MESSAGE                                                                                                                 AGE
dataservice-sample   web已部署   mysql->uuc->eureka->api-manager->auth->da-adapter->es-adapter->trd-adapter->gateway-master->gateway-web->proxy->[web]   63s
```

```
[root@k8s-master ~]# kubectl get all
NAME                                  READY   STATUS    RESTARTS   AGE
pod/api-manager-7d955b9657-c2ht7      1/1     Running   0          87m
pod/auth-64b5666f8b-wljzj             1/1     Running   0          87m
pod/ds-adapter-7579b88dd8-w99mk       1/1     Running   0          87m
pod/es-adapter-d6f9889c6-jd6d5        1/1     Running   0          87m
pod/eureka-84bc9df868-jd652           1/1     Running   0          88m
pod/gateway-master-7694c87f87-tljxn   1/1     Running   0          87m
pod/gateway-web-767b6b8bd8-f2vm7      1/1     Running   0          87m
pod/mysql-dataservice-0               1/1     Running   0          88m
pod/proxy-7d9fffd986-rpkjf            1/1     Running   0          87m
pod/trd-adapter-6bb48bfcfd-mxlqc      1/1     Running   0          87m
pod/web-794c4d7f8b-gnx7b              1/1     Running   0          87m

NAME                        TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)        AGE
service/api-manager         ClusterIP   10.96.221.93    <none>        9085/TCP       87m
service/auth                ClusterIP   10.96.12.170    <none>        9081/TCP       87m
service/ds-adapter          ClusterIP   10.96.199.58    <none>        9086/TCP       87m
service/es-adapter          ClusterIP   10.96.129.231   <none>        9087/TCP       87m
service/eureka              ClusterIP   10.96.79.129    <none>        8761/TCP       88m
service/gateway-master      ClusterIP   10.96.214.161   <none>        9184/TCP       87m
service/gateway-web         ClusterIP   10.96.34.85     <none>        9084/TCP       87m
service/kubernetes          ClusterIP   10.96.0.1       <none>        443/TCP        15d
service/mysql-dataservice   ClusterIP   10.96.225.135   <none>        3306/TCP       88m
service/proxy               ClusterIP   10.96.177.174   <none>        9094/TCP       87m
service/trd-adapter         ClusterIP   10.96.74.82     <none>        9088/TCP       87m
service/web                 NodePort    10.96.135.199   <none>        80:30180/TCP   87m

NAME                             READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/api-manager      1/1     1            1           87m
deployment.apps/auth             1/1     1            1           87m
deployment.apps/ds-adapter       1/1     1            1           87m
deployment.apps/es-adapter       1/1     1            1           87m
deployment.apps/eureka           1/1     1            1           88m
deployment.apps/gateway-master   1/1     1            1           87m
deployment.apps/gateway-web      1/1     1            1           87m
deployment.apps/proxy            1/1     1            1           87m
deployment.apps/trd-adapter      1/1     1            1           87m
deployment.apps/web              1/1     1            1           87m

NAME                                        DESIRED   CURRENT   READY   AGE
replicaset.apps/api-manager-7d955b9657      1         1         1       87m
replicaset.apps/auth-64b5666f8b             1         1         1       87m
replicaset.apps/ds-adapter-7579b88dd8       1         1         1       87m
replicaset.apps/es-adapter-d6f9889c6        1         1         1       87m
replicaset.apps/eureka-84bc9df868           1         1         1       88m
replicaset.apps/gateway-master-7694c87f87   1         1         1       87m
replicaset.apps/gateway-web-767b6b8bd8      1         1         1       87m
replicaset.apps/proxy-7d9fffd986            1         1         1       87m
replicaset.apps/trd-adapter-6bb48bfcfd      1         1         1       87m
replicaset.apps/web-794c4d7f8b              1         1         1       87m

NAME                                 READY   AGE
statefulset.apps/mysql-dataservice   1/1     88m
```

```
[root@k8s-master ~]# kubectl logs -n d3os-operator-system d3os-operator-controller-manager-b6cb76668-hmgp2 manager
I0907 09:29:17.538559       1 request.go:655] Throttling request took 1.003740997s, request: GET:https://10.96.0.1:443/apis/apiextensions.k8s.io/v1?timeout=32s
2023-09-07T09:29:18.195Z        INFO    controller-runtime.metrics      metrics server is starting to listen    {"addr": "127.0.0.1:8080"}
2023-09-07T09:29:18.196Z        INFO    setup   starting manager
I0907 09:29:18.196610       1 leaderelection.go:243] attempting to acquire leader lease d3os-operator-system/e0a9249d.com.d3os...
2023-09-07T09:29:18.196Z        INFO    controller-runtime.manager      starting metrics server {"path": "/metrics"}
I0907 09:29:18.240466       1 leaderelection.go:253] successfully acquired lease d3os-operator-system/e0a9249d.com.d3os
2023-09-07T09:29:18.240Z        DEBUG   controller-runtime.manager.events       Normal  {"object": {"kind":"ConfigMap","namespace":"d3os-operator-system","name":"e0a9249d.com.d3os","uid":"ba511072-312c-4b7e-b74d-6fb900bdc8fa","apiVersion":"v1","resourceVersion":"23353048"}, "reason": "LeaderElection", "message": "d3os-operator-controller-manager-b6cb76668-hmgp2_5338c8c8-cc33-4069-87c4-ffe5b19565f8 became leader"}
2023-09-07T09:29:18.240Z        DEBUG   controller-runtime.manager.events       Normal  {"object": {"kind":"Lease","namespace":"d3os-operator-system","name":"e0a9249d.com.d3os","uid":"4f4e2be8-ecd7-49d6-a8b7-fa7c7aa4a090","apiVersion":"coordination.k8s.io/v1","resourceVersion":"23353049"}, "reason": "LeaderElection", "message": "d3os-operator-controller-manager-b6cb76668-hmgp2_5338c8c8-cc33-4069-87c4-ffe5b19565f8 became leader"}
2023-09-07T09:29:18.240Z        INFO    controller-runtime.manager.controller.dataservice       Starting EventSource    {"reconciler group": "d3os-product.com.d3os", "reconciler kind": "DataService", "source": "kind source: /, Kind="}
2023-09-07T09:29:18.434Z        INFO    controller-runtime.manager.controller.dataservice       Starting Controller     {"reconciler group": "d3os-product.com.d3os", "reconciler kind": "DataService"}
2023-09-07T09:29:18.434Z        INFO    controller-runtime.manager.controller.dataservice       Starting workers        {"reconciler group": "d3os-product.com.d3os", "reconciler kind": "DataService", "worker count": 1}
2023-09-07T09:32:17.116Z        INFO    controller-runtime.manager.controller.dataservice       start reconcile logic   {"reconciler group": "d3os-product.com.d3os", "reconciler kind": "DataService", "name": "dataservice-sample", "namespace": "default"}
2023-09-07T09:32:17.116Z        DEBUG   controller-runtime.manager.controller.dataservice       got dsInstance, go on   {"reconciler group": "d3os-product.com.d3os", "reconciler kind": "DataService", "name": "dataservice-sample", "namespace": "default"}
2023-09-07T09:32:17.216Z        INFO    controller-runtime.manager.controller.dataservice       statefulSet mysql-dataservice not found, creating...    {"reconciler group": "d3os-product.com.d3os", "reconciler kind": "DataService", "name": "dataservice-sample", "namespace": "default"}
2023-09-07T09:32:17.322Z        INFO    controller-runtime.manager.controller.dataservice       service mysql-dataservice not found, creating...        {"reconciler group": "d3os-product.com.d3os", "reconciler kind": "DataService", "name": "dataservice-sample", "namespace": "default"}
2023-09-07T09:32:17.429Z        INFO    controller-runtime.manager.controller.dataservice       configMap mysql-dataservice not found, creating...      {"reconciler group": "d3os-product.com.d3os", "reconciler kind": "DataService", "name": "dataservice-sample", "namespace": "default"}
2023-09-07T09:32:17.433Z        INFO    controller-runtime.manager.controller.dataservice       configMap mysql-dataservice-init not found, creating... {"reconciler group": "d3os-product.com.d3os", "reconciler kind": "DataService", "name": "dataservice-sample", "namespace": "default"}
2023-09-07T09:32:22.535Z        INFO    controller-runtime.manager.controller.dataservice       即将update dataservice dataservice-sample的状态至 mysql已部署   {"reconciler group": "d3os-product.com.d3os", "reconciler kind": "DataService", "name": "dataservice-sample", "namespace": "default"}
2023-09-07T09:32:30.834Z        INFO    controller-runtime.manager.controller.dataservice       即将update dataservice dataservice-sample的状态至 uuc已部署     {"reconciler group": "d3os-product.com.d3os", "reconciler kind": "DataService", "name": "dataservice-sample", "namespace": "default"}
2023-09-07T09:32:30.946Z        INFO    controller-runtime.manager.controller.dataservice       deployment eureka not found, creating...        {"reconciler group": "d3os-product.com.d3os", "reconciler kind": "DataService", "name": "dataservice-sample", "namespace": "default"}
2023-09-07T09:32:30.950Z        INFO    controller-runtime.manager.controller.dataservice       service eureka not found, creating...   {"reconciler group": "d3os-product.com.d3os", "reconciler kind": "DataService", "name": "dataservice-sample", "namespace": "default"}
2023-09-07T09:32:35.955Z        INFO    controller-runtime.manager.controller.dataservice       即将update dataservice dataservice-sample的状态至 eureka已部署  {"reconciler group": "d3os-product.com.d3os", "reconciler kind": "DataService", "name": "dataservice-sample", "namespace": "default"}
2023-09-07T09:32:35.964Z        INFO    controller-runtime.manager.controller.dataservice       deployment api-manager not found, creating...   {"reconciler group": "d3os-product.com.d3os", "reconciler kind": "DataService", "name": "dataservice-sample", "namespace": "default"}
2023-09-07T09:32:35.969Z        INFO    controller-runtime.manager.controller.dataservice       service api-manager not found, creating...      {"reconciler group": "d3os-product.com.d3os", "reconciler kind": "DataService", "name": "dataservice-sample", "namespace": "default"}
2023-09-07T09:32:35.981Z        INFO    controller-runtime.manager.controller.dataservice       configMap api-manager not found, creating...    {"reconciler group": "d3os-product.com.d3os", "reconciler kind": "DataService", "name": "dataservice-sample", "namespace": "default"}
2023-09-07T09:32:40.987Z        INFO    controller-runtime.manager.controller.dataservice       即将update dataservice dataservice-sample的状态至 api-manager已部署     {"reconciler group": "d3os-product.com.d3os", "reconciler kind": "DataService", "name": "dataservice-sample", "namespace": "default"}
2023-09-07T09:32:40.996Z        INFO    controller-runtime.manager.controller.dataservice       deployment auth not found, creating...  {"reconciler group": "d3os-product.com.d3os", "reconciler kind": "DataService", "name": "dataservice-sample", "namespace": "default"}
2023-09-07T09:32:41.001Z        INFO    controller-runtime.manager.controller.dataservice       service auth not found, creating...     {"reconciler group": "d3os-product.com.d3os", "reconciler kind": "DataService", "name": "dataservice-sample", "namespace": "default"}
2023-09-07T09:32:41.006Z        INFO    controller-runtime.manager.controller.dataservice       configMap auth not found, creating...   {"reconciler group": "d3os-product.com.d3os", "reconciler kind": "DataService", "name": "dataservice-sample", "namespace": "default"}
2023-09-07T09:32:46.020Z        INFO    controller-runtime.manager.controller.dataservice       即将update dataservice dataservice-sample的状态至 auth已部署    {"reconciler group": "d3os-product.com.d3os", "reconciler kind": "DataService", "name": "dataservice-sample", "namespace": "default"}
2023-09-07T09:32:46.030Z        INFO    controller-runtime.manager.controller.dataservice       deployment ds-adapter not found, creating...    {"reconciler group": "d3os-product.com.d3os", "reconciler kind": "DataService", "name": "dataservice-sample", "namespace": "default"}
2023-09-07T09:32:46.033Z        INFO    controller-runtime.manager.controller.dataservice       service ds-adapter not found, creating...       {"reconciler group": "d3os-product.com.d3os", "reconciler kind": "DataService", "name": "dataservice-sample", "namespace": "default"}
2023-09-07T09:32:46.044Z        INFO    controller-runtime.manager.controller.dataservice       configMap ds-adapter not found, creating...     {"reconciler group": "d3os-product.com.d3os", "reconciler kind": "DataService", "name": "dataservice-sample", "namespace": "default"}
2023-09-07T09:32:51.049Z        INFO    controller-runtime.manager.controller.dataservice       即将update dataservice dataservice-sample的状态至 da-adapter已部署      {"reconciler group": "d3os-product.com.d3os", "reconciler kind": "DataService", "name": "dataservice-sample", "namespace": "default"}
2023-09-07T09:32:51.062Z        INFO    controller-runtime.manager.controller.dataservice       deployment es-adapter not found, creating...    {"reconciler group": "d3os-product.com.d3os", "reconciler kind": "DataService", "name": "dataservice-sample", "namespace": "default"}
2023-09-07T09:32:51.068Z        INFO    controller-runtime.manager.controller.dataservice       service es-adapter not found, creating...       {"reconciler group": "d3os-product.com.d3os", "reconciler kind": "DataService", "name": "dataservice-sample", "namespace": "default"}
2023-09-07T09:32:51.076Z        INFO    controller-runtime.manager.controller.dataservice       configMap es-adapter not found, creating...     {"reconciler group": "d3os-product.com.d3os", "reconciler kind": "DataService", "name": "dataservice-sample", "namespace": "default"}
2023-09-07T09:32:56.083Z        INFO    controller-runtime.manager.controller.dataservice       即将update dataservice dataservice-sample的状态至 es-adapter已部署      {"reconciler group": "d3os-product.com.d3os", "reconciler kind": "DataService", "name": "dataservice-sample", "namespace": "default"}
2023-09-07T09:32:56.092Z        INFO    controller-runtime.manager.controller.dataservice       deployment trd-adapter not found, creating...   {"reconciler group": "d3os-product.com.d3os", "reconciler kind": "DataService", "name": "dataservice-sample", "namespace": "default"}
2023-09-07T09:32:56.096Z        INFO    controller-runtime.manager.controller.dataservice       service trd-adapter not found, creating...      {"reconciler group": "d3os-product.com.d3os", "reconciler kind": "DataService", "name": "dataservice-sample", "namespace": "default"}
2023-09-07T09:32:56.105Z        INFO    controller-runtime.manager.controller.dataservice       configMap trd-adapter not found, creating...    {"reconciler group": "d3os-product.com.d3os", "reconciler kind": "DataService", "name": "dataservice-sample", "namespace": "default"}
2023-09-07T09:32:56.109Z        INFO    controller-runtime.manager.controller.dataservice       configMap trd-adapter-py not found, creating... {"reconciler group": "d3os-product.com.d3os", "reconciler kind": "DataService", "name": "dataservice-sample", "namespace": "default"}
2023-09-07T09:33:01.115Z        INFO    controller-runtime.manager.controller.dataservice       即将update dataservice dataservice-sample的状态至 trd-adapter已部署     {"reconciler group": "d3os-product.com.d3os", "reconciler kind": "DataService", "name": "dataservice-sample", "namespace": "default"}
2023-09-07T09:33:01.127Z        INFO    controller-runtime.manager.controller.dataservice       deployment gateway-master not found, creating...        {"reconciler group": "d3os-product.com.d3os", "reconciler kind": "DataService", "name": "dataservice-sample", "namespace": "default"}
2023-09-07T09:33:01.130Z        INFO    controller-runtime.manager.controller.dataservice       service gateway-master not found, creating...   {"reconciler group": "d3os-product.com.d3os", "reconciler kind": "DataService", "name": "dataservice-sample", "namespace": "default"}
2023-09-07T09:33:01.142Z        INFO    controller-runtime.manager.controller.dataservice       configMap gateway-master not found, creating... {"reconciler group": "d3os-product.com.d3os", "reconciler kind": "DataService", "name": "dataservice-sample", "namespace": "default"}
2023-09-07T09:33:06.151Z        INFO    controller-runtime.manager.controller.dataservice       即将update dataservice dataservice-sample的状态至 gateway-master已部署  {"reconciler group": "d3os-product.com.d3os", "reconciler kind": "DataService", "name": "dataservice-sample", "namespace": "default"}
2023-09-07T09:33:06.161Z        INFO    controller-runtime.manager.controller.dataservice       deployment gateway-web not found, creating...   {"reconciler group": "d3os-product.com.d3os", "reconciler kind": "DataService", "name": "dataservice-sample", "namespace": "default"}
2023-09-07T09:33:06.165Z        INFO    controller-runtime.manager.controller.dataservice       service gateway-web not found, creating...      {"reconciler group": "d3os-product.com.d3os", "reconciler kind": "DataService", "name": "dataservice-sample", "namespace": "default"}
2023-09-07T09:33:06.175Z        INFO    controller-runtime.manager.controller.dataservice       configMap gateway-web not found, creating...    {"reconciler group": "d3os-product.com.d3os", "reconciler kind": "DataService", "name": "dataservice-sample", "namespace": "default"}
2023-09-07T09:33:11.181Z        INFO    controller-runtime.manager.controller.dataservice       即将update dataservice dataservice-sample的状态至 gateway-web已部署     {"reconciler group": "d3os-product.com.d3os", "reconciler kind": "DataService", "name": "dataservice-sample", "namespace": "default"}
2023-09-07T09:33:11.192Z        INFO    controller-runtime.manager.controller.dataservice       deployment proxy not found, creating... {"reconciler group": "d3os-product.com.d3os", "reconciler kind": "DataService", "name": "dataservice-sample", "namespace": "default"}
2023-09-07T09:33:11.197Z        INFO    controller-runtime.manager.controller.dataservice       service proxy not found, creating...    {"reconciler group": "d3os-product.com.d3os", "reconciler kind": "DataService", "name": "dataservice-sample", "namespace": "default"}
2023-09-07T09:33:11.203Z        INFO    controller-runtime.manager.controller.dataservice       configMap proxy not found, creating...  {"reconciler group": "d3os-product.com.d3os", "reconciler kind": "DataService", "name": "dataservice-sample", "namespace": "default"}
2023-09-07T09:33:16.211Z        INFO    controller-runtime.manager.controller.dataservice       即将update dataservice dataservice-sample的状态至 proxy已部署   {"reconciler group": "d3os-product.com.d3os", "reconciler kind": "DataService", "name": "dataservice-sample", "namespace": "default"}
2023-09-07T09:33:16.220Z        INFO    controller-runtime.manager.controller.dataservice       deployment web not found, creating...   {"reconciler group": "d3os-product.com.d3os", "reconciler kind": "DataService", "name": "dataservice-sample", "namespace": "default"}
2023-09-07T09:33:16.224Z        INFO    controller-runtime.manager.controller.dataservice       service web not found, creating...      {"reconciler group": "d3os-product.com.d3os", "reconciler kind": "DataService", "name": "dataservice-sample", "namespace": "default"}
2023-09-07T09:33:16.232Z        INFO    controller-runtime.manager.controller.dataservice       configMap web-template not found, creating...   {"reconciler group": "d3os-product.com.d3os", "reconciler kind": "DataService", "name": "dataservice-sample", "namespace": "default"}
2023-09-07T09:33:16.246Z        INFO    controller-runtime.manager.controller.dataservice       configMap web not found, creating...    {"reconciler group": "d3os-product.com.d3os", "reconciler kind": "DataService", "name": "dataservice-sample", "namespace": "default"}
2023-09-07T09:33:21.252Z        INFO    controller-runtime.manager.controller.dataservice       即将update dataservice dataservice-sample的状态至 web已部署     {"reconciler group": "d3os-product.com.d3os", "reconciler kind": "DataService", "name": "dataservice-sample", "namespace": "default"}
2023-09-07T09:33:21.262Z        INFO    controller-runtime.manager.controller.dataservice       start reconcile logic   {"reconciler group": "d3os-product.com.d3os", "reconciler kind": "DataService", "name": "dataservice-sample", "namespace": "default"}
2023-09-07T09:33:21.262Z        DEBUG   controller-runtime.manager.controller.dataservice       got dsInstance, go on   {"reconciler group": "d3os-product.com.d3os", "reconciler kind": "DataService", "name": "dataservice-sample", "namespace": "default"}
2023-09-07T09:33:21.262Z        INFO    controller-runtime.manager.controller.dataservice       即将update statefulSet mysql-dataservice        {"reconciler group": "d3os-product.com.d3os", "reconciler kind": "DataService", "name": "dataservice-sample", "namespace": "default"}
2023-09-07T09:33:21.335Z        INFO    controller-runtime.manager.controller.dataservice       即将update service mysql-dataservice    {"reconciler group": "d3os-product.com.d3os", "reconciler kind": "DataService", "name": "dataservice-sample", "namespace": "default"}
2023-09-07T09:33:32.434Z        INFO    controller-runtime.manager.controller.dataservice       即将update deployment eureka    {"reconciler group": "d3os-product.com.d3os", "reconciler kind": "DataService", "name": "dataservice-sample", "namespace": "default"}
2023-09-07T09:33:32.442Z        INFO    controller-runtime.manager.controller.dataservice       即将update service eureka       {"reconciler group": "d3os-product.com.d3os", "reconciler kind": "DataService", "name": "dataservice-sample", "namespace": "default"}
2023-09-07T09:33:37.448Z        INFO    controller-runtime.manager.controller.dataservice       即将update deployment api-manager       {"reconciler group": "d3os-product.com.d3os", "reconciler kind": "DataService", "name": "dataservice-sample", "namespace": "default"}
2023-09-07T09:33:37.455Z        INFO    controller-runtime.manager.controller.dataservice       即将update service api-manager  {"reconciler group": "d3os-product.com.d3os", "reconciler kind": "DataService", "name": "dataservice-sample", "namespace": "default"}
2023-09-07T09:33:42.459Z        INFO    controller-runtime.manager.controller.dataservice       即将update deployment auth      {"reconciler group": "d3os-product.com.d3os", "reconciler kind": "DataService", "name": "dataservice-sample", "namespace": "default"}
2023-09-07T09:33:42.465Z        INFO    controller-runtime.manager.controller.dataservice       即将update service auth {"reconciler group": "d3os-product.com.d3os", "reconciler kind": "DataService", "name": "dataservice-sample", "namespace": "default"}
2023-09-07T09:33:47.471Z        INFO    controller-runtime.manager.controller.dataservice       即将update deployment ds-adapter        {"reconciler group": "d3os-product.com.d3os", "reconciler kind": "DataService", "name": "dataservice-sample", "namespace": "default"}
2023-09-07T09:33:47.477Z        INFO    controller-runtime.manager.controller.dataservice       即将update service ds-adapter   {"reconciler group": "d3os-product.com.d3os", "reconciler kind": "DataService", "name": "dataservice-sample", "namespace": "default"}
2023-09-07T09:33:52.482Z        INFO    controller-runtime.manager.controller.dataservice       即将update deployment es-adapter        {"reconciler group": "d3os-product.com.d3os", "reconciler kind": "DataService", "name": "dataservice-sample", "namespace": "default"}
2023-09-07T09:33:52.487Z        INFO    controller-runtime.manager.controller.dataservice       即将update service es-adapter   {"reconciler group": "d3os-product.com.d3os", "reconciler kind": "DataService", "name": "dataservice-sample", "namespace": "default"}
2023-09-07T09:33:57.493Z        INFO    controller-runtime.manager.controller.dataservice       即将update deployment trd-adapter       {"reconciler group": "d3os-product.com.d3os", "reconciler kind": "DataService", "name": "dataservice-sample", "namespace": "default"}
2023-09-07T09:33:57.500Z        INFO    controller-runtime.manager.controller.dataservice       即将update service trd-adapter  {"reconciler group": "d3os-product.com.d3os", "reconciler kind": "DataService", "name": "dataservice-sample", "namespace": "default"}
2023-09-07T09:34:02.504Z        INFO    controller-runtime.manager.controller.dataservice       即将update deployment gateway-master    {"reconciler group": "d3os-product.com.d3os", "reconciler kind": "DataService", "name": "dataservice-sample", "namespace": "default"}
2023-09-07T09:34:02.509Z        INFO    controller-runtime.manager.controller.dataservice       即将update service gateway-master       {"reconciler group": "d3os-product.com.d3os", "reconciler kind": "DataService", "name": "dataservice-sample", "namespace": "default"}
2023-09-07T09:34:07.513Z        INFO    controller-runtime.manager.controller.dataservice       即将update deployment gateway-web       {"reconciler group": "d3os-product.com.d3os", "reconciler kind": "DataService", "name": "dataservice-sample", "namespace": "default"}
2023-09-07T09:34:07.519Z        INFO    controller-runtime.manager.controller.dataservice       即将update service gateway-web  {"reconciler group": "d3os-product.com.d3os", "reconciler kind": "DataService", "name": "dataservice-sample", "namespace": "default"}
2023-09-07T09:34:12.524Z        INFO    controller-runtime.manager.controller.dataservice       即将update deployment proxy     {"reconciler group": "d3os-product.com.d3os", "reconciler kind": "DataService", "name": "dataservice-sample", "namespace": "default"}
2023-09-07T09:34:12.530Z        INFO    controller-runtime.manager.controller.dataservice       即将update service proxy        {"reconciler group": "d3os-product.com.d3os", "reconciler kind": "DataService", "name": "dataservice-sample", "namespace": "default"}
2023-09-07T09:34:17.537Z        INFO    controller-runtime.manager.controller.dataservice       即将update deployment web       {"reconciler group": "d3os-product.com.d3os", "reconciler kind": "DataService", "name": "dataservice-sample", "namespace": "default"}
2023-09-07T09:34:17.545Z        INFO    controller-runtime.manager.controller.dataservice       即将update service web  {"reconciler group": "d3os-product.com.d3os", "reconciler kind": "DataService", "name": "dataservice-sample", "namespace": "default"}
2023-09-07T09:34:22.550Z        INFO    controller-runtime.manager.controller.dataservice       start reconcile logic   {"reconciler group": "d3os-product.com.d3os", "reconciler kind": "DataService", "name": "dataservice-sample", "namespace": "default"}
2023-09-07T09:34:22.550Z        DEBUG   controller-runtime.manager.controller.dataservice       got dsInstance, go on   {"reconciler group": "d3os-product.com.d3os", "reconciler kind": "DataService", "name": "dataservice-sample", "namespace": "default"}
2023-09-07T09:35:50.093Z        INFO    controller-runtime.manager.controller.dataservice       start reconcile logic   {"reconciler group": "d3os-product.com.d3os", "reconciler kind": "DataService", "name": "dataservice-sample", "namespace": "default"}
2023-09-07T09:35:50.093Z        DEBUG   controller-runtime.manager.controller.dataservice       got dsInstance, go on   {"reconciler group": "d3os-product.com.d3os", "reconciler kind": "DataService", "name": "dataservice-sample", "namespace": "default"}
2023-09-07T09:39:05.952Z        INFO    controller-runtime.manager.controller.dataservice       start reconcile logic   {"reconciler group": "d3os-product.com.d3os", "reconciler kind": "DataService", "name": "dataservice-sample", "namespace": "default"}
```

