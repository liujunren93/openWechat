package memory

import (
	"fmt"
	iStore "github.com/liujunren93/openWechat/store"
	"sync"
)

type store struct {
	storeMap sync.Map
}

func NewStore() *store {
	return &store{}
}

func (s *store) Load(namespace, appId string) (data iStore.Data, err error) {
	if load, ok := s.storeMap.Load(s.buildKey(namespace, appId)); ok {
		data = load.(iStore.Data)

		return data, nil
	}
	return nil, iStore.ExpireError
}

func (s *store) Store(namespace, appId string, data iStore.Data) error {
	s.storeMap.Store(s.buildKey(namespace, appId), data)
	return nil
}

func (s *store) IsExpire(namespace, appId string) bool {
	if load, ok := s.storeMap.Load(s.buildKey(namespace, appId)); ok {
		data := load.(iStore.Data)
		return data.IsExpire()
	}
	return true
}

func (s *store) buildKey(namespace, appId string) string {
	return fmt.Sprintf("%s:%s", namespace, appId)
}
func (s *store) Close() error {
	return nil
}
