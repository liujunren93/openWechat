package notice

import (
	"encoding/xml"
	"time"
)

//ReceivingStandardMsgReq 文本消息 接收
type ReceivingStandardMsgReq struct {
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   int64  `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"` //text,image,语音为voice,语音格式：amr,视频为video,小视频为shortvideo,地理位置为location,链接为link,event
	MsgId        string `xml:"MsgId"`
	MediaId      string `xml:"MediaId"`
	text
	pic
	voice
	video
	location
	link
	userlocation
}
type text struct {
	Content string `xml:"Content"`
}

type pic struct {
	PicUrl string `xml:"PicUrl"`
}
type voice struct {
	Format      string `xml:"Format"` //语音格式 ：amr
	Recognition string `xml:"recognition"`
}

type video struct {
	ThumbMediaId string `xml:"ThumbMediaId"`
}

type location struct {
	Location_X string `xml:"Location_X"`
	Location_Y string `xml:"Location_Y"`
	Scale      string `xml:"scale"`
	Label      string `xml:"Label"`
}

type userlocation struct {
	Latitude  string `xml:"Latitude"`
	Longitude string `xml:"Longitude"`
	Precision string `xml:"Precision"`
}

type link struct {
	Title       string `xml:"Title"`
	Description string `xml:"Description"`
	Url         string `xml:"Url"`
}

type passiveUserReplyMessage struct {
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   int64  `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"` //text,image,语音为voice,语音格式：amr,视频为video,小视频为shortvideo,地理位置为location,链接为link
}

func (m *passiveUserReplyMessage) SetBase(toUserName, FromUserName, MediaId string) {
	*m = passiveUserReplyMessage{
		ToUserName:   toUserName,
		FromUserName: FromUserName,
		CreateTime:   time.Now().Local().Unix(),
	}
}

//ReplyText 文字
type ReplyText struct {
	XMLName xml.Name `xml:"xml"`
	passiveUserReplyMessage
	text
}

func (m *ReplyText) Val(context string) {
	m.MsgType = "text"
	m.Content = context

}

type ReplyImage struct {
	passiveUserReplyMessage
}

type ReplyVoice struct {
	passiveUserReplyMessage
	Voice struct {
		MediaId string `xml:"MediaId"`
	} `xml:"Voice"`
}

type ReplyVideo struct {
	passiveUserReplyMessage
	Video struct {
		MediaId     string `xml:"MediaId"`
		Title       string `xml:"Title"`
		Description string `xml:"Description"`
	} `xml:"Video"`
}

type ReplyMusic struct {
	passiveUserReplyMessage
	Music struct {
		MediaId      string `xml:"MediaId"`
		Title        string `xml:"Title"`
		Description  string `xml:"Description"`
		MusicUrl     string `xml:"MusicUrl"`
		HQMusicUrl   string `xml:"HQMusicUrl"`
		ThumbMediaId string `xml:"ThumbMediaId"`
	} `xml:"Music"`
}

// 图文
type ReplyNews struct {
	passiveUserReplyMessage
	ArticleCount int      `xml:"ArticleCount"`
	Articles     articles `xml:"articles"`
}

type articles struct {
	Item []struct {
		Title       string `xml:"Title"`
		Description string `xml:"Description"`
		PicUrl      string `xml:"PicUrl"`
		Url         string `xml:"url"`
	}
}
