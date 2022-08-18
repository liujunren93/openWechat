package qrcode

import (
	"encoding/json"
	"net/url"

	"github.com/liujunren93/openWechat/todo/officialAccount"
)

/**
* @Author: liujunren
* @Date: 2022/1/28 14:30
 */
type Api struct {
	todo *officialAccount.Todo
}

func NewApi(todo *officialAccount.Todo) *Api {
	return &Api{todo: todo}
}

func (a *Api) Create(qrcode Qrcode) (CreatQrcodeRes, error) {
	api := "https://api.weixin.qq.com/cgi-bin/qrcode/create"
	var res CreatQrcodeRes
	marshal, err := json.Marshal(qrcode)
	if err != nil {
		return res, err
	}
	err = a.todo.ToDoFuncPost(api, &res, marshal)
	if err != nil {
		return res, err
	}
	res.QrcodeUrl = "https://mp.weixin.qq.com/cgi-bin/showqrcode?ticket=" + url.QueryEscape(res.Ticket)
	return res, nil
}
