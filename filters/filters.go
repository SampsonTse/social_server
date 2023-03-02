package filters

import "github.com/beego/beego/v2/server/web"

// token检查白名单
var TOKEN_CHECK_WHITELIST = map[string]int{
	"/fourfire/account/signup": 1,
	"/fourfire/account/login":  1,
}

func init() {
	web.InsertFilter("*", web.BeforeExec, TokenFilter())
	web.InsertFilter("*", web.BeforeExec, TimeFilter())
}
