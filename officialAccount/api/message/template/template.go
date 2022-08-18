package template

import (
	"encoding/json"
	"fmt"

	"github.com/liujunren93/openWechat/todo/officialAccount"
)

type Api struct {
	todo *officialAccount.Todo
}

func NewApi(todo *officialAccount.Todo) *Api {
	return &Api{todo: todo}
}

func (a *Api) SendMsg(req Template) (TemplateRes, error) {
	api := "https://api.weixin.qq.com/cgi-bin/message/template/send"
	var res TemplateRes
	marshal, _ := json.Marshal(&req)
	fmt.Println(string(marshal))
	return res, a.todo.ToDoFuncPost(api, &res, marshal)
}
