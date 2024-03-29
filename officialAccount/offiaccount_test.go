package officialAccount

import (
	"fmt"

	"github.com/liujunren93/openWechat/officialAccount/api/material"
	"github.com/liujunren93/openWechat/officialAccount/api/menu"
	"github.com/liujunren93/openWechat/officialAccount/api/utils/qrcode"

	"os"
	"testing"
	"time"
)

var client *Client

func init() {
	//store := file.NewStore("tt.json")
	// newClient := redis2.NewClient(&redis2.Options{
	// 	Network: "tcp",
	// 	Addr:    "node1:6379",
	// })
	// store, err := redis.NewStore(newClient, "test")
	// if err != nil {
	// 	panic(err)
	// }
	fmt.Println(1111)
	client = NewOfficialAccount("wxc0f32e9fab6d06c9", "15ca38c49b2dfbcf1d1d9f4fc22efa2d", nil)

}

func TestNewOfficialAccount(t *testing.T) {

	fmt.Println(time.Now())
	list, err := client.MaterialApi().BatchGetMaterial("news", 1, 100)
	fmt.Printf("%+v", list)
	fmt.Println(list, err)
	time.Sleep(time.Second)
}

func TestA(t *testing.T) {
	for {
		list, err := client.UserApi().Info("omunIv86qHc7O1rpxuCUNeLkYies")
		fmt.Println(list, err)
		time.Sleep(time.Second * 100)
	}
}

func TestSetMenu(t *testing.T) {
	newMenu := menu.NewMenu()
	newMenu.AddClickBtn("btn1", "btn1").AddScancodePushBtn("scan", "scan")
	newMenu.AddClickBtn("btn2", "btn2").AddViewBtn("btn1-1", "btn1-1", "http://baidu.com", "", "")
	newMenu.AddClickBtn("btn3", "btn3").AddLocationSelectBtn("local", "local").AddPicWeixinBtn("wxPic", "wp")
	err := client.MenuApi().Create(newMenu)
	fmt.Println(err)
}

func TestGetMenu(t *testing.T) {

	list, err := client.MenuApi().List()
	//err := client.MenuApi().Create(list.ToMenu())
	fmt.Println(list, err)

}

func TestofficialAccount_Signature(t *testing.T) {
	time.Sleep(time.Second)
	signature, err := client.Signature().Signature("111")
	fmt.Println(signature, err)

}

func TestofficialAccount_MaterialApi_UpTemporary(t *testing.T) {
	readFile, err := os.Open("./tt.json.jpeg")
	fmt.Println(err)

	res, err := client.MaterialApi().AddTemporary(material.NewImage(readFile, ""))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v", res)
}
func TestofficialAccount_MaterialApi_UploadImg(t *testing.T) {
	open, _ := os.Open("./timg.jpeg")
	img, err := client.MaterialApi().UploadImg(open)
	fmt.Println(img, err)
}

func TestofficialAccount_Material(t *testing.T) {

	img, err := client.MaterialApi().BatchGetMaterial("", 0, 10)
	fmt.Println(img, err)
}
func TestofficialAccount_addMaterial(t *testing.T) {
	open, _ := os.Open("./tt.json.mp4")
	img, err := client.MaterialApi().AddMaterial(material.NewVideo(open, ""),
		map[string]string{"description": `{"title":"VIDEO_TITLE", "introduction":"INTRODUCTION"}`})
	fmt.Println(img, err)
}
func TestofficialAccount_MaterialApi_AddNews(t *testing.T) {
	news, err := client.MaterialApi().AddNews(material.News{
		Title:              "ttt",
		ThumbMediaID:       "Ylfx1KKkztYhcz0ZQzhogZ7oZ-nyaci-h2krZYPUPy0",
		ShowCoverPic:       true,
		Author:             "liujunren",
		Digest:             "",
		Content:            "test http://mmbiz.qpic.cn/mmbiz_jpg/dnNCVVyy7W0j0DU0CDefQRW8ulCq2ib8UnRJMfCKMUWt9Ng9xWmHfjPaeLQyPaQVReOm1Kop7bOMP5Ifhib0ylhA/0?wx_fmt=jpeg",
		ContentSourceUrl:   "https://baidu.com",
		NeedOpenComment:    true,
		OnlyFansCanComment: true,
	}, material.News{
		Title:              "ttt1",
		ThumbMediaID:       "Ylfx1KKkztYhcz0ZQzhogZ7oZ-nyaci-h2krZYPUPy0",
		ShowCoverPic:       true,
		Author:             "liujunren",
		Digest:             "",
		Content:            "test http://mmbiz.qpic.cn/mmbiz_jpg/dnNCVVyy7W0j0DU0CDefQRW8ulCq2ib8UnRJMfCKMUWt9Ng9xWmHfjPaeLQyPaQVReOm1Kop7bOMP5Ifhib0ylhA/0?wx_fmt=jpeg",
		ContentSourceUrl:   "https://baidu.com",
		NeedOpenComment:    true,
		OnlyFansCanComment: true,
	})
	fmt.Println(news, err)

}

//Ylfx1KKkztYhcz0ZQzhogei04pT46O1ZoxO1LhbcBjI
func TestofficialAccount_MaterialApiInfo(t *testing.T) {
	info, err := client.MaterialApi().MaterialInfo("cmy4CTVC2HUAA75tQFl9nRIFeG9R9k2Ff0nN6uU2zaM")
	fmt.Println(info, err)

}

//Ylfx1KKkztYhcz0ZQzhogei04pT46O1ZoxO1LhbcBjI
func TestSign(t *testing.T) {
	for {
		signature, err := client.Signature().Signature("127.0.0.1")
		fmt.Println(signature, err)
	}

}

//Ylfx1KKkztYhcz0ZQzhogei04pT46O1ZoxO1LhbcBjI
func TestUpImage(*testing.T) {
	file, err := os.Open("./aa.png")
	if err != nil {
		panic(err)
	}
	client.Utils().UploadImg(file)

}

func TestofficialAccount_Qrcode(t *testing.T) {
	create, err2 := client.Qrcode().Create(qrcode.Qrcode{
		ExpireSeconds: 3600,
		ActionName:    qrcode.QRLIMITSTRSCENE,
		ActionInfo:    qrcode.ActionInfo(1, "ppx"),
	})

	fmt.Println(create, err2)

}
