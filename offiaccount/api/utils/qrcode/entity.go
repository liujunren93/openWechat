package qrcode

/**
* @Author: liujunren
* @Date: 2022/1/28 14:37
 */

type Action string

const (
	QRSCENE         Action = "QR_SCENE"           //临时的整型;
	QRSTRSCENE      Action = "QR_STR_SCENE"       //:临时的字符串参数值
	QRLIMITSCENE    Action = "QR_LIMIT_SCENE"     //:永久的整型参数值
	QRLIMITSTRSCENE Action = "QR_LIMIT_STR_SCENE" //:永久的字符串参数值
)

type Qrcode struct {
	ExpireSeconds int        `json:"expire_seconds"`
	ActionName    Action     `json:"action_name"`
	ActionInfo    actionInfo `json:"action_info"`
}

func ActionInfo(sceneID uint32, sceneStr string) actionInfo {
	return actionInfo{struct {
		SceneID  uint32 `json:"scene_id"`
		SceneStr string `json:"scene_str"`
	}{SceneID: sceneID, SceneStr: sceneStr}}
}

type actionInfo struct {
	Scene struct {
		SceneID  uint32 `json:"scene_id"`
		SceneStr string `json:"scene_str"`
	} `json:"scene"`
}
type CreatQrcodeRes struct {
	Ticket        string `json:"ticket"`
	ExpireSeconds int    `json:"expire_seconds"`
	Url           string `json:"url"`
	QrcodeUrl     string `json:"qrcode_url"`
}
