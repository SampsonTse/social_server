package user

import (
	userModel "social_server/models/user"

	"github.com/beego/beego/v2/client/orm"
)

// 根据手机号查询用户
func GetUserByPhone(Phone int) (userModel.UserInfo, error) {
	user := userModel.UserInfo{}
	qs := orm.NewOrm().QueryTable("db_user_info")
	err := qs.Filter("phone", Phone).One(&user)
	return user, err

}

// 插入新用户
// 返回插入的id和错误信息
func AddUser(user userModel.UserInfo) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(&user)
	return id, err
}

// 判断账号是否存在重复
func JudgeUserAccountExist(account string) bool {
	o := orm.NewOrm()
	exist := o.QueryTable("db_user_info").Filter("Account", account).Exist()
	return exist
}

// 获取帐号信息
func GetUserInfoByAccount(account string) (userModel.UserInfo, error) {
	o := orm.NewOrm()
	var u userModel.UserInfo
	err := o.QueryTable("db_user_info").Filter("Account", account).One(&u)
	return u, err

}

// 判断手机号是否存在重复
func JudgeUserPhoneExist(phone int) bool {
	o := orm.NewOrm()
	exist := o.QueryTable("db_user_info").Filter("Phone", phone).Exist()
	return exist
}
