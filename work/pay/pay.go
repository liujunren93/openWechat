package pay

type Api struct {
	todo *core.Todo
}

func NewApi(todo *core.Todo) *Api {
	return &Api{todo: todo}
}

func (a *Api) List() (*resMenu, error) {

}
