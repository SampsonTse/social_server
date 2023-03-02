package relation

type RelationAttenion struct {
	Id          int
	UserId      int `orm:"description(用户id)"`
	AttentionId int `orm:"description(用户关注的id)"`
}

func (attention *RelationAttenion) TableName() string {
	return "db_relation_attention"
}
