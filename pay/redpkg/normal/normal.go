package normal

import (
	"encoding/xml"
	"errors"
	"net/url"

	"github.com/liujunren93/openWechat/pay/redpkg/utils"
	puUtils "github.com/liujunren93/openWechat/utils"
)

type Normal struct {
	XMLName     xml.Name         `xml:"xml"`
	Sign        puUtils.CDATAStr `xml:"sign,omitempty"`
	MchBillno   puUtils.CDATAStr `xml:"mch_billno,omitempty"`
	MchId       puUtils.CDATAStr `xml:"mch_id,omitempty"`
	Wxappid     puUtils.CDATAStr `xml:"wxappid,omitempty"`
	SendName    puUtils.CDATAStr `xml:"send_name"`
	ReOpenid    puUtils.CDATAStr `xml:"re_openid,omitempty"`
	TotalAmount puUtils.CDATAInt `xml:"total_amount,omitempty"`
	TotalNum    puUtils.CDATAInt `xml:"total_num,omitempty"`
	Wishing     puUtils.CDATAStr `xml:"wishing,omitempty"`
	ClientIp    puUtils.CDATAStr `xml:"client_ip,omitempty"`
	ActName     puUtils.CDATAStr `xml:"act_name,omitempty"`
	Remark      puUtils.CDATAStr `xml:"remark,omitempty"`
	SceneId     puUtils.CDATAStr `xml:"scene_id,omitempty"`
	NonceStr    puUtils.CDATAStr `xml:"nonce_str,omitempty"`
	RiskInfo    puUtils.CDATAStr `xml:"risk_info,omitempty"`
}

func (n Normal) GetVal() (url.Values, error) {
	if n.TotalAmount == 0 {
		return nil, errors.New("TotalAmount can not be 0")
	}
	if n.TotalNum == 0 {
		return nil, errors.New("TotalNum can not be 0")
	}
	values := url.Values{}
	values.Add("mch_billno", n.MchBillno.String())
	values.Add("mch_id", n.MchId.String())
	values.Add("wxappid", n.Wxappid.String())
	values.Add("send_name", n.SendName.String())
	values.Add("re_openid", n.ReOpenid.String())
	values.Add("total_amount", n.TotalAmount.String())
	values.Add("total_num", n.TotalNum.String())
	values.Add("wishing", n.Wishing.String())
	values.Add("client_ip", n.ClientIp.String())
	values.Add("act_name", n.ActName.String())
	values.Add("remark", n.Remark.String())
	values.Add("scene_id", n.SceneId.String())
	values.Add("nonce_str", n.NonceStr.String())
	values.Add("risk_info", n.RiskInfo.String())
	return values, nil
}

func (n *Normal) GetData() ([]byte, error) {
	return xml.Marshal(n)
}

func (n *Normal) GetApiPath() string {
	return "https://api.mch.weixin.qq.com/mmpaymkttransfers/sendredpack"
}

func (n *Normal) DoSign(signKey string) error {
	sign, err := utils.Sign(n, signKey)
	if err != nil {
		return err
	}
	n.Sign = puUtils.CDATAStr(sign)
	return nil
}
