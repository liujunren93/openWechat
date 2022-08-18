package todo

type ToDoFunc func(token string) ([]byte, error)
