package relation

// 点击关注/取消关注
type FollowRequest struct {
	Account      string `json:"account"`
	OtherAccount string `json:"other_account"`
}

type FollowResponse struct {
	Success int    `json:"success"`
	Reason  string `json:"reason"`
}
