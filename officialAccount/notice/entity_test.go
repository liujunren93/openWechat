package notice

import (
	"encoding/xml"
	"fmt"
	"github.com/liujunren93/openWechat/types"
	"testing"
)

func TestNewReplyText(t *testing.T) {
	var a = ReplyImage{
		passiveUserReplyMessage: passiveUserReplyMessage{
			ToUserName:   types.CDATA{"111"},
			FromUserName: types.CDATA{"3333"},
			CreateTime:   0,
			MsgType:      types.CDATA{"222"},
		},
		ArticleCount: 0,
	}
	a.AddItem("test", "test", "test1", "test1")
	a.AddItem("test", "test", "test1", "test1")
	a.AddItem("test", "test", "test1", "test1")
	a.AddItem("test", "test", "test1", "test1")
	marshal, err := xml.Marshal(a)
	fmt.Println(string(marshal), err)

}
