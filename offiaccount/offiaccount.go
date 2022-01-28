package offiAccount

import (
	"github.com/liujunren93/openWechat/internal"
	"github.com/liujunren93/openWechat/offiaccount/api/material"
	"github.com/liujunren93/openWechat/offiaccount/api/menu"
	"github.com/liujunren93/openWechat/offiaccount/api/signature"
	"github.com/liujunren93/openWechat/offiaccount/api/user"
	"github.com/liujunren93/openWechat/offiaccount/api/utils/qrcode"
	"github.com/liujunren93/openWechat/offiaccount/api/utils/upload"
	"github.com/liujunren93/openWechat/store"
	"sync"
)

type Client struct {
	toDo   *internal.Todo
	apiMap sync.Map
}

func NewOfficialAccount(appId, AppSecret string, s store.Store) *Client {
	todo := internal.NewTodo(&internal.Config{
		AppID:     appId,
		AppSecret: AppSecret,
	},"offiaccount")
	if s != nil {
		todo.SetStore(s)
	}
	return &Client{toDo: todo}

}



//UserApi 用户相关
func (o *Client) UserApi() *user.Api {
	if v, ok := o.apiMap.Load(user.Api{}); ok {
		return v.(*user.Api)
	} else {
		api := user.NewApi(o.toDo)
		o.apiMap.Store(user.Api{}, api)
		return api
	}
}

func (o *Client) MenuApi() *menu.Api {
	if v, ok := o.apiMap.Load(menu.Api{}); ok {
		return v.(*menu.Api)
	} else {
		api := menu.NewApi(o.toDo)
		o.apiMap.Store(menu.Api{}, api)
		return api
	}

}

//MaterialApi 素材
func (o *Client) MaterialApi() *material.Api {
	if v, ok := o.apiMap.Load(material.Api{}); ok {
		return v.(*material.Api)
	} else {
		api := material.NewApi(o.toDo)
		o.apiMap.Store(material.Api{}, api)
		return api
	}
}

func (o *Client) Signature() *signature.Api {
	if v, ok := o.apiMap.Load(signature.Api{}); ok {
		return v.(*signature.Api)
	} else {
		api := signature.NewApi(o.toDo)
		o.apiMap.Store(signature.Api{}, api)
		return api
	}

}

func (o *Client) Utils() *upload.Api {
	if v, ok := o.apiMap.Load(upload.Api{}); ok {
		return v.(*upload.Api)
	} else {
		api := upload.NewApi(o.toDo)
		o.apiMap.Store(upload.Api{}, api)
		return api
	}

}


func (o *Client) Qrcode() *qrcode.Api {
	if v, ok := o.apiMap.Load(qrcode.Api{}); ok {
		return v.(*qrcode.Api)
	} else {
		api := qrcode.NewApi(o.toDo)
		o.apiMap.Store(qrcode.Api{}, api)
		return api
	}

}

