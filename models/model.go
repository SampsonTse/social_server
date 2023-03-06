/**
@description：所有新建的模型都在这里init
**/
package models

import (
	relationModel "social_server/models/relation"
	userModel "social_server/models/user"

	"github.com/beego/beego/v2/client/orm"
)

func init() {
	orm.RegisterModel(new(userModel.UserInfo))
	orm.RegisterModel(new(relationModel.RelationFollow))
	orm.RegisterModel(new(relationModel.RelationFan))
}
