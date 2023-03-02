package common

type CommonResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"message"`
	Data interface{} `json:"data"`
}

type CommonResponseV2 struct {
	Code int    `json:"code"`
	Msg  string `json:"message"`
}

const (
	RES_SUCCESS = 110000 // 成功

	TOKEN_EXPIRE = 400001 // 无法获取json数据
	TIME_EXPIRE  = 400002

	GET_DATA_ERROR = 500001 // 无法获取json数据
	TOKEN_ERROR    = 500002 // JWT生成失败
	DB_ERROR       = 500003 // 数据库问题
)

var RequestMsg = map[int]string{
	RES_SUCCESS: "",

	TOKEN_EXPIRE: "服务器繁忙，请稍后重试(400001)",
	TIME_EXPIRE:  "服务器繁忙，请稍后重试(400002)",

	GET_DATA_ERROR: "无法获取数据(500001)",
	TOKEN_ERROR:    "服务器繁忙，请稍后重试(500002)",
	DB_ERROR:       "服务器繁忙，请稍后重试(500003)",
}
