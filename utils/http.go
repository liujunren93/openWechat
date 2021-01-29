package utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func HttpGet(url string) ([]byte, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	client := http.Client{}
	do, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	body := do.Body
	defer body.Close()
	re, err := ioutil.ReadAll(body)
	return re, err
}



func HttpPost(url string, header http.Header, data []byte) ([]byte, error) {

	url = strings.Trim(url, " ")
	newReader := bytes.NewReader(data)
	request, err := http.NewRequest("POST", url, newReader)
	if err != nil {
		fmt.Println(err)
	}
	request.Header = header
	client := http.Client{}
	do, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}
	body := do.Body
	all, err := ioutil.ReadAll(body)
	return all, err
}
