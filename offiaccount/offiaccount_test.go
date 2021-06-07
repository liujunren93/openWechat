package offiAccount

import (
	"encoding/json"
	"fmt"
	"github.com/liujunren93/openWechat/offiaccount/api/menu"
	"github.com/liujunren93/openWechat/store/file"
	"testing"
	"time"
)

//appId:
//"wx40a5b2247d31bddf"
//appSecret:
//"5d4677b6498b90282585c573ac324a7a"
func TestNewOfficialAccount(t *testing.T) {

	fmt.Println(time.Now())
	account := NewOfficialAccount("wx40a5b2247d31bddf", "5d4677b6498b90282585c573ac324a7a", nil)
	for  {
		list, err := account.UserApi().List("")
		fmt.Println(list, err)
		time.Sleep(time.Second)

	}

}

func TestA(t *testing.T)  {
	account := NewOfficialAccount("wx40a5b2247d31bddf", "5d4677b6498b90282585c573ac324a7a", nil)
	for  {
		list, err := account.UserApi().Info("")
		fmt.Println(list, err)
		time.Sleep(time.Second*100)
	}
}


func TestSetMenu(t *testing.T)  {
	newMenu :=menu.NewMenu()
	store := file.NewStore("/Library/WebServer/Documents/gowork/src/openWechat/offiaccount/tt.json")
	newMenu.AddClickBtn("btn1","btn1").AddScancodePushBtn("scan","scan")
	//newMenu.AddClickBtn("btn2","btn2").AddViewBtn("btn1-1","btn1-1","http://baidu.com","","")
	//newMenu.AddClickBtn("btn3","btn3").AddLocationSelectBtn("local","local").AddPicWeixinBtn("wxPic","wp")
	account := NewOfficialAccount("wxc0f32e9fab6d06c9", "15ca38c49b2dfbcf1d1d9f4fc22efa2d", store)
	fmt.Printf("%#v",newMenu)
	err := account.MenuApi().Create(newMenu)
	fmt.Println(err)
}

func TestSetMenu1(t *testing.T)  {
	var me=`[
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
	var data=menu.NewMenu()
	json.Unmarshal([]byte(me),data)
	fmt.Println(data)

	store := file.NewStore("/Library/WebServer/Documents/gowork/src/openWechat/offiaccount/tt.json")
	account := NewOfficialAccount("wxc0f32e9fab6d06c9", "15ca38c49b2dfbcf1d1d9f4fc22efa2d", store)

	fmt.Printf("%#v",data)
	err := account.MenuApi().Create(data)
	fmt.Println(err)
}

func TestGetMenu(t *testing.T)  {
	account := NewOfficialAccount("wxc0f32e9fab6d06c9", "15ca38c49b2dfbcf1d1d9f4fc22efa2d", nil)
	list ,_:= account.MenuApi().List()
	err := account.MenuApi().Create(list.ToMenu())
	fmt.Println(err)
}

func TestString(t *testing.T)  {
	var a string
	aaa(&a)
	fmt.Println(a)
}

func aaa(s interface{})  {
	s2 := s.(*string)
	*s2="232"
}

func TestOffiAccount_Signature(t *testing.T) {
	account := NewOfficialAccount("wx40a5b2247d31bddf", "5d4677b6498b90282585c573ac324a7a", nil)
	for i := 0; i <10 ; i++ {
		time.Sleep(time.Second)
		account.Signature().Signature("111")
	}

}