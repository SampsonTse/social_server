package user

const (
	// 注册&修改
	SIGN_UP_SUCCESS   = 101
	DUPLICATE_ACCOUNT = 102
	PHONE_EXIST       = 103
	INSERT_DB_ERROR   = 104

	// 登录
	LOGIN_SUCCESS   = 201
	PHONE_NOT_EXIST = 202
	PASS_ERROR      = 203
	GET_TOKEN_ERROR = 204

	// 查询用户
	USER_NOT_EXIST        = 302
	FOLLOW_USER_NOT_EXIST = 303
	ALREADY_FOLLOW        = 304
)

var SING_UP_MSG = map[int]string{
	SIGN_UP_SUCCESS:   "成功",
	DUPLICATE_ACCOUNT: "账号重复",
	INSERT_DB_ERROR:   "插入数据库失败",
	PHONE_EXIST:       "手机号重复",
}

var LOGIN_MSG = map[int]string{
	LOGIN_SUCCESS:   "登录成功",
	PHONE_NOT_EXIST: "该手机号不存在",
	PASS_ERROR:      "密码错误",
	GET_TOKEN_ERROR: "生成TOKEN失败",
}

var USER_MSG = map[int]string{
	USER_NOT_EXIST:        "您的用户不存在",
	FOLLOW_USER_NOT_EXIST: "您关注的账号不存在",
	ALREADY_FOLLOW:        "您已关注该用户",
}
