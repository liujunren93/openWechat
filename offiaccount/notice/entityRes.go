package notice

/**
* @Author: liujunren
* @Date: 2022/1/28 17:05
 */



//ReceivingStandardMsgReq 文本消息 接收
type ReceivingStandardMsg struct {
	ToUserName   string `xml:"ToUserName" json:"to_user_name,omitempty"`
	FromUserName string `xml:"FromUserName" json:"from_user_name,omitempty"`
	CreateTime   int64  `xml:"CreateTime" json:"create_time,omitempty"`
	MsgType      string `xml:"MsgType" json:"msg_type,omitempty"` //text,image,语音为voice,语音格式：amr,视频为video,小视频为shortvideo,地理位置为location,链接为link,event
	MsgId        string `xml:"MsgId" json:"msg_id,omitempty"`
	MediaId      string `xml:"MediaId" json:"media_id,omitempty"`
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
	Event string `xml:"Event" json:"event,omitempty"`
	EventKey string `xml:"EventKey" json:"event_key,omitempty"`
	Ticket   string `xml:"Ticket" json:"ticket,omitempty"`
}





type text struct {
	Content string `xml:"Content" json:"content,omitempty"`
}

type pic struct {
	PicUrl string `xml:"PicUrl" json:"pic_url,omitempty"`
}
type voice struct {
	Format      string `xml:"Format" json:"format,omitempty"` //语音格式 ：amr
	Recognition string `xml:"recognition" json:"recognition,omitempty"`
}

type video struct {
	ThumbMediaId string `xml:"ThumbMediaId" json:"thumb_media_id,omitempty"`
}

type location struct {
	Location_X string `xml:"Location_X" json:"location_x,omitempty"`
	Location_Y string `xml:"Location_Y" json:"location_y,omitempty"`
	Scale      string `xml:"scale" json:"scale,omitempty"`
	Label      string `xml:"Label" json:"label,omitempty"`
}

type userlocation struct {
	Latitude  string `xml:"Latitude" json:"latitude,omitempty"`
	Longitude string `xml:"Longitude" json:"longitude,omitempty"`
	Precision string `xml:"Precision" json:"precision,omitempty"`
}

type link struct {
	Title       string `xml:"Title" json:"title,omitempty"`
	Description string `xml:"Description" json:"description,omitempty"`
	Url         string `xml:"Url" json:"url,omitempty"`
}
