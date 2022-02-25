package pay

import (
	"context"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	utilsPay "github.com/wechatpay-apiv3/wechatpay-go/utils"
)

/**
* @Author: liujunren
* @Date: 2022/2/25 16:54
 */

type Client struct {
	Client      *core.Client
	AppID       string
	NotifyUrl   string
	MchID       string
	MchAPIv3Key string
}
type WechatPayClientConf struct {
	AppID                      string
	MchID                      string // 商户号
	MchCertificateSerialNumber string // 商户证书序列号
	MchAPIv3Key                string //商户APIv3密钥
	PrivateKeyPath             string //私钥地址
	NotifyUrl                  string
}

func NewClient(conf WechatPayClientConf) (*Client, error) {
	var cli = Client{AppID: conf.AppID,
		MchID:       conf.MchID,
		NotifyUrl:   conf.NotifyUrl,
		MchAPIv3Key: conf.MchAPIv3Key,
	}
	mchPrivateKey, err := utilsPay.LoadPrivateKeyWithPath(conf.PrivateKeyPath)
	if err != nil {
		return nil, err
	}
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(conf.MchID, conf.MchCertificateSerialNumber, mchPrivateKey, conf.MchAPIv3Key),
	}
	client, err := core.NewClient(context.TODO(), opts...)
	if err != nil {
		return nil, err
	}
	cli.Client = client
	return &cli, err
}
