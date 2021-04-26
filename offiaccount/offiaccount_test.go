package offiAccount

import (
	"fmt"
	"github.com/liujunren93/openWechat/offiaccount/api/menu"
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
	//store := file.NewStore("/Library/WebServer/Documents/gowork/src/openWechat/offiaccount/tt.json")
	newMenu.AddClickBtn("btn1","btn1").AddClickBtn("btn1-1","btn1-1")
	newMenu.AddClickBtn("btn2","btn2").AddViewBtn("btn1-1","btn1-1","http://baidu.com","","")
	newMenu.AddClickBtn("btn3","btn3").AddLocationSelectBtn("local","local").AddPicWeixinBtn("wxPic","wp")
	account := NewOfficialAccount("wx40a5b2247d31bddf", "5d4677b6498b90282585c573ac324a7a", nil)

	err := account.MenuApi().Create(newMenu)
	fmt.Println(err)
}

func TestGetMenu(t *testing.T)  {
	//store := file.NewStore("/Library/WebServer/Documents/gowork/src/openWechat/offiaccount/tt.json")
	account := NewOfficialAccount("wx40a5b2247d31bddf", "5d4677b6498b90282585c573ac324a7a", nil)
	list ,err:= account.MenuApi().List()
	fmt.Println(list,err)
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