package store

import "time"

type JsApiTicket struct {
	ErrCode   int    `json:"errcode"`
	ErrMsg    string `json:"errmsg"`
	Val       string `json:"ticket"`
	ExpireIn int64  `json:"expire_in"`
	CreateAt  int64  `json:"create_at"`
	Type      string `json:"type"`
}

func (j *JsApiTicket) SetCreatedTime(i int64) {
	j.CreateAt = time.Now().Local().Unix()
}

func (j *JsApiTicket) SetVal(s string) {
	j.Val = s
}

func (j *JsApiTicket) IsExpire() bool {
	if j.ExpireIn < time.Now().Local().Unix() {
		return true
	}
	return false
}

func (j *JsApiTicket) SetExpire(expire int64) {
	j.ExpireIn = time.Now().Local().Unix() + expire
}

func (j *JsApiTicket) GetVal() string {
	return j.Val
}
