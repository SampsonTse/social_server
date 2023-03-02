package filters

import (
	"encoding/json"
	"social_server/common"
	"social_server/utils"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web/context"
)

type TokenFilterReq struct {
	Token   string `json:"token"`
	Account string `json:"account"`
}

// 检查token是否合法
func TokenFilter() func(*context.Context) {
	return func(ctx *context.Context) {
		if _, ok := TOKEN_CHECK_WHITELIST[ctx.Input.URL()]; ok {
			return
		}

		var req TokenFilterReq
		var err error
		if ctx.Input.IsPost() {
			err = json.Unmarshal(ctx.Input.RequestBody, &req)
		} else if ctx.Input.IsGet() {
			ctx.Input.Bind(&req.Token, "token")
			ctx.Input.Bind(&req.Account, "account")
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

		sess, _ := ctx.Session()
		// 检查token是否过期
		if !utils.ValidateRegisteredClaims(req.Token) || !utils.CheckSess(sess, req.Token, req.Account) {
			resp := common.CommonResponse{
				Code: common.TOKEN_EXPIRE,
				Msg:  common.RequestMsg[common.TOKEN_EXPIRE],
				Data: "",
			}
			ctx.JSONResp(resp)
		}
	}
}
