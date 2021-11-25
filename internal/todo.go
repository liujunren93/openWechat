package internal

import (
	"encoding/json"
	"github.com/liujunren93/openWechat/client"
	"github.com/liujunren93/openWechat/offiaccount/consts"
	"github.com/liujunren93/openWechat/store"
	"github.com/liujunren93/openWechat/store/memory"
	"github.com/liujunren93/openWechat/utils"
	"log"
	"net/url"
	"reflect"
	"sync"
	"time"
)

type Todo struct {
	retry   int32
	mutex   sync.RWMutex
	Conf    *Config
	store   store.Store
	appType string
}

func NewTodo(conf *Config, appType string) *Todo {
	return &Todo{
		retry:   3,
		mutex:   sync.RWMutex{},
		Conf:    conf,
		store:   memory.NewStore(),
		appType: appType,
	}
}



func (t *Todo) SetStore(store store.Store) {
	t.store = store
}

type toDoFunc func(token string) ([]byte, error)

func buildApi(api string, token string, kv []string) string {
	var val = make(url.Values)
	val.Add("access_token", token)
	for i := 0; i < len(kv)/2; i += 2 {
		val.Add(kv[i], kv[i+1])
	}
	return api + "?" + val.Encode()
}

func (t *Todo) ToDoFuncGet(api string, res interface{}, kv ...string) error {
	if len(kv)%2 != 0 {
		log.Panic("KV has to be even ")
	}
	return t.do(func(token string) ([]byte, error) {
		api := buildApi(api, token, kv)
		re, err := client.HttpGet(api)
		if err != nil {
			return re, err
		}
		of := reflect.ValueOf(res)
		elem := of.Elem()
		switch elem.Kind() {
		case reflect.String:
			s := res.(*string)
			*s = string(re)
		default:
			err = json.Unmarshal(re, &res)
		}
		return re, err
	})
}

func (t *Todo) ToDoFuncPostForm(api string, res interface{}, file utils.File, data map[string]string, kv ...string) error {
	if len(kv)%2 != 0 {
		log.Panic("KV has to be even ")
	}
	return t.do(func(token string) ([]byte, error) {
		api = buildApi(api, token, kv)
		re, err := client.HttpPostForm(api, data, file)
		if err != nil {
			return re, err
		}
		if res != nil {
			of := reflect.ValueOf(res)
			elem := of.Elem()
			switch elem.Kind() {
			case reflect.String:
				s := res.(*string)
				*s = string(re)
			default:
				err = json.Unmarshal(re, &res)
			}
		}

		return re, nil

	})

}

func (t *Todo) ToDoFuncPost(api string, res interface{}, data []byte, kv ...string) error {
	if len(kv)%2 != 0 {
		log.Panic("KV has to be even ")
	}
	return t.do(func(token string) ([]byte, error) {
		api = buildApi(api, token, kv)
		re, err := client.HttpPost(api, data)
		if err != nil {
			return re, err
		}
		if res != nil {
			of := reflect.ValueOf(res)
			elem := of.Elem()
			switch elem.Kind() {
			case reflect.String:
				s := res.(*string)
				*s = string(re)

			default:
				err = json.Unmarshal(re, &res)
			}
		}

		return re, nil

	})

}

func (t *Todo) do(f toDoFunc) error {
	var errRes *consts.ErrorRes
	if t.Conf == nil {
		panic("Conf cannot be empty")
	}

retry:
	token, err := t.getAccessToken()
	if err != nil {
		return err
	}
	bytes, err := f(token.GetVal())
	if err != nil {
		return err
	}
	//if strings.Index(string(bytes), "errcode") > 0 {
	_ = json.Unmarshal(bytes, &errRes)
	if errRes.ErrorCode == 40001 {
		if t.retry >= 3 {
			return errRes
		}
		t.retry++

		goto retry
	} else if errRes.ErrorCode == 0 {
		t.retry = 0
		return nil
	}
	return errRes
}

// 获取accessToken
func (t *Todo) getAccessToken() (token *store.AccessToken, err error) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	apiUrl := "https://api.weixin.qq.com/cgi-bin/token"
	if t.retry > 0 || t.store.IsExpire(t.appType+"accessToken", t.Conf.AppID) {
		apiUrl += "?grant_type=client_credential&AppId=" + t.Conf.AppID + "&secret=" + t.Conf.AppSecret
		get, _ := client.HttpGet(apiUrl)
		err = json.Unmarshal(get, &token)
		if err != nil {
			return
		}
		token.SetCreatedTime(time.Now().Local().Unix())
		token.Type = "AccessToken"
		err = t.store.Store(t.appType+"accessToken", t.Conf.AppID, token)
		t.retry = 0
	} else {
		if val, ok := t.store.Load(t.appType+"accessToken", t.Conf.AppID); ok && t.retry == 0 {
			if time.Now().Local().Unix()-val.GetCreateTime() < val.GetExpire()-100 { // 未过期
				return val.(*store.AccessToken), nil
			}
		}
	}

	return

}

//
func (t *Todo) GetTicket() (string, error) {
	// 判断是否过期
	if load, ok := t.store.Load(t.appType+"ticket", t.Conf.AppID); ok {
		if time.Now().Local().Unix()-load.GetCreateTime() < load.GetExpire()-100 { // 未过期
			return load.GetVal(), nil
		}
	}
	var res store.JsApiTicket
	res.Type = "JsApiTicket"
	err := t.ToDoFuncGet("https://api.weixin.qq.com/cgi-bin/ticket/getticket", &res, "type", "jsapi")

	if err != nil {
		return "", err
	}
	if res.ErrCode != 0 {
		return "", err
	}
	err = t.store.Store(t.appType+"ticket", t.Conf.AppID, &res)
	return res.GetVal(), nil
}
