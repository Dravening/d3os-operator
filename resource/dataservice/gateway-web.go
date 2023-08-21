package dataservice

var gatewayWebResource = ReqLimit{
	ReqMem:   "1000Mi",
	ReqCpu:   "10m",
	LimitMem: "1000Mi",
	LimitCpu: "200m",
}

var gatewayWebCMList = CMDataList{
	gatewayWebCM,
}

var gatewayWebCM = CMData{
	Name:            GatewayWeb.String(),
	VolumeMountName: "configs",
	MountPath:       "/dataservice/gateway/web/config/application.yml",
	SubPath:         "application.yml",
	Data: map[string]string{
		"application.yml": `
server:
  port: 9084
spring:
  datasource:
    url: jdbc:{{mysqlUrl}}/cosmo_gateway?characterEncoding=utf-8&useJDBCCompliantTimezoneShift=true&useLegacyDatetimeCode=false&serverTimezone=Asia/Shanghai
    username: {{mysqlUsername}}
    password: {{mysqlPassword}}
    driver-class-name: com.mysql.cj.jdbc.Driver
  application:
    name: gateway-web
  jpa:
    hibernate:
      ddl-auto: ${db.ddl-auto:none}
    database-platform: org.hibernate.dialect.MySQL5InnoDBDialect
    show-sql: ${db.show-sql:true}
  security:
    user:
      name: admin
      password: cosmo
eureka:
  instance:
    prefer-ip-address: true
  client:
    fetch-registry: true
    register-with-eureka: true
    registry-fetch-interval-seconds: 30
    service-url:
      defaultZone: {{eurekaUrl}}/eureka
management:
  endpoints:
    web:
      exposure:
        include: "*"
service:
  exclude:
`,
	},
}
