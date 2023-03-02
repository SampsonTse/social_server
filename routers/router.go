package routers

import (
	accountController "social_server/controllers/account"
	userController "social_server/controllers/user"

	"github.com/beego/beego/v2/server/web"
)

// bee generate routers
func init() {
	ns := web.NewNamespace("/fourfire",
		web.NSNamespace("/account",
			web.NSInclude(
				&accountController.AccountController{},
			),
		),
		web.NSNamespace("/user",
			web.NSInclude(
				&userController.UserController{},
			),
		),
	)
	//注册 namespace
	web.AddNamespace(ns)
}
