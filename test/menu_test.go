package test

import (
	"fmt"

	officialAccount "github.com/liujunren93/openWechat/officialAccount"

	"github.com/liujunren93/openWechat/officialAccount/api/menu"

	"testing"
)

var client *officialAccount.Client

func init() {

	client = officialAccount.NewOfficialAccount("wxf990ce6f695cb376", "f7cc3a4c38a664f52fb1ab302ac2c35a", nil)

	//client = officialAccount.NewOfficialAccount("wxc0f32e9fab6d06c9", "15ca38c49b2dfbcf1d1d9f4fc22efa2d", nil)
}
func TestSetZbMenu(t *testing.T) {
	newMenu := menu.NewMenu()
	newMenu.AddViewBtn("充值中心", "elec_btn", "http://www.sharelife.club/rim_h5/#/elec/index", "", "")
	newMenu.AddClickBtn("羊毛专区", "waimai_btn").
		AddViewBtn("淘便宜", "tpy_btn", "http://www.sharelife.club/rim_h5/#/tpy/list", "", "").
		AddViewBtn("外卖红包", "meituan_btn", "http://www.sharelife.club/rim_h5/#/takeout/activity?ch=mt", "", "").
		AddViewBtn("出行有礼", "dd_btn", "http://www.sharelife.club/rim_h5/#/takeout/activity?ch=dd", "", "")
	newMenu.AddClickBtn("推广反馈", "tgfk_btn").
		AddViewBtn("意见反馈", "suggest_btn", "http://www.sharelife.club/rim_h5/#/me/suggest", "", "").
		AddClickBtn("我的分享码", "qrcode_btn")
	err := client.MenuApi().Create(newMenu)
	fmt.Println(err)
}
