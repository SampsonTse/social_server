package account

import (
	userData "social_server/data/user"
	userModel "social_server/models/user"
	"time"

	"github.com/beego/beego/v2/core/logs"
)

func SignUp(req userModel.SignUpRequest) (userModel.SignUpResponse, error) {

	var res userModel.SignUpResponse

	Phone := req.Phone
	Account := req.Account
	Password := req.Password
	Sex := req.Sex
	UserName := req.UserName

	// 判断是否重复
	if exist := userData.JudgeUserAccountExist(Account); exist {
		res.Success = userModel.DUPLICATE_ACCOUNT
		res.Reason = userModel.SING_UP_MSG[userModel.DUPLICATE_ACCOUNT]
		return res, nil
	}

	if exist := userData.JudgeUserPhoneExist(Phone); exist {
		res.Success = userModel.DUPLICATE_ACCOUNT
		res.Reason = userModel.SING_UP_MSG[userModel.DUPLICATE_ACCOUNT]
		return res, nil
	}

	ts := time.Now().Unix()
	userInfo := userModel.UserInfo{
		Phone:      Phone,
		Account:    Account,
		Password:   Password,
		CreateTime: int(ts),
		Sex:        Sex,
		UserName:   UserName,
	}

	// 插入数据库失败
	_, err := userData.AddUser(userInfo)
	if err != nil {
		logs.Info("ERROR: userData.AddUser, user:", userInfo, " error:", err)
		res.Success = userModel.INSERT_DB_ERROR
		res.Reason = userModel.SING_UP_MSG[userModel.INSERT_DB_ERROR]
		return res, err
	}

	res.Success = userModel.SIGN_UP_SUCCESS
	return res, nil

}
