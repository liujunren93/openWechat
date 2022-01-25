package store

import (
	"errors"
	"time"
)

type Store interface {
	Load(namespace, appId string) (Data, error)
	IsExpire(namespace, appId string) bool
	//Store(appId string, accessToken *AccessToken)
	Store(namespace, appId string, val Data) error
	Close() error
}

type Data interface {
	SetVal(string)
	SetExpire(expire int64)
	IsExpire() bool
	GetVal() string

}

type DataVal map[string]interface{}

func (da DataVal) SetVal(s string) {
	da["val"] = s
}

func (da DataVal) SetCreateAt() {
	da["create_at"] = time.Now().Local().Unix()
}

func (da DataVal) SetExpire(expire int64) {
	da["expire_in"] = float64(time.Now().Local().Unix() + expire)
	return
}

func (da DataVal) GetVal() string {
	return da["val"].(string)
}

func (da DataVal) IsExpire() bool {
	if da == nil {
		return true
	}
	if int64(da["expire_in"].(float64) )< time.Now().Local().Unix() {
		return true
	}
	return false
}

var ExpireError = errors.New("data was expire")
var NilError = errors.New("data was nil")
