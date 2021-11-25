package menu

import (
	"encoding/json"
	"github.com/liujunren93/openWechat/internal"
)

type Api struct {
	todo *internal.Todo
}
func NewApi(todo *internal.Todo) *Api {
	return &Api{todo: todo}
}


func (a *Api) List() (*resMenu, error) {
	var res resMenu
	api := "https://api.weixin.qq.com/cgi-bin/get_current_selfmenu_info"
	err := a.todo.ToDoFuncGet(api, &res)
	return &res, err

}

func (a *Api) Create(menus *Menu) (err error) {
	var req = map[string]*Menu{"button": menus}
	api := "https://api.weixin.qq.com/cgi-bin/menu/create"
	marshal, _ := json.Marshal(&req)
	err = a.todo.ToDoFuncPost(api, nil, marshal)
	return err
}
