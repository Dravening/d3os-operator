package dataservice

import "fmt"

var webResource = ReqLimit{
	ReqMem:   "1000Mi",
	ReqCpu:   "10m",
	LimitMem: "1000Mi",
	LimitCpu: "200m",
}

var webCMList = CMDataList{
	webCM1,
	webCM2,
}

var webCM1 = CMData{
	Name: Web.String(),
	Data: map[string]string{
		"app.config.js": `
/**
  * 海尔 应用配置文件  DataService
  */
window.HAIER_APP_CONFIG = {
  // 系统标识 不要改
  APP: 'DataService',
  // 单点登录: 是否开启UUC单点登录 共享Cookie方式 与【UUC_IS_OPEN】配置互斥
  UUC_IS_OPEN_PUBLIC_COOKIE: false,
  // 单点登录: 是否开启UUC单点登录 跳转方式 与【UUC_IS_OPEN_PUBLIC_COOKIE】配置互斥
  UUC_IS_OPEN: true,
  // 单点登录: UUC统一登录地址
  UUC_URL: 'https://uuc.ymmii.com/duc/login/',
  // UI: favicon.ico图标  如果为空字符显示原系统图标；eg：'/client_config/tianneng/favicon.ico'
  UI_FAVICON: '',
  // UI: logo图片地址 如果为空字符显示原系统图标；eg：'/client_config/tianneng/logo.png'
  UI_LOGO_URL: '',
  // UI: logo名称 如果logo是带标题的,那么不用加
  UI_LOGO_TITLE: '',
};
`,
	},
}

var webCM2 = CMData{
	Name: fmt.Sprintf("%s-template", Web.String()),
	Data: map[string]string{
		"default.conf.template": `
server {
  listen       80;
  listen  [::]:80;
  server_name  localhost;

  # 静态资源路径,不用改
  location / {
    root   /usr/share/nginx/html;
    index  index.html index.htm;
    try_files $uri $uri/ /index.html;
  }

  # 路由代理 location不带后缀/,使用alias
  location /project-dataservice {
    alias   /usr/share/nginx/html;
    index  index.html index.htm;
    try_files $uri $uri/ /project-dataservice/index.html;
  }
  # 静态资源代理 location不带后缀/,使用alias
  location /project-dataservice/ui {
    alias   /usr/share/nginx/html;
  }

  # 代理
  location /project-dataservice/base-api1/ {
    proxy_pass http://proxy.$NAMESPACE.svc:9094/;
  }
  # 代理
  location /project-dataservice/base-api2/ {
    proxy_pass http://gateway-master.$NAMESPACE.svc:9184/;
  }

  error_page   500 502 503 504  /50x.html;
  location = /50x.html {
    root  /usr/share/nginx/html;
  }
}
`,
	},
}
