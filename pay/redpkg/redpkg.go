package redpkg

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/liujunren93/openWechat/pay/redpkg/normal"
)

type Client struct {
	// wechatPayCertPath string //apiclient_cert.pem
	// wechatPayKeyPath  string //apiclient_key.pem
	// rootCPath         string //apiclient_cert.p12
	SignKey string //mchAPIv3Key
	client  *http.Client
}

func NewClient(wechatPayCertPath, wechatPayKeyPath, rootCPath, mchAPIv3Key string, timeout time.Duration) (*Client, error) {
	//  Client{
	// 	wechatPayCertPath: wechatPayCertPath,
	// 	wechatPayKeyPath:  wechatPayKeyPath,
	// 	rootCPath:         rootCPath,
	// }
	var c Client
	c.SignKey = mchAPIv3Key
	certs, e := tls.LoadX509KeyPair(wechatPayCertPath, wechatPayKeyPath)
	if e != nil {
		return nil, e
	}

	rootCa, e := ioutil.ReadFile(rootCPath)
	if e != nil {
		return nil, e
	}
	pool := x509.NewCertPool()
	if !pool.AppendCertsFromPEM(rootCa) {
		// return nil, fmt.Errorf("AppendCertsFromPEM not ok")
	}

	c.client = &http.Client{
		Timeout: timeout,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
				RootCAs:            pool,
				Certificates:       []tls.Certificate{certs},
			},
		},
	}
	return &c, nil

}

func (c *Client) SendNormalRedPkg(req normal.Normal) error {
	var api = "https://api.mch.weixin.qq.com/mmpaymkttransfers/sendredpack"
	req.DoSign(c.SignKey)
	b, err := req.GetData()
	if err != nil {
		return err
	}

	request, err := http.NewRequest("POST", api, bytes.NewBuffer(b))
	if err != nil {
		fmt.Println(err)
	}
	res, err := c.client.Do(request)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	all, _ := io.ReadAll(res.Body)
	fmt.Println(string(all))
	return nil
}
