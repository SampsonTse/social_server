package relation

import (
	"context"
	relationModel "social_server/models/relation"
	"social_server/utils"
	"strconv"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

var followConuntKey = "follow_count_"
var fanConuntKey = "fan_count_"

// 判断是否已经关注
func JudgeIfFollowed(fanId int, followId int) bool {

	o := orm.NewOrm()
	qs := o.QueryTable("db_relation_follow")
	return qs.Filter("user_id", fanId).Filter("follow_id", fanId).Exist()
}

// fanId:粉丝帐号
// followId:被关注的帐号
func FollowUser(fanId int, followId int) error {
	follow := relationModel.RelationFollow{
		UserId:   fanId,
		FollowId: followId,
	}
	fan := relationModel.RelationFan{
		UserId: followId,
		FanId:  fanId,
	}

	o := orm.NewOrm()
	// 启动事务，保证原子性
	to, err := o.Begin()
	if err != nil {
		logs.Info("ERROR: Data.relation.FollowUser start transaction failed:", err)
		return err
	}

	// 更新关注数据库和缓存
	_, err = to.Insert(&follow)
	if err != nil {
		logs.Info("ERROR: Data.relation.FollowUser insert db_relation_attention failed:", err)
		err = to.Rollback()
		if err != nil {
			logs.Info("ERROR: Data.relation.FollowUser rollback failed:", err)
		}
		return err
	}

	err = DelFollowCountRd(fanId)
	if err != nil {
		err = to.Rollback()
		if err != nil {
			logs.Info("ERROR: Data.relation.FollowUser rollback failed:", err)
		}
		return err
	}

	// 更新粉丝数据库和缓存
	_, err = to.Insert(&fan)
	if err != nil {
		logs.Info("ERROR: Data.relation.FollowUser insert db_relation_fan failed:", err)
		err = to.Rollback()
		if err != nil {
			logs.Info("ERROR: Data.relation.FollowUser rollback failed:", err)
		}
		return err
	}

	err = DelFanCountRd(followId)
	if err != nil {
		err = to.Rollback()
		if err != nil {
			logs.Info("ERROR: Data.relation.FollowUser rollback failed:", err)
		}
		return err
	}

	// 提交事务
	err = to.Commit()
	if err != nil {
		logs.Info("ERROR: Data.Follow transaction commit failed:", err)
	}

	return nil
}

// 删除关注数的缓存
func DelFollowCountRd(userId int) error {

	key := followConuntKey + strconv.Itoa(userId)

	err := utils.RedisCache.Del(context.Background(), key).Err()
	if err != nil {
		logs.Info("ERROR: Data.relation.DelFollowCountRd transaction commit failed:", err)
	}
	return err
}

// 删除粉丝数的缓存
func DelFanCountRd(userId int) error {
	key := fanConuntKey + strconv.Itoa(userId)

	err := utils.RedisCache.Del(context.Background(), key).Err()
	if err != nil {
		logs.Info("ERROR: Data.relation.DelFanCountRd transaction commit failed:", err)
	}
	return err
}
