package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/liujunren93/openWechat/store"
	"github.com/liujunren93/openWechat/utils"
	"net/url"
	"strings"
	"sync"
	"time"
)

type Todo struct {
	retry int32
	mutex sync.RWMutex
	Conf  *Config
	Store store.Store
}

type toDoFunc func(token string) ([]byte, error)

func ToDoFuncGet(api string, res interface{}, kv ...string) (toDoFunc, error) {
	var val = make(url.Values)
	if len(kv)%2 != 0 {
		return nil, errors.New("KV has to be even ")
	}
	for i := 0; i < len(kv)/2; i += 2 {
		val.Add(kv[i], kv[i+1])
	}
	return func(token string) ([]byte, error) {
		val.Add("access_token", token)
		api := api + "?" + val.Encode()
		re, err := utils.HttpGet(api)
		if err != nil {
			return re, err
		}
		err = json.Unmarshal(re, &res)

		return re, err
	}, nil
}

func (t *Todo) Do(f toDoFunc) error {
	var errRes *ErrorRes
	if t.Conf == nil {
		panic("Conf cannot be empty")
	}
	token, ok := t.Store.Load(t.Conf.AppID)
	if !ok {
		token=t.getToken()
	}
fmt.Println(token.AccessToken)
	bytes, err := f(token.AccessToken)
	if err != nil {
		return err
	}
	if strings.Index(string(bytes), "errcode") > 0 {
		_ = json.Unmarshal(bytes, &errRes)
		if errRes.ErrorCode == 40001 {
			if t.retry >= 3 {
				return errRes
			}
			t.retry++
			return t.Do(f)
		}
		return errRes
	}
	return nil
}

func (t *Todo) getToken() (token *store.AccessToken){
	t.mutex.Lock()
	defer t.mutex.Unlock()
	apiUrl := "https://api.weixin.qq.com/cgi-bin/token"

	if t.retry > 0 || t.Store.IsExpire(t.Conf.AppID) {
		url := apiUrl + "?grant_type=client_credential&AppId=" + t.Conf.AppID + "&secret=" + t.Conf.AppSecret
		get, _ := utils.HttpGet(url)
		json.Unmarshal(get, &token)
		token.CreateAt = time.Now().Unix()

		t.Store.Store(t.Conf.AppID, token)
		t.retry = 0
	}
	return token

}
