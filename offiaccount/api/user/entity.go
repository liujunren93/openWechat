package user

type ListRes struct {
	Total int `json:"total"`
	Count int `json:"count"`
	Data  struct {
		OpenID []string `json:"openid"`
	} `json:"data"`
	NextOpenID string `json:"next_openid"`
}

type InfoRes struct {
	Subscribe int    `json:"subscribe"`
	OpenID    string `json:"openid"`
	NickName  string `json:"nickname"`
	UnionID   string `json:"unionid"`
	Avatar    string `json:"headimgurl"`
}
