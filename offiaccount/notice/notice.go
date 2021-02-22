package notice

import (
	"github.com/liujunren93/openWechat/utils"
	"sort"
	"strings"
)

//CheckNoticeSignature 检查消息签名
func CheckNoticeSignature(signature, token, timestamp, nonce string) bool {
	tmp := []string{token, timestamp, nonce}
	sort.Strings(tmp)
	join := strings.Join(tmp, "")
	sha1 := utils.Sha1(join)
	if sha1 == signature {
		return true
	}
	return false

}
