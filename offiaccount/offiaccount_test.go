package offiAccount

import (
	"encoding/json"
	"fmt"
	"github.com/liujunren93/openWechat/offiaccount/api/material"
	"github.com/liujunren93/openWechat/offiaccount/api/menu"
	"github.com/liujunren93/openWechat/offiaccount/consts"
	"github.com/liujunren93/openWechat/store/file"
	"os"
	"testing"
	"time"
)

var client *Client

func init() {
	store := file.NewStore("/Library/WebServer/Documents/gowork/src/openWechat/offiaccount/tt")
	//client = NewOfficialAccount("wx40a5b2247d31bddf", "5d4677b6498b90282585c573ac324a7a", store)
	client = NewOfficialAccount("wxc0f32e9fab6d06c9", "15ca38c49b2dfbcf1d1d9f4fc22efa2d", store)

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
		list, err := client.UserApi().Info("ouKLl6cS6JaR7Rqvxl220gwRjnC0")
		fmt.Println(list, err)
		time.Sleep(time.Second * 100)
	}
}

func TestSetMenu(t *testing.T) {
	newMenu := menu.NewMenu()
	newMenu.AddClickBtn("btn1", "btn1").AddScancodePushBtn("scan", "scan")
	//newMenu.AddClickBtn("btn2","btn2").AddViewBtn("btn1-1","btn1-1","http://baidu.com","","")
	//newMenu.AddClickBtn("btn3","btn3").AddLocationSelectBtn("local","local").AddPicWeixinBtn("wxPic","wp")
	err := client.MenuApi().Create(newMenu)
	fmt.Println(err)
}

func TestSetMenu1(t *testing.T) {
	var me = `[
            {
                "name":"btn1111",
                "sub_button":[
                        {
                            "type":"click",
                            "name":"btn111111-1",
                            "key":"btn1-1"
                        }
                    ]
                
            },
            {
                "name":"btn2",
                "sub_button":[
                        {
                            "type":"view",
                            "name":"btn1-1",
                            "url":"http:\/\/baidu.com"
                        }
                    ]
                
            },
            {
                "name":"btn3",
                "sub_button":[
                        {
                            "type":"location_select",
                            "name":"local",
                            "key":"local"
                        },
                        {
                            "type":"pic_weixin",
                            "name":"wxPic",
                            "key":"wp"
                        }
                    ]
                
            }
        ]`
	var data = menu.NewMenu()
	json.Unmarshal([]byte(me), data)
	fmt.Println(data)

	fmt.Printf("%#v", data)
	err := client.MenuApi().Create(data)
	fmt.Println(err)
}

func TestGetMenu(t *testing.T) {
	list, _ := client.MenuApi().List()
	err := client.MenuApi().Create(list.ToMenu())
	fmt.Println(err)
}

func TestOffiAccount_Signature(t *testing.T) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		client.Signature().Signature("111")
	}

}

func TestOffiAccount_MaterialApi_UpTemporary(t *testing.T) {
	readFile, err := os.Open("./timg.jpeg")
	fmt.Println(err)

	res, err := client.MaterialApi().AddTemporary(material.NewImage(readFile, ""))
	if err != nil {
		if er, ok := err.(*consts.ErrorRes); ok {
			fmt.Println(er)
		}

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

	img, err := client.MaterialApi().BatchGetMaterial("video", 0, 10)
	fmt.Println(img, err)
}
func TestOffiAccount_addMaterial(t *testing.T) {
	open, _ := os.Open("./tt.jpeg")
	img, err := client.MaterialApi().AddMaterial(material.NewImage(open, ""),
		map[string]string{"description": `{"title":"VIDEO_TITLE", "introduction":"INTRODUCTION"}`})
	fmt.Println(img, err)
}
func TestOffiAccount_MaterialApi_AddNews(t *testing.T) {
	news, err := client.MaterialApi().AddNews(material.News{
		Title:              "test",
		ThumbMediaID:       "cmy4CTVC2HUAA75tQFl9neng-VmHJzdTANsCRUPK1Ys",
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
