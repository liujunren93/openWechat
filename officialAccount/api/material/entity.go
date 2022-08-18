package material

import (
	"os"
)

type media struct {
	_type string
	name  string
	*os.File
}

func (m media) GetFieldName() string {
	return "media"
}
func (m media) GetType() string {
	return m._type
}

func (m media) GetName() string {
	if m.name == "" {
		return m.File.Name()
	}
	return m.name
}
func (m media) GetData() *os.File {
	return m.File
}

// 图文
type News struct {
	NeedOpenComment    bool   `json:"need_open_comment,omitempty"`     // 是否打开评论，0不打开，1打开
	OnlyFansCanComment bool   `json:"only_fans_can_comment,omitempty"` // 是否粉丝才可评论，0所有人可评论，1粉丝才可评论
	ShowCoverPic       bool   `json:"show_cover_pic,omitempty"  `      //是否显示封面，0为false，即不显示，1为true，即显示
	Title              string `json:"title,omitempty" `
	ThumbMediaID       string `json:"thumb_media_id,omitempty" ` //图文消息的封面图片素材id（必须是永久mediaID）
	Author             string `json:"author,omitempty"`
	Digest             string `json:"digest,omitempty"`  //图文消息的摘要，仅有单图文消息才有摘要，多图文此处为空。如果本字段为没有填写，则默认抓取正文前54个字。
	Content            string `json:"content,omitempty"` //图文消息的具体内容，支持HTML标签，必须少于2万字符，小于1M，且此处会去除JS,涉及图片url必须来源 "上传图文消息内的图片获取URL"接口获取。外部图片url将被过滤。
	Url                string `json:"url,omitempty"`
	ContentSourceUrl   string `json:"content_source_url,omitempty"` //图文消息的原文地址，即点击“阅读原文”后的URL
}
type AddMaterialRes struct {
	MediaID string `json:"media_id"`
	Url     string `json:"url"`
}

type Temporary struct {
	MType     string `json:"type"`
	CreatedAt int64  `json:"created_at"`
	Media     media  `json:"media"`
}

func NewImage(file *os.File, name string) media {
	return media{_type: "image", File: file, name: name}
}
func NewVoice(file *os.File, name string) media {
	return media{_type: "voice", File: file, name: name}
}
func NewVideo(file *os.File, name string) media {
	return media{_type: "video", File: file, name: name}
}
func NewThumb(file *os.File, name string) media {
	return media{_type: "thumb", File: file, name: name}
}

type AddTemporaryRes struct {
	Type      string `json:"type"`
	MediaID   string `json:"media_id"`
	CreatedAt int64  `json:"created_at"`
}
type BatchMaterialRes struct {
	TotalCount int            `json:"total_count"`
	ItemCount  int            `json:"item_count"`
	Item       []MaterialItem `json:"item,omitempty"`
}

type MaterialItem struct {
	MediaID    string `json:"media_id,omitempty"`
	Name       string `json:"name,omitempty"`
	Url        string `json:"url,omitempty"`
	UpdateTime int64  `json:"update_time"`
	Content    struct {
		NewsItem []News `json:"news_item"`
	} `json:"content,omitempty"`
}

type MaterialsInfoRes struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	DownUrl     string `json:"down_url,omitempty"`
	NewsItem    []News `json:"news_item,omitempty"`
}
