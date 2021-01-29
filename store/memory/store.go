package memory

import (
	iStore "github.com/liujunren93/openWechat/store"
	"sync"
	"time"
)

type store struct {
	storeMap sync.Map
}

func NewStore() *store {
	return &store{}
}

func (s *store) Load(appId string) (*iStore.AccessToken, bool) {
	if load, ok := s.storeMap.Load(appId); ok {
		token := load.(*iStore.AccessToken)
		if time.Now().Unix()-token.CreateAt >= 7100 {
			return nil, false
		}
		return token, true
	}
	return nil, false
}

func (s *store) Store(appId string, accessToken *iStore.AccessToken) {
	s.storeMap.Store(appId, accessToken)
}

func (s *store) IsExpire(appId string) bool {
	if load, ok := s.storeMap.Load(appId); ok {
		token := load.(*iStore.AccessToken)
		if time.Now().Unix()-token.CreateAt >= 7100 {
			return true
		}
		return false
	}
	return true
}
