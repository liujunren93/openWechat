package redis

type val map[string]interface{}

func (v val) SetCreatedTime(i int64) {

}

func (v val) SetVal(s string) {

}

func (v val) GetCreateTime() int64 {
	return v["created_at"].(int64)
}

func (v val) GetExpire() int64 {
	return v["expires_in"].(int64)
}

func (v val) GetVal() string {
	return v["val"].(string)
}
