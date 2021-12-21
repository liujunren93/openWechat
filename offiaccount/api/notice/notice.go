package notice

import (
	"github.com/liujunren93/openWechat/internal"
	"github.com/liujunren93/openWechat/utils"
	"sort"
	"strings"
)

type Api struct {
	todo *internal.Todo
}
func NewApi(todo *internal.Todo) *Api {
	return &Api{todo: todo}
}

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
