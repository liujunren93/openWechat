package user

import (
	"encoding/json"
	"fmt"
	"github.com/liujunren93/openWechat/offiaccount/internal"
	"github.com/liujunren93/openWechat/utils"
	"net/url"
)

type Api struct {
	todo *internal.Todo
}

func NewApi(todo *internal.Todo) *Api {
	return &Api{todo: todo}
}

//GetAccessToken 获取auth2 token js
func (a *Api) GetAccessToken(code string) (res *AccessTokenRes, err error) {
	api := "https://api.weixin.qq.com/sns/oauth2/access_token?%s"
	apiQuery := url.Values{}
	apiQuery.Set("appid", a.todo.Conf.AppID)
	apiQuery.Set("secret",  a.todo.Conf.AppSecret)
	apiQuery.Set("code", code)
	apiQuery.Set("grant_type", "authorization_code")
	api = fmt.Sprintf(api, apiQuery.Encode())

	get, err := utils.HttpGet(api)
	if err != nil {
		return
	}
	err = json.Unmarshal(get, &res)
	if err != nil {
		return
	}
	return res, err
}


//GetUserInfo 用户详情unionid
func (a *AccessTokenRes)Info() (res *SnsInfoRes, err error) {
	api := "https://api.weixin.qq.com/sns/userinfo?%s"
	apiQuery := url.Values{}
	apiQuery.Set("access_token", a.AccessToken)
	apiQuery.Set("openid", a.OpenID)
	apiQuery.Set("lang", "zh_CN")
	api = fmt.Sprintf(api, apiQuery.Encode())
	get, err := utils.HttpGet(api)
	if err != nil {
		return
	}
	err = json.Unmarshal(get, &res)
	if err != nil {
		return
	}
	return
}

//GetUserList 拉取公众号粉丝列表
func (a *Api) List(next string) (res *ListRes, err error) {
	api := "https://api.weixin.qq.com/cgi-bin/user/get"
	err = a.todo.ToDoFuncGet(api, &res, "next_openid", next)
	return res, err
}

//GetUserInfo 用户详情unionid
func (a *Api) Info(openId string) (res *InfoRes, err error) {
	api := "https://api.weixin.qq.com/cgi-bin/user/info"
	err = a.todo.ToDoFuncGet(api, &res, "openid", openId, "lang", "zh_CN")
	return res, err
}
