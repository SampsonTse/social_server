/**
@LastEditDate: 2023/2/28
@description: 账号登录、登出、注册相关接口
**/

package account

import (
	"github.com/beego/beego/v2/server/web"
)

type AccountController struct {
	web.Controller
}

func (ctrl *AccountController) URLMapping() {
	ctrl.Mapping("Login", ctrl.Login)
	ctrl.Mapping("SingUp", ctrl.SingUp)
}
