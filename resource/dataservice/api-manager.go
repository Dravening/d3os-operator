package dataservice

var apiManagerResource = ReqLimit{
	ReqMem:   "1000Mi",
	ReqCpu:   "10m",
	LimitMem: "1000Mi",
	LimitCpu: "200m",
}

var apiManagerCMList = CMDataList{
	apiManagerCM,
}

var apiManagerCM = CMData{
	Name:            ApiManager.String(),
	VolumeMountName: "configs",
	MountPath:       "/dataservice/api-manager/config/application.yml",
	SubPath:         "application.yml",
	Data: map[string]string{
		"application.yml": `
server:
	port: 9085
eureka:
	instance:
		prefer-ip-address: true

	client:
		fetch-registry: true
		register-with-eureka: true
		registry-fetch-interval-seconds: 30
		service-url:
			defaultZone: {{eurekaUrl}}/eureka
spring:
	application:
		name: api-manager-service
	datasource:
		hikari:
			minimum-idle: 5
			maximum-pool-size: 15
			idle-timeout: 30000
			auto-commit: true
			max-lifetime: 1800000
			connection-timeout: 30000
			connection-test-query: select 1
			pool-name: 'ManagerDataSourceHikariCP'

manager:
	datasource:
		typeCode: MYSQL
		url: jdbc:{{mysqlUrl}}/cosmo_api_manager?serverTimezone=Asia/Shanghai
		userName: {{mysqlUsername}}
		password: {{mysqlPassword}}
		maximum-pool-size: 30

api-gateway:
	datasource:
		url: jdbc:{{mysqlUrl}}/cosmo_gateway?serverTimezone=Asia/Shanghai&characterEncoding=utf-8
		userName: {{mysqlUsername}}
		password: {{mysqlPassword}}
		driverClass: com.mysql.cj.jdbc.Driver

data-source:
	datasource:
		url: jdbc:{{mysqlUrl}}/data_space?useLegacyDatetimeCode=false&useSSL=false&characterEncoding=UTF-8&serverTimezone=Asia/Shanghai
		userName: {{mysqlUsername}}
		password: {{mysqlPassword}}
		driverClass: com.mysql.cj.jdbc.Driver

mybatis-plus:
	configuration:
		log-impl: org.apache.ibatis.logging.stdout.StdOutImpl

rest-template:
	connection-timeout: 300
	read-timeout: 600

gateway:
	serviceName: GATEWAY-WEB
	routeDefine: routeDefine
	application: application
	appRouteRelation: appRouteRelation
	userName: admin
	password: cosmo
	wipePathService: udp-sc-user-center-service,udp-sc-persona-service,udp-sc-ldosapp-service
api:
	dsServiceName: API-DSADAPTER-SERVICE
	esServiceName: API-ESADAPTER-SERVICE
	redisServiceName: API-CHADAPTER-SERVICE
	thirdPartServiceName: API-TPADAPTER-SERVICE
	scriptServiceName: API-SCADAPTER-SERVICE
	signatureMethod: md5
	signatureVersion: v1

gateway-admin:
	#添加网关RouteDefine
	addRouteDefine: true
	#添加网关应用
	addApplication: true
	#查询网关应用详情
	searchOneApp: true
	#删除网关应用，同时删除相关relation
	deleteApp: true
	#添加 app-route relation
	addRelation: true
	#删除 app-route relation
	delRelation: true
	#删除RouteDefine
	deleteRouteDefine: true
	#API授权
	saveRDAuth: true
	#取消API授权
	deleteRdAuth: true
	#开启路由
	rdEnabled: true
	#禁用路由
	rdDisabled: true
	#路由开关
	routeDefineSwitch: true

common:
	mybatis:
		tenant:
			ignore: swagger_api_pre_import,channel_api_mapping,tree_channel,channel,app_rd_statistics_day,app_rd_statistics_hour,enterprise_data_set_relation,enterprise_user,mail_template,mail_relation,mail_record`,
	},
}
