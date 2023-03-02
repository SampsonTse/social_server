package utils

func init() {

	// 初始化mysql
	initMysql()
	// 初始化Redis
	initRedis()
	// 解决跨域访问
	initCORS()
	// 初始化session
	initSession()
	// 初始化log
	initLogger()
}
