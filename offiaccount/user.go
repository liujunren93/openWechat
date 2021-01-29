package officialAccount

import (
	"github.com/liujunren93/openWechat/offiaccount/entity"
	"github.com/liujunren93/openWechat/offiaccount/internal"
)

func (o *OfficialAccount) GetUserList(next string) (res *entity.UserListRes, err error) {
	api := "https://api.weixin.qq.com/cgi-bin/user/get"
	get, err := internal.ToDoFuncGet(api, &res, "next_openid", next)
	err = o.toDo.Do(get)
	return res, err
}

func (o *OfficialAccount) GetUserInfo(openId string) (res *entity.UserInfoRes, err error) {
	api := "https://api.weixin.qq.com/cgi-bin/user/info"
	get, err := internal.ToDoFuncGet(api, &res, "openid", openId, "lang")
	err = o.toDo.Do(get)
	return res, err

}
