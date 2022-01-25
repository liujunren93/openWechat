package file

import (
	"encoding/json"
	"errors"
	"fmt"
	iStore "github.com/liujunren93/openWechat/store"
	"github.com/liujunren93/openWechat/store/memory"
	"io"
	"os"
)

type store struct {
	fileName string
	file     *os.File
	mstore   iStore.Store
}

func (s *store) Load(namespace, appId string) (iStore.Data, error) {
	if load, ok := s.mstore.Load(namespace, appId); ok == nil {
		return load, nil
	}
	return s.syncMem(namespace, appId)
}

//syncMem
func (s *store) syncMem(namespace, appId string) (iStore.Data, error) {
	all, err := s.localAll()
	if err != nil {
		return nil, err
	}
	if val, ok := all[namespace]; ok {
		if data, ok := val[appId]; ok {
			err = s.mstore.Store(namespace, appId, data)
			return data, err
		}
	}
	return nil, iStore.NilError
}
func (s *store) syncFile(namespace, appId string, data iStore.Data) error {
	all, err := s.localAll()
	if err != nil {
		return err
	}
	if all == nil {
		all = map[string]map[string]iStore.DataVal{namespace: {appId: data.(iStore.DataVal)}}
	} else {
		all[namespace]= map[string]iStore.DataVal{appId:data.(iStore.DataVal)}
	}
	marshal, err := json.Marshal(all)
	if err != nil {
		return err
	}
	s.file.Truncate(0)
	_, err = s.file.WriteAt(marshal, 0)
	return err
}
func (s *store) localAll() (map[string]map[string]iStore.DataVal, error) {
	if s.file == nil {
		open, err := os.OpenFile(s.fileName, os.O_CREATE|os.O_RDWR, 0666)
		if err != nil {
			return nil, err
		}
		s.file = open
	}
	s.file.Seek(0, 0)
	var data = make(map[string]map[string]iStore.DataVal)
	all, err := io.ReadAll(s.file)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(all, &data)
	if err != nil {
		return nil, nil
	}

	return data, err
}

func (s *store) IsExpire(namespace, appId string) bool {

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

func (s *store) Store(namespace, appId string, val iStore.Data) error {
	err := s.mstore.Store(namespace, appId, val)
	if err != nil {
		return err
	}
	return s.syncFile(namespace, appId, val)

}

func (s *store) Close() error {
	if s.file == nil {
		return nil
	}
	err := s.file.Close()
	s.file = nil
	if err != nil {
		return err
	}
	return nil
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
	return &store{fileName: fileName, mstore: memory.NewStore()}
}

func (s *store) buildKey(namespace, appId string) string {
	return fmt.Sprintf("%s:%s", namespace, appId)
}
