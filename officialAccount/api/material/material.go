package material

import (
	"bytes"
	"encoding/json"
	"os"

	"github.com/liujunren93/openWechat/todo/officialAccount"
)

type Api struct {
	todo *officialAccount.Todo
}

func NewApi(todo *officialAccount.Todo) *Api {
	return &Api{todo: todo}
}

// 新增临时素材
func (a *Api) AddTemporary(media media) (AddTemporaryRes, error) {
	var res AddTemporaryRes
	api := "https://api.weixin.qq.com/cgi-bin/media/upload"
	err := a.todo.ToDoFuncPostForm(api, &res, media, nil, "type", media._type)
	return res, err
}

//AddNews 新增永久图文素材
//@return media_id,error
func (a *Api) AddNews(news ...News) (string, error) {
	var res AddMaterialRes
	api := "https://api.weixin.qq.com/cgi-bin/material/add_news"
	buf := bytes.Buffer{}
	buf.WriteString(`{"articles":`)
	marshal, err := json.Marshal(news)
	if err != nil {
		return "", err
	}
	buf.Write(marshal)
	buf.WriteString("}")
	err = a.todo.ToDoFuncPost(api, &res, buf.Bytes())
	return res.MediaID, err
}

func (a *Api) UploadImg(file *os.File) (string, error) {
	var res AddMaterialRes
	api := "https://api.weixin.qq.com/cgi-bin/media/uploadimg"

	err := a.todo.ToDoFuncPostForm(api, &res, NewImage(file, file.Name()), nil)
	return res.Url, err
}

func (a *Api) AddMaterial(media media, videDesc map[string]string) (AddMaterialRes, error) {
	var res AddMaterialRes
	api := "https://api.weixin.qq.com/cgi-bin/material/add_material"
	err := a.todo.ToDoFuncPostForm(api, &res, media, videDesc, "type", media._type)
	return res, err
}

func (a *Api) MaterialInfo(id string) (MaterialsInfoRes, error) {
	var res MaterialsInfoRes
	api := "https://api.weixin.qq.com/cgi-bin/material/get_material"
	err := a.todo.ToDoFuncPost(api, &res, []byte(`{"media_id":`+id+`}`))
	return res, err
}

//BatchGetMaterial 批量获取素材
//@param mType 图片（image）、视频（video）、语音 （voice）、图文（news）
//@param offset
//@param count
func (a *Api) BatchGetMaterial(mType string, offset, count int) (BatchMaterialRes, error) {
	var res BatchMaterialRes
	api := "https://api.weixin.qq.com/cgi-bin/material/batchget_material"
	var req = map[string]interface{}{"type": mType, "offset": offset, "count": count}
	marshal, _ := json.Marshal(req)
	err := a.todo.ToDoFuncPost(api, &res, marshal)
	return res, err
}
