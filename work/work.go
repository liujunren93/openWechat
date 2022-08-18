package work

import (
	"github.com/liujunren93/openWechat/store"
	"github.com/liujunren93/openWechat/store/memory"
)

func NewOfficialAccount(appId, AppSecret string, s store.Store) *Client {
	if s == nil {
		s = memory.NewStore()
	}
	todo := core.Todo{}
	todo.SetStore(s)
	todo.SetConf(&core.Config{
		AppID:     appId,
		AppSecret: AppSecret,
	})
	return &Client{toDo: &todo}

}
