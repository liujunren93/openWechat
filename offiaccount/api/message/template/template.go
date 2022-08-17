package template

type Api struct {
	todo *internal.Todo
}

func NewApi(todo *internal.Todo) *Api {
	return &Api{todo: todo}
}
