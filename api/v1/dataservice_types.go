/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	"d3os-operator/resource/dataservice"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

const (
	Deployment MiddlewareKind = iota
	StatefulSet
	DaemonSet
	ThirdParty
)

type MiddlewareKind int

// MiddlewareThirdParty 如果中间件采用真机部署，填写此对象
type MiddlewareThirdParty struct {
	Url      string            `json:"url"`      // mysql://127.0.0.1:3306
	CheckCMD string            `json:"checkCMD"` // 使用checkCMD确定是否正常
	Username string            `json:"username,omitempty"`
	Password string            `json:"password,omitempty"`
	Custom   map[string]string `json:"custom,omitempty"`
}

// Uuc 统一登陆中心地址
//type Uuc struct {
//	// backend
//	GetUserUrl   string // https://uuc.ymmii.com/duc/service/sys/user/getUserByToken
//	GetTenantUrl string // https://uuc.ymmii.com/duc/service/tenant/select 这个一般不用
//	// frontend
//	LoginUrl string // https://uuc.ymmii.com/duc/login/
//}

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

// DataServiceSpec defines the desired state of DataService
type DataServiceSpec struct {
	// 中间件 可选连接
	Mysql  *Middleware `json:"mysql"`
	Uuc    *Middleware `json:"uuc"`
	Eureka *Middleware `json:"eureka"`

	// 服务
	ApiManager    *Service `json:"api-manager"`
	Auth          *Service `json:"auth"`
	DsAdapter     *Service `json:"ds-adapter"`
	EsAdapter     *Service `json:"es-adapter"`
	TrdAdapter    *Service `json:"trd-adapter"`
	GatewayMaster *Service `json:"gateway-master"`
	GatewayWeb    *Service `json:"gateway-web"`
	Proxy         *Service `json:"proxy"`
	// 前台, 一般是nginx
	Web *Service `json:"web"`
}

// DataServiceStatus defines the observed state of DataService
type DataServiceStatus struct {
	// +optional
	Conditions []DsBackendCondition `json:"conditions,omitempty"`
	// +optional
	Phase DsPhase `json:"phase,omitempty"`
	// +optional
	Message string `json:"message,omitempty"`
}

type DsBackendCondition struct {
	// Type of order condition.   Type OrderConditionType `json:"type"`
	// Phase of the condition, one of True, False, Unknown.
	Status corev1.ConditionStatus `json:"status"`
	// Last time the condition transitioned from one status to another.
	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty"`
	// The reason for the condition's last transition.
	Reason string `json:"reason,omitempty"`
	// A human-readable message indicating details about the transition.
	Message string `json:"message,omitempty"`
}

const (
	MysqlDone  DsPhase = "mysql已部署"
	UucDone    DsPhase = "uuc已部署"
	EurekaDone DsPhase = "eureka已部署"

	ApiManagerDone    DsPhase = "api-manager已部署"
	AuthDone          DsPhase = "auth已部署"
	DsAdapterDone     DsPhase = "da-adapter已部署"
	EsAdapterDone     DsPhase = "es-adapter已部署"
	TrdAdapterDone    DsPhase = "trd-adapter已部署"
	GatewayMasterDone DsPhase = "gateway-master已部署"
	GatewayWebDone    DsPhase = "gateway-web已部署"
	ProxyDone         DsPhase = "proxy已部署"
	WebDone           DsPhase = "web已部署"

	MessageNone          = "无特别信息"
	MessageMysql         = "[mysql]->uuc->eureka->api-manager->auth->da-adapter->es-adapter->trd-adapter->gateway-master->gateway-web->proxy->web"
	MessageUuc           = "mysql->[uuc]->eureka->api-manager->auth->da-adapter->es-adapter->trd-adapter->gateway-master->gateway-web->proxy->web"
	MessageEureka        = "mysql->uuc->[eureka]->api-manager->auth->da-adapter->es-adapter->trd-adapter->gateway-master->gateway-web->proxy->web"
	MessageApiManager    = "mysql->uuc->eureka->[api-manager]->auth->da-adapter->es-adapter->trd-adapter->gateway-master->gateway-web->proxy->web"
	MessageAuth          = "mysql->uuc->eureka->api-manager->[auth]->da-adapter->es-adapter->trd-adapter->gateway-master->gateway-web->proxy->web"
	MessageDsAdapter     = "mysql->uuc->eureka->api-manager->auth->[da-adapter]->es-adapter->trd-adapter->gateway-master->gateway-web->proxy->web"
	MessageEsAdapter     = "mysql->uuc->eureka->api-manager->auth->da-adapter->[es-adapter]->trd-adapter->gateway-master->gateway-web->proxy->web"
	MessageTrdAdapter    = "mysql->uuc->eureka->api-manager->auth->da-adapter->es-adapter->[trd-adapter]->gateway-master->gateway-web->proxy->web"
	MessageGatewayMaster = "mysql->uuc->eureka->api-manager->auth->da-adapter->es-adapter->trd-adapter->[gateway-master]->gateway-web->proxy->web"
	MessageGatewayWeb    = "mysql->uuc->eureka->api-manager->auth->da-adapter->es-adapter->trd-adapter->gateway-master->[gateway-web]->proxy->web"
	MessageProxy         = "mysql->uuc->eureka->api-manager->auth->da-adapter->es-adapter->trd-adapter->gateway-master->gateway-web->[proxy]->web"
	MessageWeb           = "mysql->uuc->eureka->api-manager->auth->da-adapter->es-adapter->trd-adapter->gateway-master->gateway-web->proxy->[web]"
)

type DsPhase string

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// DataService is the Schema for the dataservices API
type DataService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DataServiceSpec   `json:"spec,omitempty"`
	Status DataServiceStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// DataServiceList contains a list of DataService
type DataServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataService `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DataService{}, &DataServiceList{})
}
