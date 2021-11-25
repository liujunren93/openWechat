package pay

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"github.com/liujunren93/openWechat/client"
	"io/ioutil"
	"net/http"
)

type Client struct {
	PayCert  string
	PayKey   string
	RootCert string
	SignKey  string
}
type Payer interface {
	GetApiPath() string
	GetData() ([]byte, error)
	DoSign(key string) error
}

func (c Client) Do(p Payer) ([]byte, error) {
	certs, err := tls.LoadX509KeyPair(c.PayCert, c.PayKey)
	if err != nil {
		return nil, err
	}

	rootCa, e := ioutil.ReadFile(c.RootCert)
	if e != nil {
		return nil, err
	}
	pool := x509.NewCertPool()
	if pool.AppendCertsFromPEM(rootCa) {
		return nil, errors.New("AppendCertsFromPEM err")
	}
	t := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
			RootCAs:            pool,
			Certificates:       []tls.Certificate{certs},
		},
	}
	err = p.DoSign(c.SignKey)
	if err != nil {
		return nil, err
	}
	data, err := p.GetData()
	if err != nil {
		return nil, err
	}
	return client.HttpPost(p.GetApiPath(), data, client.WithTransport(t))

}
