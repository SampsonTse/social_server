package user

type LoginRequest struct {
	Phone    int    `json:"phone"`
	Password string `json:"password"`
}

type LoginRespsonse struct {
	Success int    `json:"success"`
	Msg     string `json:"message"`
	Token   string `json:"token"`
	Account string `json:"account"`
}

type SignUpRequest struct {
	Phone    int    `json:"phone"`
	Password string `json:"password"`
	Account  string `json:"account"`
	Sex      uint8  `json:"sex"`
	UserName string `json:"username"`
}

type SignUpResponse struct {
	Success int    `json:"success"` // 0失败 1成功
	Reason  string `json:"reason"`  // 失败的原因
}

type GetUserInfoRequest struct {
	Account string `json:"Account" form:"account"`
}
