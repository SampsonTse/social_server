package user

type Test struct {
	Msg string `json:"msg"`
}

// @router /getUserInfo [get]
// @description: 获取自己的用户信息
func (ctrl *UserController) GetUserInfo() {
	var response Test
	response.Msg = "success"
	ctrl.Data["json"] = response
	ctrl.ServeJSON()
}
