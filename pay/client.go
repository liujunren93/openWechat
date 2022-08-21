package pay

import (
	"context"
	"fmt"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/jsapi"
	"github.com/wechatpay-apiv3/wechatpay-go/services/refunddomestic"
	utilsPay "github.com/wechatpay-apiv3/wechatpay-go/utils"
)

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
	NotifyUrl                  string //通知地址
}

func NewClient(conf WechatPayClientConf) (*Client, error) {
	var cli = Client{
		AppID:       conf.AppID,
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

func (c *Client) JspCreateOrder(ctx context.Context, desc, outTradeNo, attach, openid string, amount int64) (response *jsapi.PrepayWithRequestPaymentResponse, result *core.APIResult, err error) {
	fmt.Println("ttt:", desc, outTradeNo, attach, openid, amount)
	svc := jsapi.JsapiApiService{Client: c.Client}
	return svc.PrepayWithRequestPayment(ctx, jsapi.PrepayRequest{
		Appid:       &c.AppID,
		Mchid:       &c.MchID,
		Description: &desc,
		OutTradeNo:  &outTradeNo,
		Attach:      &attach,
		NotifyUrl:   &c.NotifyUrl,
		Amount: &jsapi.Amount{
			Total: core.Int64(amount),
		},
		Payer: &jsapi.Payer{
			Openid: &openid,
		},
	})
}

func (c *Client) Refund(ctx context.Context, transactionId, orderCode, outRefundCode, reason string, amount int64) (resp *refunddomestic.Refund, result *core.APIResult, err error) {
	svc := refunddomestic.RefundsApiService{Client: c.Client}
	currency := "CNY"
	return svc.Create(ctx, refunddomestic.CreateRequest{
		//SubMchid:      &c.mchID,
		TransactionId: &transactionId,
		OutTradeNo:    &orderCode,
		OutRefundNo:   &outRefundCode,
		Reason:        &reason,
		NotifyUrl:     &c.NotifyUrl,
		Amount: &refunddomestic.AmountReq{
			Refund:   &amount,
			Total:    &amount,
			Currency: &currency,
		},
	})
}

func NewRedPkgClient() {}
