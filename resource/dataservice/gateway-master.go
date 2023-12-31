package dataservice

var gatewayMasterResource = ReqLimit{
	ReqMem:   "1000Mi",
	ReqCpu:   "10m",
	LimitMem: "1000Mi",
	LimitCpu: "200m",
}

var gatewayMasterCMList = CMDataList{
	gatewayMasterCM,
}

var gatewayMasterCM = CMData{
	Name:            GatewayMaster.String(),
	VolumeMountName: "configs",
	MountPath:       "/dataservice/gateway/master/config/application.yml",
	SubPath:         "application.yml",
	Data: map[string]string{
		"application.yml": `
server:
  port: 9184
zuul:
  sensitiveHeaders: Access-Control-Allow-Origin,Access-Control-Allow-Methods
  prefix: /services
ribbon:
  ReadTimeout: 10000
  ConnectTimeout: 10000
signature-auth:
  enable: true

spring:
  liquibase:
    enabled: false
    change-log: classpath:/db/change-log/change-log-master.yml
  datasource:
    url: jdbc:{{mysqlUrl}}/cosmo_gateway?characterEncoding=utf-8&useJDBCCompliantTimezoneShift=true&useLegacyDatetimeCode=false&serverTimezone=Asia/Shanghai
    username: {{mysqlUsername}}
    password: {{mysqlPassword}}
    driver-class-name: com.mysql.cj.jdbc.Driver
  application:
    name: gateway-master
  jpa:
    hibernate:
      ddl-auto: update
    database-platform: org.hibernate.dialect.MySQL5InnoDBDialect
    show-sql: false
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

statistics:
  enable: true
  period-seconds: 5

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
