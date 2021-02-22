package menu

type Type string

const (
	CLICK              Type = "click"
	VIEW               Type = "view"
	SCANCODE_PUSH      Type = "scancode_push"
	SCANCODE_WAITMSG   Type = "scancode_waitmsg"
	PIC_SYSPHOTO       Type = "pic_sysphoto"
	PIC_PHOTO_OR_ALBUM Type = "pic_photo_or_album"
	PIC_WEIXIN         Type = "pic_weixin"
	LOCATION_SELECT    Type = "location_select"
	MEDIA_ID           Type = "media_id"
	VIEW_LIMITED       Type = "view_limited"
)

type Button struct {
	Name      string `json:"name,omitempty"`
	Type      Type   `json:"type,omitempty"` //click
	Value     string `json:"value,omitempty"`
	Key       string `json:"key,omitempty"`
	AppID     string `json:"appid,omitempty"`
	PagePath  string `json:"pagepath,omitempty"`
	Url       string `json:"url,omitempty"`
	MediaID   string `json:"media_id,omitempty"`
	SubButton *menu  `json:"sub_button,omitempty"`
}

type menu []*Button

type ReceiveMenu struct {
	IsMenuOpen   int              `json:"is_menu_open"`
	SelfmenuInfo map[string]*menu `json:"selfmenu_info"`
}

func NewMenu() *menu {
	return new(menu)
}

type MenuFunc func(*[]*Button) *menu

//WithClickBtn 点击推事件用户点击click类型按钮
func (m *menu) AddClickBtn(name, key string) *menu {
	subButton := NewMenu()
	*m = append(*m, &Button{
		Name:      name,
		Type:      CLICK,
		Key:       key,
		SubButton: subButton,
	})
	return subButton
}

func (m *menu) AddViewBtn(name, key, url, appId, pagePath string) *menu {
	*m = append(*m, &Button{
		Name:     name,
		Type:     VIEW,
		Key:      key,
		Url:      url,
		AppID:    appId,
		PagePath: pagePath,
	})
	return nil
}

//AddScancodeWaitMsgBtn 扫码带提示
func (m *menu) AddScancodeWaitMsgBtn(name, key, url, appid, pagepath string) *menu {
	subButton := NewMenu()
	*m = append(*m, &Button{
		Url:       url,
		AppID:     appid,
		PagePath:  pagepath,
		Name:      name,
		Type:      SCANCODE_WAITMSG,
		Key:       key,
		SubButton: subButton,
	})
	return subButton

}

//AddScancodePushBtn 扫码推事件
func (m *menu) AddScancodePushBtn(name, key string) *menu {
	subButton := NewMenu()

	*m = append(*m, &Button{
		Name:      name,
		Type:      SCANCODE_PUSH,
		Key:       key,
		SubButton: subButton,
	})
	return subButton

}

//AddPicSysPhotoBtn 系统拍照发图
func (m *menu) AddPicSysPhotoBtn(name, key string) *menu {

	subButton := NewMenu()
	*m = append(*m, &Button{
		Name:      name,
		Type:      PIC_SYSPHOTO,
		Key:       key,
		SubButton: subButton,
	})
	return subButton

}

//AddPicPhotoOrAlbumBtn 拍照或者相册发图
func (m *menu) AddPicPhotoOrAlbumBtn(name, key string) *menu {

	subButton := NewMenu()
	*m = append(*m, &Button{
		Name:      name,
		Type:      PIC_PHOTO_OR_ALBUM,
		Key:       key,
		SubButton: subButton,
	})
	return subButton

}

//AddPicWeixinBtn 微信相册发图
func (m *menu) AddPicWeixinBtn(name, key string) *menu {

	subButton := NewMenu()
	*m = append(*m, &Button{
		Name:      name,
		Type:      PIC_WEIXIN,
		Key:       key,
		SubButton: subButton,
	})
	return subButton

}

//AddLocationSelectBtn 发送位置
func (m *menu) AddLocationSelectBtn(name, key string) *menu {
	*m = append(*m, &Button{
		Name: name,
		Type: LOCATION_SELECT,
		Key:  key,
	})
	return m

}

func (m *menu) AddMediaBtn(name, key string) *menu {
	*m = append(*m, &Button{
		Type: MEDIA_ID,
		Name: name,
		Key:  key,
	})
	return m
}

//AddViewLimitedBtn 图文消息
func (m *menu) AddViewLimitedBtn(name, key string) *menu {
	*m = append(*m, &Button{
		Type: VIEW_LIMITED,
		Name: name,
		Key:  key,
	})
	return m
}
