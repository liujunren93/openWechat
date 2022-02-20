package qrcode

import (
	"encoding/json"
	"github.com/liujunren93/openWechat/internal"
)

/**
* @Author: liujunren
* @Date: 2022/1/28 14:30
 */
type Api struct {
	todo *internal.Todo
}

func NewApi(todo *internal.Todo) *Api {
	return &Api{todo: todo}
}

func (a *Api) Create(qrcode Qrcode) (CreatQrcodeRes, error) {
	api := "https://api.weixin.qq.com/cgi-bin/qrcode/create"
	var res CreatQrcodeRes
	marshal, err := json.Marshal(qrcode)
	if err != nil {
		return res, err
	}
	return res, a.todo.ToDoFuncPost(api, &res, marshal)
}

