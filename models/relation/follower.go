package relation

type RelationFollower struct {
	Id         int
	UserId     int `orm:"description(用户id)"`
	FollowerId int `orm:"description(用户粉丝的id)"`
}

func (follower *RelationFollower) TableName() string {
	return "db_relation_follower"
}
