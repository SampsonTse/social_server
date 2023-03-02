package filters

import (
	"encoding/json"
	"social_server/common"
	"time"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web/context"
)

type TimeFilterReq struct {
	TTT int64 `json:"ttt"`
}

const WhiteTTT = 6998233482468

// 检查token是否合法
func TimeFilter() func(*context.Context) {
	return func(ctx *context.Context) {
		var req TimeFilterReq
		var err error
		if ctx.Input.IsPost() {
			err = json.Unmarshal(ctx.Input.RequestBody, &req)
		} else if ctx.Input.IsGet() {
			err = ctx.Input.Bind(&req.TTT, "ttt")
		}

		if err != nil {
			resp := common.CommonResponse{
				Code: common.GET_DATA_ERROR,
				Msg:  common.RequestMsg[common.GET_DATA_ERROR],
				Data: "",
			}
			logs.Info("ERROR:", "time_filter get param error:", err)
			ctx.JSONResp(resp)
		}

		ts := time.Now().Unix()
		// 不在0~10秒的请求直接返回
		// 时间戳不合法
		if (ts-req.TTT > 10 || ts-req.TTT < 0) && req.TTT != WhiteTTT {
			resp := common.CommonResponse{
				Code: common.TIME_EXPIRE,
				Msg:  common.RequestMsg[common.TIME_EXPIRE],
				Data: "",
			}
			ctx.JSONResp(resp)
		}

	}
}
