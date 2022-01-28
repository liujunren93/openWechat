package notice

/**
* @Author: liujunren
* @Date: 2022/1/28 17:05
 */



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
	event
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
	if r.MsgType == "event" {
		return "event"
	}
	return ""
}
type event struct {
	Event string `xml:"Event"`
	EventKey string `xml:"EventKey"`
	Ticket   string `xml:"Ticket"`
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
