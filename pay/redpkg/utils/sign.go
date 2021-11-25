package utils

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"net/url"
	"sort"
	"strings"
)

type Signer interface {
	GetVal() (url.Values,error)
}

func Sign(s Signer, key string)(string,error) {
	val,err	 := s.GetVal()

	if err != nil {
		return "", err
	}
	for k, v := range val {
		if v[0] == "" {
			val.Del(k)
		}
	}
	stringSignTemp :=Encode(val) + "&key=" + key
	return md5Up(stringSignTemp)

}

func md5Up(str string) (string, error) {
	hash := md5.New()
	_, err := hash.Write([]byte(str))
	if err != nil {
		return "", err
	}
	return 	strings.ToUpper(hex.EncodeToString(hash.Sum(nil))), nil
}

func sha256Up(str,key string)(string,error)  {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(str))
	return strings.ToUpper(hex.EncodeToString(mac.Sum(nil))), nil
}

func Encode(v url.Values) string {
	if v == nil {
		return ""
	}
	var buf strings.Builder
	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		vs := v[k]
		keyEscaped := k
		for _, v := range vs {
			if buf.Len() > 0 {
				buf.WriteByte('&')
			}
			buf.WriteString(keyEscaped)
			buf.WriteByte('=')
			buf.WriteString(v)
		}
	}
	return buf.String()
}
