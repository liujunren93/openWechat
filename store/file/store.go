package file

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
	"time"

	iStore "github.com/liujunren93/openWechat/store"
)

type store struct {
	fileName string
	sync.RWMutex
	storeMap map[string]iStore.Data
}

func (s *store) Load(namespace, appId string) (iStore.Data, bool) {
begin:
	s.RWMutex.RLock()
	token, ok := s.storeMap[s.buildKey(namespace, appId)]
	s.RWMutex.RUnlock()
	if ok {
		return token, true
	} else {
		open, err := os.Open(s.fileName)
		defer open.Close()
		if err != nil {
			return nil, false
		}
		var data map[string]map[string]interface{}
		all, err := ioutil.ReadAll(open)
		if err != nil || len(all)==0{
			return nil, false
		}
		all = bytes.Trim(all, " ")
		err = json.Unmarshal(all, &data)
		if err != nil {
			return nil, false
		}
		s.Lock()
		for k, m := range data {
			var tmp iStore.Data
			switch m["type"] {
			case "AccessToken":
				tmp = &iStore.AccessToken{
					Val:       m["access_token"].(string),
					ExpiresIn: int64(m["expires_in"].(float64)),
					CreateAt:  int64(m["create_at"].(float64)),
					Type:      "AccessToken",
				}
			case "JsApiTicket":
				tmp = &iStore.JsApiTicket{
					Val:       m["ticket"].(string),
					ExpiresIn: int64(m["expires_in"].(float64)),
					CreateAt:  int64(m["create_at"].(float64)),
					Type:      "JsApiTicket",
				}

			}
			s.storeMap[k] = tmp
		}
		s.Unlock()
		goto begin

	}
}

func (s *store) IsExpire(namespace, appId string) bool {
	if load, ok := s.Load(namespace, appId); ok {
		if time.Now().Local().Unix()-load.GetCreateTime() >= load.GetExpire()-100 {
			return true
		}
		return false
	}
	return true
}

func (s *store) Store(namespace, appId string, val iStore.Data) error {
	s.Lock()
	defer s.Unlock()
	s.storeMap[s.buildKey(namespace, appId)] = val
	marshal, err := json.Marshal(&s.storeMap)
	file, err := os.OpenFile(s.fileName, os.O_RDWR, 0666)
	defer file.Close()
	if err != nil {
		return err
	}
	_, err = file.Write(marshal)
	return err
}

func NewStore(fileName string) *store {
	_, err := os.Stat(fileName)
	if err != nil {
		create, err := os.Create(fileName)
		defer create.Close()
		if err != nil {
			panic(err)
		}
	}
	return &store{fileName: fileName, storeMap: map[string]iStore.Data{}}
}

func (s *store) buildKey(namespace, appId string) string {
	return fmt.Sprintf("%s:%s", namespace, appId)
}
