package main

import (
	_ "social_server/filters"
	_ "social_server/models"
	_ "social_server/routers"
	_ "social_server/utils"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}
