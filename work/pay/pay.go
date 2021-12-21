package pay

import "github.com/liujunren93/openWechat/internal"

type Api struct {
	todo *internal.Todo
}
func NewApi(todo *internal.Todo) *Api {
	return &Api{todo: todo}
}


func (a *Api) List() (*resMenu, error) {


}
