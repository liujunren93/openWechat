package redis

import (
	"encoding/json"
	"github.com/liujunren93/openWechat/store"
	"time"

	"github.com/go-redis/redis"
)

type rStore struct {
	db        *redis.Client
	prefix    string
	namespace string
}

func (r *rStore) Load(namespace, appId string) (store.Data, bool) {
	var res val
	get := r.db.HGet(r.buildKey(), appId)
	bytes, err := get.Bytes()
	if err != nil {
		return nil, false
	}
	err = json.Unmarshal(bytes, &res)
	if err != nil {
		return nil, false
	}else{
		return res, true
	}


}

func (r *rStore) IsExpire(namespace, appId string) bool {
	if data, ok := r.Load(namespace, appId);ok{
		return true
	}else{
		if time.Now().Unix()-data.GetCreateTime() >= data.GetExpire()-100 {
			return true
		}
	}
	return false
}

func (r *rStore) Store(namespace, appId string, val store.Data) error {
	panic("implement me")
}

//func (r *rStore) Store(appId string, data store.Data) {
//	//script := ` if redis.call('hmset',%s,'data',%s,"expireIn",%d,"createdAt",%s) then redis.call('expire',%s,%d);  return 1 else 	return 0 end`
//	//sprintf := fmt.Sprintf(script, r.buildKey(appId), data.GetVal(), data.GetExpire(), data.GetCreateTime(), r.buildKey(appId), data.GetExpire()-100)
//	//r.db.Eval(sprintf,)
//}

func (r *rStore) buildKey() string {

	return r.prefix + ":" + r.namespace
}
