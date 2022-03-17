package jsapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/liujunren93/openWechat/pay"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/jsapi"
	"github.com/wechatpay-apiv3/wechatpay-go/services/refunddomestic"
)

/**
* @Author: liujunren
* @Date: 2022/2/25 17:00
 */

type JspCreateOrderReq struct {
	OutTradeNo string
	Attach     string
	OpenID     string
	Desc       string
	Amount     int64
}

type PayResult struct {
	Mchid          string    `json:"mchid"`
	Appid          string    `json:"appid"`
	OutTradeNo     string    `json:"out_trade_no"`
	TransactionID  string    `json:"transaction_id"` // 微信交易编号
	TradeType      string    `json:"trade_type"`
	TradeState     string    `json:"trade_state"`
	TradeStateDesc string    `json:"trade_state_desc"`
	BankType       string    `json:"bank_type"`
	Attach         string    `json:"attach"`
	SuccessTime    time.Time `json:"success_time"`
	Payer          struct {
		Openid string `json:"openid"`
	} `json:"payer"`
	Amount struct {
		Total         int    `json:"total"`
		PayerTotal    int    `json:"payer_total"`
		Currency      string `json:"currency"`
		PayerCurrency string `json:"payer_currency"`
	} `json:"amount"`
}

func (r PayResult) Json() ([]byte, error) {
	return json.Marshal(&r)

}

//CreateOrder 创建订单
func CreateOrder(c *pay.Client, ctx context.Context, req JspCreateOrderReq) (response *jsapi.PrepayWithRequestPaymentResponse, err error) {
	svc := jsapi.JsapiApiService{Client: c.Client}
	payment, apiResult, err := svc.PrepayWithRequestPayment(ctx, jsapi.PrepayRequest{
		Appid:       &c.AppID,
		Mchid:       &c.MchID,
		Description: &req.Desc,
		OutTradeNo:  &req.OutTradeNo,
		Attach:      &req.Attach,
		NotifyUrl:   &c.NotifyUrl,
		Amount: &jsapi.Amount{
			Total: core.Int64(req.Amount),
		},
		Payer: &jsapi.Payer{
			Openid: &req.OpenID,
		},
	})
	if apiResult != nil && apiResult.Response.StatusCode != 200 {
		body := apiResult.Response.Body
		defer body.Close()
		all, _ := io.ReadAll(body)

		return nil, fmt.Errorf("refund http errCode:%d,data:%v", apiResult.Response.StatusCode, string(all))
	}
	return payment, err
}

type RefundReq struct {
	TransactionID string
	OrderCode     string
	OutRefundCode string
	Reason        string
	Amount        int64
}

func Refund(c *pay.Client, ctx context.Context, req RefundReq) (resp *refunddomestic.Refund, result *core.APIResult, err error) {
	svc := refunddomestic.RefundsApiService{Client: c.Client}
	currency := "CNY"
	return svc.Create(ctx, refunddomestic.CreateRequest{
		//SubMchid:      &c.mchID,
		TransactionId: &req.TransactionID,
		OutTradeNo:    &req.OrderCode,
		OutRefundNo:   &req.OutRefundCode,
		Reason:        &req.Reason,
		NotifyUrl:     &c.NotifyUrl,
		Amount: &refunddomestic.AmountReq{
			Refund:   &req.Amount,
			Total:    &req.Amount,
			Currency: &currency,
		},
	})
}
