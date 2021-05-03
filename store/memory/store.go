package memory

import (
	"fmt"
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

func (s *store) Load(namespace, appId string) (data iStore.Data, isExpire bool) {
	if load, ok := s.storeMap.Load(s.buildKey(namespace, appId)); ok {
		data = load.(iStore.Data)

		return data, true
	}
	return nil, false
}

func (s *store) Store(namespace,appId string, data iStore.Data) error {
	data.SetCreatedTime(time.Now().Local().Unix())
	s.storeMap.Store(s.buildKey(namespace, appId), data)
	return nil
}

func (s *store) IsExpire(namespace, appId string) bool {
	if load, ok := s.storeMap.Load(s.buildKey(namespace, appId)); ok {
		data := load.(iStore.Data)
		if time.Now().Unix()-data.GetCreateTime() >= data.GetExpire()-100 {
			return true
		}
		return false
	}
	return true
}

func (s *store) buildKey(namespace, appId string) string {
	return fmt.Sprintf("%s:%s", namespace, appId)
}
