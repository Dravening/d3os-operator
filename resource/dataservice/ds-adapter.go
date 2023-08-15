package dataservice

var dsAdapterResource = ReqLimit{
	ReqMem:   "1000Mi",
	ReqCpu:   "10m",
	LimitMem: "1000Mi",
	LimitCpu: "200m",
}

var dsAdapterCMList = CMDataList{
	dsAdapterCM,
}

var dsAdapterCM = CMData{
	Name: DsAdapter.String(),
	Data: map[string]string{
		"application.yml": `
server:
  port: 9086
spring:
  application:
    name: API-DSADAPTER-SERVICE
  datasource:
    url: jdbc:mysql://10.253.0.28:3306/cosmo_api_manager?useUnicode=true&characterEncoding=utf8&serverTimezone=Asia/Shanghai
    username: ht
    password: HT@2023#!ht0704
    driver-class-name: com.mysql.cj.jdbc.Driver

data-source:
  datasource:
    url: jdbc:mysql://10.253.0.28:3306/data_space?useLegacyDatetimeCode=false&useSSL=false&characterEncoding=UTF-8&serverTimezone=Asia/Shanghai
    userName: ht
    password: HT@2023#!ht0704
    driverClass: com.mysql.cj.jdbc.Driver

  cloud:
    loadbalancer:
      retry:
        enabled: true
    gateway:
      discovery:
        locator:
          enabled: true
          lowerCaseServiceId: true
      default-filters:

  jackson:
    default-property-inclusion: NON_NULL
    date-format: yyyy-MM-dd hh:mm:ss
    time-zone: GMT+8

hystrix:
  command:
    default:
      execution:
        isolation:
          thread:
            timeoutInMilliseconds: 60000

eureka:
  instance:
    prefer-ip-address: true
    lease-expiration-duration-in-seconds: 10
    lease-renewal-interval-in-seconds: 5
  client:
    healthcheck:
      enabled: true
    registry-fetch-interval-seconds: 5
    serviceUrl:
      defaultZone: http://eureka.{{ .Release.Namespace }}.svc:8761/eureka
  enabled: true

feign:
  hystrix:
    enabled: true

management:
  health:
    elasticsearch:
      enabled: false
`,
	},
}
