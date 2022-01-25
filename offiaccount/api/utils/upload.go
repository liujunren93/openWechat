package utils

import (
	"fmt"
	"github.com/liujunren93/openWechat/internal"
	"os"
)

/**
* @Author: liujunren
* @Date: 2022/1/25 17:04
 */
type Api struct {
	todo *internal.Todo
}

func NewApi(todo *internal.Todo) *Api {
	return &Api{todo: todo}
}

type image struct {
	f *os.File
}

func (i image) GetFieldName() string {
	return "buffer"
}

func (i image) GetName() string {
	return i.f.Name()
}

func (i image) GetData() *os.File {
	return i.f
}

func (a Api) UploadImg(f *os.File) {
	i := image{f:f}
	api := "https://api.weixin.qq.com/cgi-bin/media/uploadimg"
	var data interface{}
	 a.todo.ToDoFuncPostForm(api, &data, i,nil)
	fmt.Println(data)
}
