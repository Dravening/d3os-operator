package dataservice

var proxyResource = ReqLimit{
	ReqMem:   "1000Mi",
	ReqCpu:   "10m",
	LimitMem: "1000Mi",
	LimitCpu: "200m",
}

var proxyCMList = CMDataList{
	proxyCM,
}

var proxyCM = CMData{
	Name: Proxy.String(),
	Data: map[string]string{
		"application.yml": `
server:
  servlet:
    context-path: /cosmo-api
  tomcat:
    max-http-form-post-size: -1
  max-http-header-size: -1
  port: 9094

spring:
  application:
    name: api-proxy-service
  datasource:
    driver-class-name: com.mysql.cj.jdbc.Driver
    username: ht
    password: HT@2023#!ht0704
    url: jdbc:mysql://10.253.0.28:3306/cosmo_proxy?serverTimezone=Asia/Shanghai

eureka:
  instance:
    prefer-ip-address: true
  client:
    fetch-registry: true
    register-with-eureka: true
    registry-fetch-interval-seconds: 180
    service-url:
      defaultZone: http://eureka.{{ .Release.Namespace }}.svc:8761/eureka

ribbon:
  ReadTimeout: 180000
  ConnectTimeout: 180000

zuul:
  host:
    socket-timeout-millis: 180000
    connect-timeout-millis: 180000
    max-total-connections: 500
  white-urls: /api-auth/login,/**/v2/api-docs,/api-auth/swapToken
  #  prefix: ${zuul.prefix:/sdh-api}
  #  strip-prefix: ${zuul.strip-prefix:true}
  #  ignored-patterns: ${zuul.ignored-patterns:/api-auth/login,/**/v2/api-docs}
  sensitive-headers: Cookie,Set-Cookie
  routes:
    log-service: /api-log/**
    datasource-service: /api-datasource/**
    auth-service: /api-auth/**
    api-manager-service: /api-manager/**
    hermes-service: /api-hermes/**
    mirror-data-service: /api-mirror/**

hystrix:
  command:
    default:
      execution:
        isolation:
          thread:
            timeoutInMilliseconds: 180000

jwt:
  token:
    secret: 4d13002bd1ac7207f4c673e7c2764c57
    expiration: 43200000
common:
  core:
    interceptor:
      login:
        enable: false
      feign:
        enable: false
`,
	},
}
