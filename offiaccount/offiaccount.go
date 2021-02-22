package offiAccount

import (
	"github.com/liujunren93/openWechat/offiaccount/api/menu"
	"github.com/liujunren93/openWechat/offiaccount/api/user"
	"github.com/liujunren93/openWechat/offiaccount/internal"
	"github.com/liujunren93/openWechat/store"
	"github.com/liujunren93/openWechat/store/memory"
)

type OffiAccount struct {
	toDo *internal.Todo
}

func NewOfficialAccount(appId, AppSecret string, s store.Store) *OffiAccount {
	if s == nil {
		s = memory.NewStore()
	}
	todo := internal.Todo{}
	todo.SetStore(s)
	todo.SetConf(&internal.Config{
		AppID:     appId,
		AppSecret: AppSecret,
	})
	return &OffiAccount{toDo: &todo}

}

//UserApi 用户相关
func (o *OffiAccount) UserApi() *user.Api {
	return &user.Api{
		Todo: o.toDo,
	}
}

func (o *OffiAccount) MenuApi() *menu.Api {
	return &menu.Api{
		Todo: o.toDo,
	}
}

