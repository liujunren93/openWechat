package store

// token 系统accessToken
type AccessToken struct {
	Val string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	CreateAt    int64
}

func (a *AccessToken) SetCreatedTime(i int64) {
	a.CreateAt = i
}

func (a *AccessToken) SetVal(s string) {
	a.Val = s
}

func (a *AccessToken) GetCreateTime() int64 {
	return a.CreateAt
}

func (a *AccessToken) GetExpire() int64 {
	return a.ExpiresIn
}

func (a *AccessToken) GetVal() string {
	return a.Val
}
