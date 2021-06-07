package user

import (
	"encoding/json"
	"fmt"
	"github.com/liujunren93/openWechat/offiaccount/internal"
	"github.com/liujunren93/openWechat/utils"
	"net/url"
)

type Api struct {
	*internal.Todo
}


//GetAccessToken 获取token
func (a *Api) GetAccessToken(code string) (res *AccessTokenRes, err error) {
	api := "https://api.weixin.qq.com/sns/oauth2/access_token?%s"
	apiQuery := url.Values{}
	apiQuery.Set("appid", a.Conf.AppID)
	apiQuery.Set("secret",  a.Conf.AppSecret)
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