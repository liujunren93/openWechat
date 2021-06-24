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

func (j *JsApiTicket) GetCreateTime() int64 {
	return j.CreateAt
}

func (j *JsApiTicket) GetExpire() int64 {
	return j.ExpiresIn
}

func (j *JsApiTicket) GetVal() string {
	return j.Val
}
