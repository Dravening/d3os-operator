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

// DataServiceStatus defines the observed state of DataService
type DataServiceStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Mysql bool `json:"mysql"`
	Uuc   bool `json:"uuc"`

	ApiManagerStatus  bool `json:"api-manager"`
	AuthStatus        bool `json:"auth"`
	DsAdapterStatus   bool `json:"ds-adapter"`
	EsAdapterStatus   bool `json:"es-adapter"`
	Eureka            bool `json:"eureka"`
	TrdAdapterStatus  bool `json:"trd-adapter"`
	GatewayMastStatus bool `json:"gateway-master"`
	GatewayWebStatus  bool `json:"gateway-web"`
	ProxyStatus       bool `json:"proxy"`

	Web bool `json:"web"`
}

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
