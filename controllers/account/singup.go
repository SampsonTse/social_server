package account

import (
	"encoding/json"
	"social_server/common"
	userModel "social_server/models/user"
	accountService "social_server/services/account"

	"github.com/beego/beego/v2/core/logs"
)

// @router /signup [post]
func (ctrl *AccountController) SingUp() {

	req := userModel.SignUpRequest{}
	err := json.Unmarshal(ctrl.Ctx.Input.RequestBody, &req)
	if err != nil {
		resp := common.CommonResponse{
			Code: common.GET_DATA_ERROR,
			Msg:  common.RequestMsg[common.GET_DATA_ERROR],
			Data: "",
		}
		logs.Info("ERROR:", "/account/signup get param error:", err)
		ctrl.Data["json"] = resp
		ctrl.ServeJSON()
		return
	}

	res, err := accountService.SignUp(req)
	if err != nil {
		logs.Info("ERROR:", "/account/signup error:", err)
		resp := common.CommonResponse{
			Code: common.GET_DATA_ERROR,
			Msg:  common.RequestMsg[common.GET_DATA_ERROR],
			Data: "",
		}
		ctrl.Data["json"] = resp
		ctrl.ServeJSON()
		return
	}

	resp := common.CommonResponse{
		Code: common.RES_SUCCESS,
		Data: res,
	}
	ctrl.Data["json"] = resp
	ctrl.ServeJSON()
}
