package internal

import (
	"encoding/json"
	"fmt"
	"github.com/liujunren93/openWechat/store"
	"github.com/liujunren93/openWechat/utils"
	"log"
	"net/url"
	"reflect"
	"strings"
	"sync"
	"time"
)

type Todo struct {
	retry int32
	mutex sync.RWMutex
	Conf  *Config
	store store.Store
}

func (t *Todo) SetConf(conf *Config) {
	t.Conf = conf
}
func (t *Todo) SetStore(store store.Store) {
	t.store = store
}

type toDoFunc func(token string) ([]byte, error)

func ToDoFuncGet(api string, res interface{}, kv ...string) toDoFunc {
	if len(kv)%2 != 0 {
		log.Panic("KV has to be even ")
	}
	return func(token string) ([]byte, error) {
		api := buildApi(api, token, kv)
		re, err := utils.HttpGet(api)
		if err != nil {
			return re, err
		}
		of := reflect.ValueOf(res)
		elem := of.Elem()
		switch elem.Kind(){
		case reflect.String:
			s := res.(*string)
			*s=string(re)
		default:
			err = json.Unmarshal(re, &res)
		}
		return re, err
	}
}

func buildApi(api string, token string, kv []string) string {
	var val = make(url.Values)
	val.Add("access_token", token)
	for i := 0; i < len(kv)/2; i += 2 {
		val.Add(kv[i], kv[i+1])
	}
	return api + "?" + val.Encode()
}

func ToDoFuncPost(api string, res interface{}, data []byte, kv ...string) toDoFunc {
	if len(kv)%2 != 0 {
		log.Panic("KV has to be even ")
	}
	return func(token string) ([]byte, error) {
		api = buildApi(api, token, kv)
		fmt.Println(api)
		return utils.HttpPost(api, nil, data)
	}

}

func (t *Todo) Do(f toDoFunc) error {
	var errRes *ErrorRes
	if t.Conf == nil {
		panic("Conf cannot be empty")
	}
	token, ok := t.store.Load(t.Conf.AppID)
	if !ok {
		token = t.getToken()
	}
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

func (t *Todo) getToken() (token *store.AccessToken) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	apiUrl := "https://api.weixin.qq.com/cgi-bin/token"

	if t.retry > 0 || t.store.IsExpire(t.Conf.AppID) {
		url := apiUrl + "?grant_type=client_credential&AppId=" + t.Conf.AppID + "&secret=" + t.Conf.AppSecret
		get, _ := utils.HttpGet(url)
		json.Unmarshal(get, &token)
		token.CreateAt = time.Now().Unix()

		t.store.Store(t.Conf.AppID, token)
		t.retry = 0
	}
	return token

}
