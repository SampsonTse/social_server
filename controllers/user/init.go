package user

import "github.com/beego/beego/v2/server/web"

type UserController struct {
	web.Controller
}

func (ctrl *UserController) URLMapping() {
	ctrl.Mapping("GetUserInfo", ctrl.GetUserInfo)
}
