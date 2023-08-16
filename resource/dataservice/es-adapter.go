package dataservice

var esAdapterResource = ReqLimit{
	ReqMem:   "1000Mi",
	ReqCpu:   "10m",
	LimitMem: "1000Mi",
	LimitCpu: "200m",
}

var esAdapterCMList = CMDataList{
	esAdapterCM,
}

var esAdapterCM = CMData{
	Name: EsAdapter.String(),
	Data: map[string]string{
		"application.yml": `
server:
  compression:
    enabled: true
  tomcat:
    uri-encoding: UTF-8
    #最小线程数
    min-spare-threads: 100
    #最大线程数
    max-threads: 500
    #最大链接数
    max-connections: 1000
    #最大等待队列长度
    accept-count: 300
    #链接建立超时时间
    connection-timeout: 120000
  port: 9087

spring:
  application:
    name: API-ESADAPTER-SERVICE
  datasource:
    url: jdbc:{{mysqlUrl}}/cosmo_api_manager?useUnicode=true&characterEncoding=utf8&serverTimezone=Asia/Shanghai
    username: {{mysqlUsername}}
    password: {{mysqlPassword}}
    driver-class-name: com.mysql.cj.jdbc.Driver
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

jetcache:
  statIntervalMinutes: 15       #统一间隔
  areaInCacheName: false        #设置不把areaName作为Key前缀
  hiddenPackages: com.alibaba   #如果@Cached和@CreateCache的name是自动生成的，会默认的将包//名和方法名作为前缀，为了不让name太长，该设置时将制定的包名截掉
  local:
    default:
      type: caffeine            #缓存类型
      limit: 100                #
      keyConvertor: fastjson    #Key转换器的全局变量
      expireAfterWriteInMillis: 100000
    otherArea:
      type: linkedhashmap
      limit: 100
      keyConvertor: none
      expireAfterWriteInMillis: 100000

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
      defaultZone: {{eurekaUrl}}/eureka
  enabled: true

feign:
  hystrix:
    enabled: true

management:
  health:
    elasticsearch:
      enabled: false

  endpoints:
    web:
      exposure:
        include: "*"
`,
	},
}
