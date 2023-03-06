package relation

import (
	relationData "social_server/data/relation"
	userData "social_server/data/user"
	relationModel "social_server/models/relation"
	userModel "social_server/models/user"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

func FollowUser(data relationModel.FollowRequest) (relationModel.FollowResponse, error) {

	var resp relationModel.FollowResponse

	// 粉丝的用户信息
	fanUserInfo, err := userData.GetUserInfoByAccount(data.Account)
	if err != nil {
		// 不是空数据报错
		if err != orm.ErrNoRows {
			logs.Info("ERROR: services.relation.FollowUser get database user info error:", err, "  account:", data.Account)
			return resp, err
		}
		// 空数据报错,正常返回
		resp.Success = 0
		resp.Reason = userModel.USER_MSG[userModel.USER_NOT_EXIST]
		logs.Info("WARN: services.relation.FollowUser no rows:", data.Account)
		return resp, nil
	}

	// 被关注的
	followUserInfo, err := userData.GetUserInfoByAccount(data.OtherAccount)
	if err != nil {
		// 不是空数据报错
		if err != orm.ErrNoRows {
			logs.Info("ERROR: services.relation.FollowUser get database user info error:", err, "  account:", data.Account)
			return resp, err
		}
		// 空数据报错
		resp.Success = 0
		resp.Reason = userModel.USER_MSG[userModel.FOLLOW_USER_NOT_EXIST]
		logs.Info("WARN: services.relation.FollowUser no rows:", data.Account)
		return resp, nil
	}
	err = relationData.FollowUser(fanUserInfo.Id, followUserInfo.Id)
	if err != nil {
		logs.Info("ERROR: services.relation.FollowUser err:", err)
		return resp, err
	}

	resp.Success = 1
	return resp, nil
}
