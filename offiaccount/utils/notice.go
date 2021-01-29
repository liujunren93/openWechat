package utils

import (
	"github.com/liujunren93/openWechat/utils"
	"sort"
	"strings"
)

//CheckNoticeSignature 检查消息签名
func CheckNoticeSignature(signature, timestamp, nonce, token string) bool {
	tmp := []string{signature, timestamp, nonce, token}
	sort.Strings(tmp)
	join := strings.Join(tmp, " ")
	sha1 := utils.Sha1(join)
	if sha1 == signature {
		return true
	}
	return false

}
