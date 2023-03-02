package account

import (
	userData "social_server/data/user"
	userModel "social_server/models/user"
	"social_server/utils"

	"github.com/beego/beego/v2/core/logs"
)

func CheckLogin(req userModel.LoginRequest) (userModel.LoginRespsonse, error) {
	Phone := req.Phone
	Password := req.Password

	userInfo, err := userData.GetUserByPhone(Phone)

	// 手机号查不到
	if err != nil {
		logs.Info("warnning: Data.User.GetUserByPhone Error:", err, "; error Number:", Phone)
		res := userModel.LoginRespsonse{
			Success: userModel.PHONE_NOT_EXIST,
			Msg:     userModel.LOGIN_MSG[userModel.PHONE_NOT_EXIST],
		}
		return res, err
	}

	// 密码错误
	if userInfo.Password != Password {
		res := userModel.LoginRespsonse{
			Success: userModel.PASS_ERROR,
			Msg:     userModel.LOGIN_MSG[userModel.PASS_ERROR],
		}
		return res, nil
	}

	// 生成token
	token, err := utils.GenRegisteredClaims(userInfo.Account)
	if err != nil {
		logs.Info("ERROR: utils.GenRegisteredClaims, error:", err)
		res := userModel.LoginRespsonse{
			Success: userModel.GET_TOKEN_ERROR,
			Msg:     userModel.LOGIN_MSG[userModel.GET_TOKEN_ERROR],
		}
		return res, err
	}

	// 登录成功
	res := userModel.LoginRespsonse{
		Success: userModel.LOGIN_SUCCESS,
		Msg:     userModel.LOGIN_MSG[userModel.LOGIN_SUCCESS],
		Account: userInfo.Account,
		Token:   token,
	}

	return res, nil
}
