package dataservice

import "fmt"

const (
	ApiManager    DSName = "api-manager"
	Auth          DSName = "auth"
	DsAdapter     DSName = "ds-adapter"
	EsAdapter     DSName = "es-adapter"
	GatewayMaster DSName = "gateway-master"
	GatewayWeb    DSName = "gateway-web"
	Proxy         DSName = "proxy"
	TrdAdapter    DSName = "trd-adapter"
	Web           DSName = "web"

	Mysql  DSName = "mysql-dataservice"
	UUC    DSName = "uuc"
	Eureka DSName = "eureka"
)

const (
	MysqlUrlTrans      string = "{{mysqlUrl}}"
	MysqlUsernameTrans string = "{{mysqlUsername}}"
	MysqlPasswordTrans string = "{{mysqlPassword}}"
	EurekaUrlTrans     string = "{{eurekaUrl}}"
	UucUrlTrans        string = "{{uucUrl}}"
	NamespaceTrans     string = "{{namespace}}"
)

var (
	MysqlInternalUrl = fmt.Sprintf("mysql://%s.%s.svc:3306", Mysql, NamespaceTrans)
	// UucInternalUrl   string = "https://uuc.ymmii.com/duc"
	EurekaInternalUrl = fmt.Sprintf("http://%s.%s.svc:8761", Eureka, NamespaceTrans)
)

type DSName string

func (name DSName) String() string {
	return string(name)
}

// ReqLimit request & limit
type ReqLimit struct {
	ReqMem   string
	ReqCpu   string
	LimitMem string
	LimitCpu string
}

var ResourceMap = map[DSName]ReqLimit{
	ApiManager:    apiManagerResource,
	Auth:          authResource,
	DsAdapter:     dsAdapterResource,
	EsAdapter:     esAdapterResource,
	GatewayMaster: gatewayMasterResource,
	GatewayWeb:    gatewayWebResource,
	Proxy:         proxyResource,
	TrdAdapter:    trdAdapterResource,
	Web:           webResource,

	Eureka: eurekaResource,
	Mysql:  mysqlResource,
}

type CMData struct {
	Name            string
	VolumeMountName string
	MountPath       string
	SubPath         string
	Data            map[string]string
}

type CMDataList []CMData

var CMDataMap = map[DSName]CMDataList{
	ApiManager:    apiManagerCMList,
	Auth:          authCMList,
	DsAdapter:     dsAdapterCMList,
	EsAdapter:     esAdapterCMList,
	Eureka:        eurekaCMList,
	GatewayMaster: gatewayMasterCMList,
	GatewayWeb:    gatewayWebCMList,
	Proxy:         proxyCMList,
	TrdAdapter:    trdAdapterCMList,
	Web:           webCMList,

	Mysql: mysqlCMList,
}
