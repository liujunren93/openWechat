package client

import (
	"bytes"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"time"
)

func do(r *http.Request, opts ...Option) (*http.Response, error) {
	client := http.DefaultClient
	var op options
	op.Timeout=time.Second*2
	for _, opt := range opts {
		opt(&op)
	}
	client.Transport = op.Transport
	r.Header = op.Header

	client.Timeout = op.Timeout
	return client.Do(r)

}

func HttpGet(url string, opts ...Option) ([]byte, error) {
	url = strings.Trim(url, " ")
	request, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}
	response, err := do(request, opts...)
	if err != nil {
		return nil, err
	}
	body := response.Body
	defer body.Close()
	re, err := ioutil.ReadAll(body)
	return re, err
}

func HttpPost(url string, data []byte, opts ...Option) ([]byte, error) {

	url = strings.Trim(url, " ")
	newReader := bytes.NewReader(data)
	request, err := http.NewRequest("POST", url, newReader)
	if err != nil {
		return nil, err
	}
	response, err := do(request, opts...)
	if err != nil {
		return nil, err
	}
	body := response.Body
	defer body.Close()
	all, err := ioutil.ReadAll(body)
	return all, err
}

type File interface {
	GetFieldName() string
	GetName() string
	GetData() *os.File
}

func HttpPostForm(url string, data map[string]string, files []File, opts ...Option) ([]byte, error) {

	buf := new(bytes.Buffer)
	writer := multipart.NewWriter(buf)
	for _, file := range files {
		if file.GetData() != nil {
			formFile, err := writer.CreateFormFile(file.GetFieldName(), file.GetName())
			if err != nil {
				return nil, err
			}
			_, err = io.Copy(formFile, file.GetData())
			if err != nil {
				return nil, err
			}
			opts = append(opts, WithHeader("content-type", writer.FormDataContentType()))
		}
	}

	for k, v := range data {
		err := writer.WriteField(k, v)
		if err != nil {
			return nil, err
		}
	}
	writer.Close()
	request, err := http.NewRequest("POST", url, buf)
	if err != nil {
		panic(err)
	}

	response, err := do(request, opts...)
	if err != nil {
		return nil, err
	}
	body := response.Body
	defer body.Close()
	return io.ReadAll(body)
}
