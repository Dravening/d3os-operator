package dataservice

var authResource = ReqLimit{
	ReqMem:   "1000Mi",
	ReqCpu:   "10m",
	LimitMem: "1000Mi",
	LimitCpu: "200m",
}

var authCMList = CMDataList{
	authCM,
}

var authCM = CMData{
	Name:            Auth.String(),
	VolumeMountName: "configs",
	MountPath:       "/dataservice/auth/config/application.properties",
	SubPath:         "application.properties",
	Data: map[string]string{
		"application.properties": `
server.port=9081

eureka.instance.prefer-ip-address=true
eureka.client.fetch-registry=true
eureka.client.register-with-eureka=true
eureka.client.registry-fetch-interval-seconds=30
eureka.client.service-url.defaultZone={{eurekaUrl}}/eureka

#sso相关
authentication.callback.url=http://sdh.njdev.datago.vip/auth/callback.htm
#客户端过滤器不过滤的url， 不用动
authentication.skip.uri=/auth/callback.htm
#以下都不要动
authentication.get.access.token.url=https://saic.saicmotor.com/api/v1/oauth/token
authentication.refresh.access.token.url=https://saic.saicmotor.com/api/v1/oauth/token
authentication.get.user.info.v2.url=https://saic.saicmotor.com/api/v1/user/get_info2
authentication.logout.url=https://saic.saicmotor.com/api/v1/user/logout
authentication.login.url=https://saic.saicmotor.com/api/v1/login/login_with_identity


#spring 相关
spring.application.name=auth-service
spring.datasource.driver-class-name=com.mysql.cj.jdbc.Driver
spring.datasource.url=jdbc:{{mysqlUrl}}/sdh_auth?characterEncoding=UTF-8&useUnicode=true&useSSL=false&serverTimezone=Asia/Shanghai&zeroDateTimeBehavior=round
spring.datasource.username={{mysqlUsername}}
spring.datasource.password={{mysqlPassword}}
#rabitmq
spring.rabbitmq.host=${rabbitmq.host}
spring.rabbitmq.port=${rabbitmq.port}
spring.rabbitmq.username=${rabbitmq.username}
spring.rabbitmq.password=${rabbitmq.password}
spring.rabbitmq.virtual-host=${rabbitmq.virtual.host}
feign.httpclient.enabled=true
logging.level.web=info

#tenant.sync.url=${tenant.sync.url}

appName = cosmo-auth
#appId = ${appId}
#appSecret = ${appSecret}
tenant.api.baseUri=${tenant.a} http://c01-gateway8083.c01.shjq-uat-a.sxc.sh
tenant.api.token=${tenant.api.baseUri}/gateway/token?appId=${appId}
tenant.api.add=${tenant.api.baseUri}/am/user/createOrUpdate

#mybatis
mybatis-plus.configuration.map-underscore-to-camel-case=true
mybatis-plus.global-config.db-config.logic-delete-field=deleted
mybatis-plus.global-config.db-config.logic-delete-value=true
mybatis-plus.global-config.db-config.logic-not-delete-value=false
mybatis-plus.type-aliases-package=com.haier.cosmo.auth.beans
#mybatis-plus.configuration.log-impl=org.apache.ibatis.logging.stdout.StdOutImpl

jwt.token.secret=4d13002bd1ac7207f4c673e7c2764c57
jwt.token.expiration=43200000

userCenter.endpoint={{uucUrl}}/service/sys/user/getUserByToken
userCenter.tenantInfo={{uucUrl}}/service/tenant/select
`,
	},
}
