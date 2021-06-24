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
	SubButton *Menu  `json:"sub_button,omitempty"`
}

type Menu []*Button

type ReceiveMenu struct {
	IsMenuOpen   int              `json:"is_menu_open"`
	SelfMenuInfo map[string]*Menu `json:"selfmenu_info"`
}

func NewMenu() *Menu{
	return new(Menu)
}

type MenuFunc func(*[]*Button) *Menu

//WithClickBtn 点击推事件用户点击click类型按钮
func (m *Menu) AddClickBtn(name, key string) *Menu{
	subButton := NewMenu()
	*m = append(*m, &Button{
		Name:      name,
		Type:      CLICK,
		Key:       key,
		SubButton: subButton,
	})
	return subButton
}

func (m *Menu) AddViewBtn(name, key, url, appId, pagePath string) *Menu{
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
func (m *Menu) AddScancodeWaitMsgBtn(name, key, url, appid, pagepath string) *Menu{
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
func (m *Menu) AddScancodePushBtn(name, key string) *Menu{
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
func (m *Menu) AddPicSysPhotoBtn(name, key string) *Menu{

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
func (m *Menu) AddPicPhotoOrAlbumBtn(name, key string) *Menu{

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
func (m *Menu) AddPicWeixinBtn(name, key string) *Menu{

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
func (m *Menu) AddLocationSelectBtn(name, key string) *Menu{
	*m = append(*m, &Button{
		Name: name,
		Type: LOCATION_SELECT,
		Key:  key,
	})
	return m

}

func (m *Menu) AddMediaBtn(name, key string) *Menu{
	*m = append(*m, &Button{
		Type: MEDIA_ID,
		Name: name,
		Key:  key,
	})
	return m
}

//AddViewLimitedBtn 图文消息
func (m *Menu) AddViewLimitedBtn(name, key string) *Menu{
	*m = append(*m, &Button{
		Type: VIEW_LIMITED,
		Name: name,
		Key:  key,
	})
	return m
}

type resMenu struct {
	IsMenuOpen   int8 `json:"is_menu_open"`
	SelfMenuInfo struct {
		Button []ResButton
	} `json:"selfmenu_info"`
}

type ResButton struct {
	Name      string `json:"name,omitempty"`
	Type      Type   `json:"type,omitempty"` //click
	Value     string `json:"value,omitempty"`
	Key       string `json:"key,omitempty"`
	AppID     string `json:"appid,omitempty"`
	PagePath  string `json:"pagepath,omitempty"`
	Url       string `json:"url,omitempty"`
	MediaID   string `json:"media_id,omitempty"`
	SubButton struct {
		List *Menu`json:"list"`
	} `json:"sub_button"`
}

func (m *resMenu) ToMenu() *Menu{
	var mu Menu
	for _, button := range m.SelfMenuInfo.Button {
		tmpBtn := Button{
			Name:      button.Name,
			Type:      button.Type,
			Value:     button.Value,
			Key:       button.Key,
			AppID:     button.AppID,
			PagePath:  button.PagePath,
			Url:       button.Url,
			MediaID:   button.MediaID,
			SubButton: button.SubButton.List,
		}
		mu = append(mu, &tmpBtn)
	}
	return &mu
}
