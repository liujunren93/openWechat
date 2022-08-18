package template

import "github.com/liujunren93/openWechat/officialAccount/api"

type Template struct {
	Touser      string             `json:"touser"`
	TemplateID  string             `json:"template_id"`
	URL         string             `json:"url"`
	Miniprogram *Miniprogram       `json:"miniprogram"`
	ClientMsgID string             `json:"client_msg_id"`
	Data        map[string]Keyword `json:"data"`
}
type Miniprogram struct {
	Appid    string `json:"appid"`
	Pagepath string `json:"pagepath"`
}

type Keyword struct {
	Value string `json:"value"`
	Color string `json:"color"`
}

type TemplateRes struct {
	api.BaseRes
	Msgid int `json:"msgid"`
}
