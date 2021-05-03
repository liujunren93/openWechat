package file

import (
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
	if token, ok := s.storeMap[s.buildKey(namespace,appId)]; ok {
		if time.Now().Unix()-token.GetCreateTime() >= 7100 {
			return nil, false
		}
		return token, true
	} else {
		open, err := os.Open(s.fileName)
		defer open.Close()
		if err != nil {
			return nil, false
		}
		var data map[string]iStore.Data
		all, err := ioutil.ReadAll(open)
		err = json.Unmarshal(all, &data)
		if err != nil {
			return nil, false
		}
		s.Lock()
		s.storeMap = data
		s.Unlock()
		if token, ok := data[s.buildKey(namespace,appId)]; ok {
			if time.Now().Unix()-token.GetCreateTime() >= 7100 {
				return nil, false
			}
			return token, true
		}
		return nil, false
	}
}

func (s *store) IsExpire(namespace, appId string) bool {
	if load, ok := s.Load(namespace,appId); ok {
		if time.Now().Unix()-load.GetCreateTime() >= 7100 {
			return true
		}
		return false
	}
	return true
}

func (s *store) Store(namespace, appId string, val iStore.Data) error {
	s.Lock()
	defer s.Unlock()
	s.storeMap[s.buildKey(namespace,appId)] = val
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
