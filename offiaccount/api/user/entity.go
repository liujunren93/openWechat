package user

type AccessTokenRes struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	OpenID       string `json:"openid"`
	Scope        string `json:"scope"`
	ErrCode      int    `json:"errcode"`
	ErrMsg       string `json:"errmsg"`
}
type ListRes struct {
	Total int `json:"total"`
	Count int `json:"count"`
	Data  struct {
		OpenID []string `json:"openid"`
	} `json:"data"`
	NextOpenID string `json:"next_openid"`
}

//SnsInfoRes 用户授权获取信息
type SnsInfoRes struct {
	Gender    int8     `json:"sex"`
	OpenID    string   `json:"openid"`
	Nickname  string   `json:"nickname"`
	Province  string   `json:"province"`
	City      string   `json:"city"`
	Country   string   `json:"country"`
	Avatar    string   `json:"headimgurl"`
	Privilege []string `json:"privilege"`
	UnionID   string   `json:"union_id"`
	ErrCode   int      `json:"errcode"`
	ErrMsg    string   `json:"errmsg"`
}
type InfoRes struct {
	Subscribe int    `json:"subscribe"`
	OpenID    string `json:"openid"`
	NickName  string `json:"nickname"`
	UnionID   string `json:"unionid"`
	Avatar    string `json:"headimgurl"`
}
