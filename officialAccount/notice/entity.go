package notice

import (
	"encoding/xml"
	"time"

	"github.com/liujunren93/openWechat/types"
)

type passiveUserReplyMessage struct {
	XMLName      xml.Name    `xml:"xml"`
	ToUserName   types.CDATA `xml:"ToUserName"`
	FromUserName types.CDATA `xml:"FromUserName"`
	CreateTime   int64       `xml:"CreateTime"`
	MsgType      types.CDATA `xml:"MsgType"` //text,image,语音为voice,语音格式：amr,视频为video,小视频为shortvideo,地理位置为location,链接为link
}

func (m *passiveUserReplyMessage) SetBase(toUserName, FromUserName, MediaId string) {
	m.ToUserName = types.CDATA{Text: toUserName}
	m.FromUserName = types.CDATA{Text: FromUserName}
	m.CreateTime = time.Now().Local().Unix()
}

//ReplyText 文字
type ReplyText struct {
	XMLName xml.Name `xml:"xml"`
	passiveUserReplyMessage
	Content types.CDATA `xml:"Content"`
}

func NewReplyText(toUserName, FromUserName, content string) ReplyText {
	text := ReplyText{}
	text.SetBase(toUserName, FromUserName, "")
	text.Content = types.CDATA{Text: content}
	text.MsgType = types.CDATA{Text: "text"}
	return text
}

func (m *ReplyText) SetVal(context string) {
	m.MsgType = types.CDATA{Text: "text"}
	m.Content = types.CDATA{Text: context}

}

type ReplyImage struct {
	passiveUserReplyMessage
}

type ReplyVoice struct {
	passiveUserReplyMessage
	Voice struct {
		MediaId types.CDATA `xml:"MediaId"`
	} `xml:"Voice"`
}

type ReplyVideo struct {
	passiveUserReplyMessage
	Video struct {
		MediaId     types.CDATA `xml:"MediaId"`
		Title       types.CDATA `xml:"Title"`
		Description types.CDATA `xml:"Description"`
	} `xml:"Video"`
}

type ReplyMusic struct {
	passiveUserReplyMessage
	Music struct {
		MediaId      types.CDATA `xml:"MediaId"`
		Title        types.CDATA `xml:"Title"`
		Description  types.CDATA `xml:"Description"`
		MusicUrl     types.CDATA `xml:"MusicUrl"`
		HQMusicUrl   types.CDATA `xml:"HQMusicUrl"`
		ThumbMediaId types.CDATA `xml:"ThumbMediaId"`
	} `xml:"Music"`
}

// 图文
type ReplyNews struct {
	passiveUserReplyMessage
	ArticleCount int       `xml:"ArticleCount"`
	Articles     *articles `xml:"Articles"`
}

func NewReplyNews() *ReplyNews {
	return &ReplyNews{
		passiveUserReplyMessage: passiveUserReplyMessage{
			MsgType: types.CDATA{Text: "news"},
		},
		ArticleCount: 0,
		Articles:     &articles{},
	}

}

func (r *ReplyNews) AddItem(title, description, picUrl, url string) {
	r.ArticleCount++
	if r.Articles == nil {
		r.Articles = &articles{}
	}

	r.Articles.Item = append(r.Articles.Item, ReplyNewsItem{
		Title:       types.CDATA{Text: title},
		Description: types.CDATA{Text: description},
		PicUrl:      types.CDATA{Text: picUrl},
		Url:         types.CDATA{Text: url},
	})
}

type articles struct {
	Item []ReplyNewsItem
}

type ReplyNewsItem struct {
	XMLName     xml.Name    `xml:"item"`
	Title       types.CDATA `xml:"Title"`
	Description types.CDATA `xml:"Description"`
	PicUrl      types.CDATA `xml:"PicUrl"`
	Url         types.CDATA `xml:"Url"`
}
