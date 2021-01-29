package store

type Store interface {
	Load(appId string) (*AccessToken, bool)
	IsExpire(appId string) bool
	Store(appId string, accessToken *AccessToken)
}
