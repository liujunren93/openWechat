package file

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"sync"
	"time"

	iStore "github.com/liujunren93/openWechat/store"
)

type store struct {
	fileName string
	sync.RWMutex
	storeMap map[string]*iStore.AccessToken
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
	return &store{fileName: fileName, storeMap: map[string]*iStore.AccessToken{}}
}

func (s *store) Load(appId string) (res *iStore.AccessToken, ok bool) {
	if token, ok := s.storeMap[appId]; ok {

		if time.Now().Unix()-token.CreateAt >= 7100 {
			return nil, false
		}
		return token, true
	} else {
		open, err := os.Open(s.fileName)
		defer open.Close()
		if err != nil {
			return nil, false
		}
		var data map[string]*iStore.AccessToken
		all, err := ioutil.ReadAll(open)
		err = json.Unmarshal(all, &data)
		if err != nil {
			return nil, false
		}
		s.Lock()
		s.storeMap = data
		s.Unlock()
		if token, ok := data[appId]; ok {
			if time.Now().Unix()-token.CreateAt >= 7100 {
				return nil, false
			}
			return token, true
		}
		return nil, false
	}

}

func (s *store) IsExpire(appId string) bool {

	if load, ok := s.Load(appId); ok {
		if time.Now().Unix()-load.CreateAt >= 7100 {
			return true
		}
		return false
	}
	return true

}

func (s *store) Store(appId string, accessToken *iStore.AccessToken) {
	s.Lock()
	defer s.Unlock()
	s.storeMap[appId] = accessToken
	marshal, err := json.Marshal(&s.storeMap)
	file, err := os.OpenFile(s.fileName, os.O_RDWR, 0666)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	_, err = file.Write(marshal)
	if err != nil {
		panic(err)
	}
}
