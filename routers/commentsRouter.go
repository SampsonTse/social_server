package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["social_server/controllers/account:AccountController"] = append(beego.GlobalControllerRouter["social_server/controllers/account:AccountController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["social_server/controllers/account:AccountController"] = append(beego.GlobalControllerRouter["social_server/controllers/account:AccountController"],
        beego.ControllerComments{
            Method: "SingUp",
            Router: `/signup`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["social_server/controllers/user:UserController"] = append(beego.GlobalControllerRouter["social_server/controllers/user:UserController"],
        beego.ControllerComments{
            Method: "FollowUser",
            Router: `/followUser`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["social_server/controllers/user:UserController"] = append(beego.GlobalControllerRouter["social_server/controllers/user:UserController"],
        beego.ControllerComments{
            Method: "GetUserInfo",
            Router: `/getUserInfo`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
