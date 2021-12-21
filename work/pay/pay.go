package pay

import "github.com/liujunren93/openWechat/internal"

type Api struct {
	todo *core.Todo
}
func NewApi(todo *core.Todo) *Api {
	return &Api{todo: todo}
}


func (a *Api) List() (*resMenu, error) {


}
