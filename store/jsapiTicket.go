package store

import "time"

type JsApiTicket struct {
	ErrCode   int    `json:"errcode"`
	ErrMsg    string `json:"errmsg"`
	Val       string `json:"ticket"`
	ExpiresIn int64  `json:"expires_in"`
	CreateAt  int64
	Type      string
}

func (j *JsApiTicket) SetCreatedTime(i int64) {
	j.CreateAt = time.Now().Local().Unix()
}

func (j *JsApiTicket) SetVal(s string) {
	j.Val = s
}

func (j *JsApiTicket) IsExpire() bool {
	if j.ExpiresIn < time.Now().Local().Unix() {
		return true
	}
	return false
}

func (j *JsApiTicket) SetExpire(expire int64) {
	j.ExpiresIn = time.Now().Local().Unix() + expire
}

func (j *JsApiTicket) GetVal() string {
	return j.Val
}

