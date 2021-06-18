package offiAccount

import (
	"github.com/liujunren93/openWechat/offiaccount/api/material"
	"github.com/liujunren93/openWechat/offiaccount/api/menu"
	"github.com/liujunren93/openWechat/offiaccount/api/signature"
	"github.com/liujunren93/openWechat/offiaccount/api/user"
	"github.com/liujunren93/openWechat/offiaccount/internal"
	"github.com/liujunren93/openWechat/store"
	"github.com/liujunren93/openWechat/store/memory"
)

type Client struct {
	toDo *internal.Todo
}

func NewOfficialAccount(appId, AppSecret string, s store.Store) *Client {
	if s == nil {
		s = memory.NewStore()
	}
	todo := internal.Todo{}
	todo.SetStore(s)
	todo.SetConf(&internal.Config{
		AppID:     appId,
		AppSecret: AppSecret,
	})
	return &Client{toDo: &todo}

}

//UserApi 用户相关
func (o *Client) UserApi() *user.Api {
	return &user.Api{
		Todo: o.toDo,
	}
}

func (o *Client) MenuApi() *menu.Api {
	return &menu.Api{
		Todo: o.toDo,
	}
}

//MaterialApi 素材
func (o *Client) MaterialApi() *material.Api {
	return &material.Api{
		Todo: o.toDo,
	}
}

func (o *Client) Signature() *signature.Api {
	return &signature.Api{
		Todo: o.toDo,
	}

}
