package relation

type RelationFollow struct {
	Id       int
	UserId   int `orm:"description(用户id)"`
	FollowId int `orm:"description(用户关注的id)"`
}

func (attention *RelationFollow) TableName() string {
	return "db_relation_follow"
}
