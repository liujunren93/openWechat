package menu

import (
	"encoding/json"
	"github.com/liujunren93/openWechat/offiaccount/internal"
)

type Api struct {
	*internal.Todo
}

func (a *Api) List() (*ResMenu, error) {
	var res ResMenu
	api := "https://api.weixin.qq.com/cgi-bin/get_current_selfmenu_info"
	f := internal.ToDoFuncGet(api, &res)

	err := a.Do(f)

	return &res, err

}

func (a *Api) Create(menus *menu) (err error) {
	var req = map[string]*menu{"button": menus}
	api := "https://api.weixin.qq.com/cgi-bin/menu/create"
	marshal, _ := json.Marshal(&req)
	f := internal.ToDoFuncPost(api, nil, marshal)
	err = a.Do(f)
	return err
}
