package store

import "time"

// token 系统accessToken
type AccessToken struct {
	Val      string `json:"access_token"`
	ExpireIn int64  `json:"expire_in"` // 过期时间
	CreateAt int64  `json:"create_at"`
	Type     string `json:"type"`
}

func (a *AccessToken) SetVal(s string) {
	a.Val = s
}

func (a *AccessToken) SetExpire(expire int64) {
	a.ExpireIn = time.Now().Local().Unix() + expire

}

func (a *AccessToken) IsExpire() bool {
	if a.ExpireIn < time.Now().Local().Unix() {
		return true
	}
	return false
}

func (a *AccessToken) GetVal() string {
	return a.Val
}
