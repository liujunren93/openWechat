package offiAccount

import (
	"fmt"
	"github.com/liujunren93/openWechat/offiaccount/api/material"
	"github.com/liujunren93/openWechat/offiaccount/api/menu"
	"github.com/liujunren93/openWechat/store/file"
	"os"
	"testing"
	"time"
)

var client *Client

func init() {
	store := file.NewStore("tt.json")
	client = NewOfficialAccount("wx40a5b2247d31bddf", "5d4677b6498b90282585c573ac324a7a", store)
	//client = NewOfficialAccount("wx40a5b2247d31bddf", "5d4677b6498b90282585c573ac324a7a", nil)

}

//appId:
//"wx40a5b2247d31bddf"
//appSecret:
//"5d4677b6498b90282585c573ac324a7a"
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
	newMenu.AddClickBtn("btn2","btn2").AddViewBtn("btn1-1","btn1-1","http://baidu.com","","")
	newMenu.AddClickBtn("btn3","btn3").AddLocationSelectBtn("local","local").AddPicWeixinBtn("wxPic","wp")
	err := client.MenuApi().Create(newMenu)
	fmt.Println(err)
}

func TestSetZbMenu(t *testing.T) {
	newMenu := menu.NewMenu()
	newMenu.AddViewBtn("淘便宜","tpy_btn","http://www.sharelife.club/rim_h5//#/tpy/list","","")
	newMenu.AddViewBtn("淘便宜","tpy_btn","http://www.sharelife.club/rim_h5//#/tpy/list","","")
	err := client.MenuApi().Create(newMenu)
	fmt.Println(err)
}


func TestGetMenu(t *testing.T) {
	list, err := client.MenuApi().List()
	//err := client.MenuApi().Create(list.ToMenu())
	fmt.Println(list,err)
}

func TestOffiAccount_Signature(t *testing.T) {
		time.Sleep(time.Second)
		signature, err := client.Signature().Signature("111")
		fmt.Println(signature, err )

}

func TestOffiAccount_MaterialApi_UpTemporary(t *testing.T) {
	readFile, err := os.Open("./tt.jpeg")
	fmt.Println(err)

	res, err := client.MaterialApi().AddTemporary(material.NewImage(readFile, ""))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v", res)
}
func TestOffiAccount_MaterialApi_UploadImg(t *testing.T) {
	open, _ := os.Open("./timg.jpeg")
	img, err := client.MaterialApi().UploadImg(open)
	fmt.Println(img, err)
}

func TestOffiAccount_Material(t *testing.T) {

	img, err := client.MaterialApi().BatchGetMaterial("", 0, 10)
	fmt.Println(img, err)
}
func TestOffiAccount_addMaterial(t *testing.T) {
	open, _ := os.Open("./tt.mp4")
	img, err := client.MaterialApi().AddMaterial(material.NewVideo(open, ""),
		map[string]string{"description": `{"title":"VIDEO_TITLE", "introduction":"INTRODUCTION"}`})
	fmt.Println(img, err)
}
func TestOffiAccount_MaterialApi_AddNews(t *testing.T) {
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
	},material.News{
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
func TestOffiAccount_MaterialApiInfo(t *testing.T) {
		info, err := client.MaterialApi().MaterialInfo("cmy4CTVC2HUAA75tQFl9nRIFeG9R9k2Ff0nN6uU2zaM")
		fmt.Println(info, err )

}