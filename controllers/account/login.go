package account

import (
	"encoding/json"
	"social_server/common"
	userModel "social_server/models/user"
	accountService "social_server/services/account"

	"github.com/beego/beego/v2/core/logs"
)

// @router /login [post]
func (ctrl *AccountController) Login() {
	req := userModel.LoginRequest{}
	err := json.Unmarshal(ctrl.Ctx.Input.RequestBody, &req)
	if err != nil {
		resp := common.CommonResponse{
			Code: common.GET_DATA_ERROR,
			Msg:  common.RequestMsg[common.GET_DATA_ERROR],
			Data: "",
		}
		logs.Info("ERROR:", "/account/login get param error:", err)
		ctrl.Data["json"] = resp
		ctrl.ServeJSON()
		return
	}
	res, err := accountService.CheckLogin(req)
	if err != nil {
		logs.Info("ERROR:", "/account/login get param error:", err)
	}

	// JWT TOKEN生成报错
	if res.Success == userModel.GET_TOKEN_ERROR {
		resp := common.CommonResponse{
			Code: common.TOKEN_ERROR,
			Msg:  common.RequestMsg[common.TOKEN_ERROR],
			Data: "",
		}
		logs.Info("ERROR:", "/account/login get param error:", err)
		ctrl.Data["json"] = resp
		ctrl.ServeJSON()
		return
	}
	// 生成session
	err = ctrl.SetSession(res.Token, res.Account)
	if err != nil {
		resp := common.CommonResponse{
			Code: common.SET_SESS_ERROR,
			Msg:  common.RequestMsg[common.SET_SESS_ERROR],
			Data: "",
		}
		logs.Info("ERROR:", "/account/login set session error:", err)
		ctrl.Data["json"] = resp
		ctrl.ServeJSON()
		return
	}

	resp := common.CommonResponse{
		Code: common.RES_SUCCESS,
		Msg:  common.RequestMsg[common.RES_SUCCESS],
		Data: res,
	}
	ctrl.Data["json"] = resp
	ctrl.ServeJSON()
}
