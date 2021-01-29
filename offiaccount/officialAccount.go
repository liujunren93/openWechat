package officialAccount

import (
	"github.com/liujunren93/openWechat/offiaccount/internal"
	"github.com/liujunren93/openWechat/store"
	"github.com/liujunren93/openWechat/store/memory"
)

type OfficialAccount struct {
	toDo *internal.Todo
}

func NewOfficialAccount(appId, AppSecret string, s store.Store) *OfficialAccount {
	if s == nil {
		s = memory.NewStore()
	}
	todo := internal.Todo{
		Conf: &internal.Config{
			AppID:     appId,
			AppSecret: AppSecret,
		},
		Store: s,
	}
	return &OfficialAccount{toDo: &todo}

}


