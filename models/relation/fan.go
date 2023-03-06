package relation

type RelationFan struct {
	Id     int
	UserId int `orm:"description(用户id)"`
	FanId  int `orm:"description(用户粉丝的id)"`
}

func (follower *RelationFan) TableName() string {
	return "db_relation_fan"
}
