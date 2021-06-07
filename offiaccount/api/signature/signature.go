package signature

import (
	"fmt"
	"github.com/liujunren93/openWechat/offiaccount/internal"
	"github.com/liujunren93/openWechat/utils"
	"strings"
	"time"
)

type Api struct {
	*internal.Todo
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
	ticket, err := a.GetTicket()
	if err != nil {
		return sign,err
	}
	urlSlice := strings.Split(uri, "#")

	uri = urlSlice[0]
	bufString := fmt.Sprintf("jsapi_ticket=%s&noncestr=%s&timestamp=%d&url=%s", ticket, sign.Noncestr, sign.Timestamp, uri)
	//urlVal:=url.Values{}
	//urlVal.Add("jsapi_ticket",ticket)
	//urlVal.Add("noncestr",sign.Noncestr)
	//urlVal.Add("timestamp",strconv.FormatInt(sign.Timestamp,10))
	//urlVal.Add("url",uri)
	//fmt.Println(urlVal.Encode())
	//fmt.Println(bufString)
	sign.Signature = utils.Sha1(bufString)
	return sign,nil
}

