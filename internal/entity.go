package internal

import "fmt"

type ErrorRes struct {
	ErrorCode int    `json:"errcode"`
	ErrMsg    string `json:"errmsg"`
}
func (e *ErrorRes) Error() string {
	return fmt.Sprintf("errCode:%d,msg:%s", e.ErrorCode, e.ErrMsg)
}


