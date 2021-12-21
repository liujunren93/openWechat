package normal

import (
	"encoding/xml"
	"errors"
	"github.com/liujunren93/openWechat/pay/redpkg"
	"github.com/liujunren93/openWechat/pay/redpkg/utils"
	"net/url"
)

type Normal struct {
	XMLName     xml.Name        `xml:"xml"`
	Sign        redpkg.CDATAStr `xml:"sign,omitempty"`
	MchBillno   redpkg.CDATAStr `xml:"mch_billno,omitempty"`
	MchId       redpkg.CDATAStr `xml:"mch_id,omitempty"`
	Wxappid     redpkg.CDATAStr `xml:"wxappid,omitempty"`
	SendName    redpkg.CDATAStr `xml:"send_name"`
	ReOpenid    redpkg.CDATAStr `xml:"re_openid,omitempty"`
	TotalAmount redpkg.CDATAInt `xml:"total_amount,omitempty"`
	TotalNum    redpkg.CDATAInt `xml:"total_num,omitempty"`
	Wishing     redpkg.CDATAStr `xml:"wishing,omitempty"`
	ClientIp    redpkg.CDATAStr `xml:"client_ip,omitempty"`
	ActName     redpkg.CDATAStr `xml:"act_name,omitempty"`
	Remark      redpkg.CDATAStr `xml:"remark,omitempty"`
	SceneId     redpkg.CDATAStr `xml:"scene_id,omitempty"`
	NonceStr    redpkg.CDATAStr `xml:"nonce_str,omitempty"`
	RiskInfo    redpkg.CDATAStr `xml:"risk_info,omitempty"`
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
	n.Sign = redpkg.CDATAStr(sign)
	return nil
}
