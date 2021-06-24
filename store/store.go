package store

type Store interface {
	Load(namespace, appId string) (Data, bool)
	IsExpire(namespace, appId string) bool
	//Store(appId string, accessToken *AccessToken)
	Store(namespace, appId string, val Data) error
}

type Data interface {
	SetCreatedTime(int64)
	SetVal(string)
	GetCreateTime() int64
	GetExpire() int64
	GetVal() string

}
