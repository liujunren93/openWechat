package test

import (
	"fmt"

	"os"
	"testing"
	"time"

	"github.com/liujunren93/openWechat/officialAccount/api/material"
	"github.com/liujunren93/openWechat/officialAccount/api/utils/qrcode"
)

//func init() {
//
//	client = officialAccount.NewOfficialAccount("wxf990ce6f695cb376", "f7cc3a4c38a664f52fb1ab302ac2c35a", nil)
//
//}

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
