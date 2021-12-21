package notice

import (
	"encoding/xml"
	"github.com/liujunren93/openWechat/types"
	"time"
)

//ReceivingStandardMsgReq 文本消息 接收
type ReceivingStandardMsg struct {
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

func (r ReceivingStandardMsg) GetMsgType() string {
	if r.text.Content != "" {
		return "text"
	}
	if r.pic.PicUrl != "" {
		return "pic"
	}
	if r.voice.Format != "" {
		return "voice"
	}
	if r.video.ThumbMediaId != "" {
		return "video"
	}
	if r.location.Location_X != "" {
		return "location"
	}
	if r.userlocation.Latitude != "" {
		return "userlocation"
	}
	if r.link.Url != "" {
		return "link"
	}
	return ""
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
	ToUserName   types.CDATA `xml:"ToUserName"`
	FromUserName types.CDATA `xml:"FromUserName"`
	CreateTime   int64       `xml:"CreateTime"`
	MsgType      types.CDATA `xml:"MsgType"` //text,image,语音为voice,语音格式：amr,视频为video,小视频为shortvideo,地理位置为location,链接为link
}

func (m *passiveUserReplyMessage) SetBase(toUserName, FromUserName, MediaId string) {
	*m = passiveUserReplyMessage{
		ToUserName:   types.CDATA{Text: toUserName},
		FromUserName: types.CDATA{Text: FromUserName},
		CreateTime:   time.Now().Local().Unix(),
	}
}

//ReplyText 文字
type ReplyText struct {
	XMLName xml.Name `xml:"xml"`
	passiveUserReplyMessage
	Content types.CDATA `xml:"context"`
}

func (m *ReplyText) Val(context string) {
	m.MsgType = types.CDATA{Text: "text"}
	m.Content =types.CDATA{ Text:context}

}

type ReplyImage struct {
	passiveUserReplyMessage
}

type ReplyVoice struct {
	passiveUserReplyMessage
	Voice struct {
		MediaId types.CDATA  `xml:"MediaId"`
	} `xml:"Voice"`
}

type ReplyVideo struct {
	passiveUserReplyMessage
	Video struct {
		MediaId     types.CDATA  `xml:"MediaId"`
		Title       types.CDATA  `xml:"Title"`
		Description types.CDATA  `xml:"Description"`
	} `xml:"Video"`
}

type ReplyMusic struct {
	passiveUserReplyMessage
	Music struct {
		MediaId      types.CDATA  `xml:"MediaId"`
		Title        types.CDATA  `xml:"Title"`
		Description  types.CDATA  `xml:"Description"`
		MusicUrl     types.CDATA  `xml:"MusicUrl"`
		HQMusicUrl   types.CDATA  `xml:"HQMusicUrl"`
		ThumbMediaId types.CDATA  `xml:"ThumbMediaId"`
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
		Title       types.CDATA  `xml:"Title"`
		Description types.CDATA  `xml:"Description"`
		PicUrl      types.CDATA  `xml:"PicUrl"`
		Url         types.CDATA  `xml:"url"`
	}
}
