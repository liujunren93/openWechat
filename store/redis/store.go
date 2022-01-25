package redis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	iStore "github.com/liujunren93/openWechat/store"
	"github.com/liujunren93/openWechat/store/memory"
	"time"
)

type store struct {
	db     *redis.Client
	prefix string
	mstore iStore.Store
}

func NewStore(db *redis.Client, prefix string) (*store, error) {
	ping := db.Ping(getContext())
	if ping.Err() != nil {
		return nil, ping.Err()
	}
	return &store{
		db:     db,
		prefix: prefix,
		mstore: memory.NewStore(),
	}, nil
}

func (s store) syncMem(namespace, appId string) (iStore.Data, error) {
	get := s.db.HGet(getContext(), s.buildKey(namespace), appId)
	if get.Err() == redis.Nil {
		return nil, iStore.NilError
	}
	var data iStore.DataVal
	bytes, err := get.Bytes()
	if err != nil {
		return nil, nil
	}
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, nil
	}
	return data, nil
}
func (s store) syncDB(namespace, appId string, data iStore.Data) error {
	marshal, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return s.db.HSet(getContext(), s.buildKey(namespace), appId, marshal).Err()
}

func (s store) Load(namespace, appId string) (iStore.Data, error) {
	if load, ok := s.mstore.Load(namespace, appId); ok == nil {
		return load, nil
	}
	return s.syncMem(namespace, appId)
}

func (s store) IsExpire(namespace, appId string) bool {
	expire := s.mstore.IsExpire(namespace, appId)
	if !expire { //未过期
		return expire
	}
	data, err := s.syncMem(namespace, appId)
	if errors.Is(err, iStore.NilError) {
		return true
	}
	return data.IsExpire()
}

func (s store) Store(namespace, appId string, val iStore.Data) error {
	err := s.mstore.Store(namespace, appId, val)
	if err != nil {
		return err
	}
	return s.syncDB(namespace, appId, val)
}

func (s store) Close() error {
	//if s.db == nil {
	//	return nil
	//}
	//err := s.db.Close()
	//if err == redis.ErrClosed {
	//	return nil
	//}
	//return s.db.Close()
	return nil
}

func (s *store) buildKey(namespace string) string {
	return fmt.Sprintf("%s:%s", s.prefix, namespace)
}

func getContext() context.Context {
	timeout, _ := context.WithTimeout(context.TODO(), time.Second*2)
	return timeout
}
