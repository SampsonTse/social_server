package user

import (
	userModel "social_server/models/user"
)

type Test struct {
	Msg string `json:"msg"`
}

// @router /getUserInfo [get]
// @description: 获取自己的用户信息
func (ctrl *UserController) GetUserInfo() {
	var req userModel.GetUserInfoRequest
	ctrl.Ctx.Input.Bind()
}
