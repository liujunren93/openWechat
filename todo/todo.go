package todo

type toDoFunc func(token string) ([]byte, error)
