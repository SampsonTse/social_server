package common

type CommonResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"message"`
	Data interface{} `json:"data"`
}

const (
	RES_SUCCESS = 110000 // 成功

	TOKEN_EXPIRE   = 400001 // tokne过期 & token验证失败
	TIME_EXPIRE    = 400002
	USER_NOT_EXIST = 400003

	GET_DATA_ERROR = 500001 // 无法获取json数据
	TOKEN_ERROR    = 500002 // JWT生成失败
	DB_ERROR       = 500003 // 数据库问题
	SET_SESS_ERROR = 500004 // 设置session失败
)

var RequestMsg = map[int]string{
	RES_SUCCESS: "",

	TOKEN_EXPIRE:   "服务器繁忙，请稍后重试(400001)",
	TIME_EXPIRE:    "服务器繁忙，请稍后重试(400002)",
	USER_NOT_EXIST: "用户不存在",

	GET_DATA_ERROR: "无法获取数据(500001)",
	TOKEN_ERROR:    "服务器繁忙，请稍后重试(500002)",
	DB_ERROR:       "服务器繁忙，请稍后重试(500003)",
	SET_SESS_ERROR: "服务器繁忙，请稍后重试(500004)",
}
