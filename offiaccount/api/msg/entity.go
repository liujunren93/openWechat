package msg

/**
* @Author: liujunren
* @Date: 2021/12/21 10:30
 */
import (
	"encoding/xml"
	"github.com/liujunren93/openWechat/types"
)

type Base struct {
	XMLName      xml.Name    `xml:"xml"`
	ToUserName   types.CDATA `xml:"ToUserName"`
	FromUserName types.CDATA `xml:"FromUserName"`
	CreateTime   int64       `xml:"CreateTime"`
	MsgType      types.CDATA `xml:"msg_type"`
}

//ImgTxtReq 图文
type ImgTxtReq struct {
	Base
	ArticleCount int64 `xml:"ArticleCount"`
	Articles     Item  `xml:"Articles"`
}

func (req *ImgTxtReq) AddItems(items ...ArticlesItem) {
	if len(req.Articles.Items)==0 {
		req.Articles.Items=items
	}else{
		for _, item := range items {
			req.Articles.Items = append(req.Articles.Items, item)
		}
	}


}

type Item struct {
	Items []ArticlesItem
}

type ArticlesItem struct {
	XMLName     xml.Name    `xml:"Item"`
	Title       types.CDATA `xml:"Title"`
	Description types.CDATA `xml:"Description"`
	PicUrl      types.CDATA `xml:"PicUrl"`
	Url         types.CDATA `xml:"Url"`
}
