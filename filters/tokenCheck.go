package filters

import (
	"encoding/json"
	"fmt"
	"social_server/common"
	"social_server/utils"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web/context"
)

type TokenFilterReq struct {
	Token string `json:"token"`
}

// 检查token是否合法
func TokenFilter() func(*context.Context) {
	return func(ctx *context.Context) {
		fmt.Println(ctx.Input.URL())
		if _, ok := TOKEN_CHECK_WHITELIST[ctx.Input.URL()]; ok {
			return
		}

		var req TokenFilterReq
		var err error
		if ctx.Input.IsPost() {
			err = json.Unmarshal(ctx.Input.RequestBody, &req)
		} else if ctx.Input.IsGet() {
			ctx.Input.Bind(&req.Token, "token")
		}
		// 获取数据失败
		if err != nil {
			resp := common.CommonResponse{
				Code: common.GET_DATA_ERROR,
				Msg:  common.RequestMsg[common.GET_DATA_ERROR],
				Data: "",
			}
			logs.Info("ERROR:", "token_filter get param error:", err)
			ctx.JSONResp(resp)
		}

		if !utils.ValidateRegisteredClaims(req.Token) {
			fmt.Println(req.Token)
			resp := common.CommonResponse{
				Code: common.TOKEN_EXPIRE,
				Msg:  common.RequestMsg[common.TOKEN_EXPIRE],
				Data: "",
			}
			ctx.JSONResp(resp)
		}
	}
}
