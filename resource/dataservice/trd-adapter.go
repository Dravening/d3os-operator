package dataservice

import "fmt"

var trdAdapterResource = ReqLimit{
	ReqMem:   "1000Mi",
	ReqCpu:   "10m",
	LimitMem: "1000Mi",
	LimitCpu: "200m",
}

var trdAdapterCMList = CMDataList{
	trdAdapterCM1,
	trdAdapterCM2,
}

var trdAdapterCM1 = CMData{
	Name: Proxy.String(),
	Data: map[string]string{
		"application.yml": `
server:
  port: 9088
spring:
  application:
    name: API-TPADAPTER-SERVICE

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
  datasource:
    driver-class-name: com.mysql.cj.jdbc.Driver
    url: jdbc:{{mysqlUrl}}/cosmo_api_manager?useUnicode=true&characterEncoding=utf8&serverTimezone=Asia/Shanghai
    username: {{mysqlUsername}}
    password: {{mysqlPassword}}

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
`,
	},
}

var trdAdapterCM2 = CMData{
	Name: fmt.Sprintf("%s-py", Proxy.String()),
	Data: map[string]string{
		"pythonScript.py": `
#!/usr/bin/python
# -*- coding=utf-8 -*-
import hashlib
import time
import urllib2
import urllib
import json

g_appKey = ''
g_secret = ''
g_service = ''
version = 'v1'
signatureMethod = 'md5'
ProductExtGatewayUrl = 'https://gateway.saicdt.com/services/'
ProductGatewayUrl = 'http://gateway.saicdt.com/services/'
TestGatewayUrl = 'http://gateway.njtest.datago.vip/services/'
DevGatewayUrl = 'http://gateway.njdev.datago.vip/services/'

def getMessage(map):
    sb = ''
    if map is None:
        return ''
    for key in map:
        if sb != '':
            sb = sb + ('&')
        sb = sb + key + '=' + map[key]
    return sb

def md5(str):
    m = hashlib.md5()
    m.update(str)
    pas = m.hexdigest()
    return pas.upper()

def sign(api, message):
    message = message
    if api is None:
        return None
    method = ''
    pos = api.index('/')
    if pos > 0:
        arr = api.split('?')
        if arr[0].startswith('/'):
            method = arr[0][1:pos]
        else:
            method = arr[0]

        if message is None and len(arr) > 1:
            map = {}
            arrParams = arr[1].split('&')

            for s in arrParams:
                arrParam = s.split('=')
                if len(arrParam) > 1:
                    map[arrParam[0]] = arrParam[1]

            message = getMessage(map)
        if message is None:
            message = ''
        timestamp = time.strftime("%Y%m%d%H%M%S", time.localtime())
        str = g_service + method + message + 'appKey=' + g_appKey + 'format=json' + 'timestamp=' + timestamp + 'signatureMethod=' + signatureMethod + 'version=' + version + g_secret
        sign = md5(str)
        headers = {
                'appKey': g_appKey,
                'timestamp': timestamp,
                'signatureMethod': signatureMethod,
                'signature': sign,
                'version': version
        }
        return headers

def getApiUrl(gateway, api):
    return gateway + g_service + '/' + api

def getEnvApiUrl(api, env):
    if env == 'DEV':
        return getApiUrl(DevGatewayUrl, api)

    elif env == 'TEST':
        return getApiUrl(TestGatewayUrl, api)

    elif env == 'PRODUCT_EXT':
        return getApiUrl(ProductExtGatewayUrl, api)

    elif env == 'PRODUCT':
        return getApiUrl(ProductExtGatewayUrl, api)

    else:
        return getApiUrl(DevGatewayUrl, api)

def setGlobal(service, appKey, secret):
    global g_appKey,g_secret,g_service
    g_appKey = appKey
    g_secret = secret
    g_service = service

def doGet(api, appKey, secret, service, env):
    setGlobal(service, appKey, secret)
    signHeaders = sign(api, None)
    print
    signHeaders
    url = getEnvApiUrl(api, env)
    print
    'doGetUrl:' + url
    request = urllib2.Request(url, headers=signHeaders)
    response = ''
    try:
         response = urllib2.urlopen(request).read()
    except urllib2.URLError as e:
        response = e.read()
    print
    'response:' + response
    return response

def doPost(api, body, appKey, secret, service, env):
    setGlobal(service, appKey, secret)
    signHeaders = sign(api, body)
    url = getEnvApiUrl(api, env)
    print 'doPostUrl:' + url
    request = urllib2.Request(url, body, headers=signHeaders)
    response = ''
    try:
         response = urllib2.urlopen(request).read()
    except urllib2.URLError as e:
        response = e.read()
    print
    'response:' + response
    return response
`,
	},
}
