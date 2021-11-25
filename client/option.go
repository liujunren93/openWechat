package client

import (
	"net/http"
	"time"
)

type options struct {
	Header   http.Header
	Transport *http.Transport
	Timeout time.Duration
}

type Option func(*options)

func WithHeader(k,v string) Option {
	return func(o *options) {
		if o.Header==nil {
			o.Header=http.Header{}
		}
		o.Header.Set(k,v)
	}
}

func WithTransport(transport *http.Transport) Option {
	return func(o *options) {
		o.Transport = transport
	}
}
