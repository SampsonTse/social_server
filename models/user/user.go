package user

type UserInfo struct {
	Id         int
	IsDelete   bool
	CreateTime int
	Account    string `orm:"unique;description(用户的唯一标识)"`
	UserName   string `orm:"description(用户名)"`
	Phone      int    `orm:"description(手机号)"`
	Password   string `orm:"description(密码)"`
	Sex        uint8  `orm:"description(性别 0男 1女)"`
}

func (model *UserInfo) TableName() string {
	return "db_user_info"
}
