package msg

import "github.com/liujunren93/openWechat/internal"

/**
* @Author: liujunren
* @Date: 2021/12/21 10:14
 */
type Api struct {
	todo *internal.Todo
}
func NewApi(todo *internal.Todo) *Api {
	return &Api{todo: todo}
}
func (Api) PassiveUserReplyMessage() {
	//var res AddTemporaryRes
	//api := "https://api.weixin.qq.com/cgi-bin/media/upload"
	//err := a.todo.ToDoFuncPostForm(api, &res, media, nil, "type", media._type)
	//return res, err
}
