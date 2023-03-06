/**
@description:用来关注/取消关注的接口
**/

package user

import (
	"encoding/json"
	"social_server/common"
	relationModel "social_server/models/relation"
	relationService "social_server/services/relation"

	"github.com/beego/beego/v2/core/logs"
)

// @router /followUser
func (ctrl *UserController) FollowUser() {

	var req relationModel.FollowRequest

	if err := json.Unmarshal(ctrl.Ctx.Input.RequestBody, &req); err != nil {
		resp := common.CommonResponse{
			Code: common.GET_DATA_ERROR,
			Msg:  common.RequestMsg[common.GET_DATA_ERROR],
			Data: "",
		}
		logs.Info("ERROR:", "/user/followUser get param error:", err)
		ctrl.Data["json"] = resp
		ctrl.ServeJSON()
		return
	}

	res, err := relationService.FollowUser(req)
	if err != nil {
		logs.Info("ERROR:", "/user/followUser error:", err)
		resp := common.CommonResponse{
			Code: common.DB_ERROR,
			Msg:  common.RequestMsg[common.DB_ERROR],
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
