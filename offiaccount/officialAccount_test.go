package officialAccount

import (
	"fmt"
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
		list, err := account.GetUserList("")
		fmt.Println(list, err)
		time.Sleep(time.Second*100)

	}

}

func TestA(t *testing.T)  {
	account := NewOfficialAccount("wx40a5b2247d31bddf", "5d4677b6498b90282585c573ac324a7a", nil)
	for  {
		list, err := account.GetUserList("")
		fmt.Println(list, err)
		time.Sleep(time.Second*100)
	}
}
