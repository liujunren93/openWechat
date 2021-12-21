package signature

import (
	"fmt"
	"github.com/liujunren93/openWechat/internal"
	"github.com/liujunren93/openWechat/utils"
	"strings"
	"time"
)

type Api struct {
	todo *internal.Todo
}

func NewApi(todo *internal.Todo) *Api {
	return &Api{todo: todo}
}

type signature struct {
	Timestamp int64  `json:"timestamp"`
	Url       string `json:"url"`
	Noncestr  string `json:"noncestr"`
	Signature string `json:"signature"`
}

func (a *Api) Signature(uri string) (signature,error) {
	sign := signature{
		Timestamp: time.Now().Local().Unix(),
		Url:       uri,
		Noncestr: utils.RandString(10),
	}
	ticket, err := a.todo.GetTicket()
	if err != nil {
		return sign,err
	}
	urlSlice := strings.Split(uri, "#")

	uri = urlSlice[0]
	bufString := fmt.Sprintf("jsapi_ticket=%s&noncestr=%s&timestamp=%d&url=%s", ticket, sign.Noncestr, sign.Timestamp, uri)

	sign.Signature = utils.Sha1(bufString)
	return sign,nil
}

