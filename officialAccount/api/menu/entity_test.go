package menu

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestNewMenu(t *testing.T) {
	newMenu := NewMenu()
	newMenu.AddClickBtn("btn1","afdsa").AddClickBtn("btn1_1","afdsa_1")
	newMenu.AddClickBtn("btn2","afdsa")

	marshal, err := json.Marshal(newMenu)
	fmt.Println(string(marshal),err)

}