package utils

import (
	beego "github.com/beego/beego/v2/server/web"
	cors "github.com/beego/beego/v2/server/web/filter/cors"
)

func initCORS() {
	// 解决跨域(CORS)的问题
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins: []string{"http://localhost:9528", "http://120.77.153.12:8000"}, // 允许访问的域名
		//AllowAllOrigins:  true, // 允许所有域名访问（vue中开启允许请求头带上 Cookie 后不能设置为 true
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},                                                                                               // 允许的请求类型
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type", "X-Token", "X-Requested-With"}, // 允许的头部信息
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},                                         // 允许暴露的头信息
		AllowCredentials: true,                                                                                                                                              // 如果设置为 true
	}))
}
