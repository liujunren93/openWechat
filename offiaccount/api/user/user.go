package user

import (
	"github.com/liujunren93/openWechat/offiaccount/internal"
)

type Api struct {
	*internal.Todo
}

//GetUserList 拉取公众号粉丝列表
func (a *Api) List(next string) (res *ListRes, err error) {
	api := "https://api.weixin.qq.com/cgi-bin/user/get"
	get := internal.ToDoFuncGet(api, &res, "next_openid", next)
	err = a.Do(get)
	return res, err
}

//GetUserInfo 用户详情unionid
func (a *Api) Info(openId string) (res *InfoRes, err error) {
	api := "https://api.weixin.qq.com/cgi-bin/user/info"
	get := internal.ToDoFuncGet(api, &res, "openid", openId, "lang", "zh_CN")
	err = a.Do(get)
	return res, err
}
