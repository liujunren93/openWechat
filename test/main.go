package main

import (
	"encoding/xml"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
)

const (
	appid         = "wxa3c6a833afddef99"
	token         = "sharelife"
	key           = "hCs6q86iQNeEzyZ6ijK3mRruYcB4NMncNVj42JdhXTj"
	nonce         = "1614133472"
	timestamp     = "1658733735"
	msg_signature = "05fe659c06dfc25b0657a6486fff30e2f9f208de"
	signature     = "620e1c0922db6c206ac176ccd077244873b146f7"
)

var data = "+5yJXCLAM5w8pQF1DHYykFOc3QoEQGOxC+EubEo/E4Luyy1K6IEH4PLLfmXyVhTUrh27uTz/nfIYpYY49F5BA/UrSjhQ/lsctBGQ4N7Vd1PQjk6P/Zd5ERMzmOmzclDF2rU2jZnHy0Mlgdae8kamyT00cLUaheEtYhER7iGrhanV8KItIvsQEHzxj2NM1luqcBIrxw+RkPOPa6CCJ3e9ECpKd6G4bkmQKXHlZ53ILqNg+uRLE1w6UclAKyy8FNWzgoWWdopB8+6R1mQdp/WSJMdqK9yT7eZNPqbRxy3Jrm7Ls7R0vKAAjVlPqmB7WPi5HnQJTSiW05srp63e7eQs3SFlUe5RUo9aI2BY61Jk5IJsyJtGIiiqKVXAJy/sfbihMWby5fW8gmlNdiWuZ3+q9dZdY0vx4eEGXG7aTknU0xhGcjeYtld5rIuwv9CFdkGka7oIFSXrJbu0BiETefC0Tw=="

// https://mp.weixin.qq.com/cgi-bin/componentloginpage?component_appid=wxa3c6a833afddef99&pre_auth_code=preauthcode@@@jkDJ0NYWEJAFLghi-Ze68tRUgj9b3D3Lvi2u2NMLvsQ6Zv1tSRs8QHiP2zdxou6hurA5jeJ0e_ujOQXshS-dDw&redirect_uri=https://xq2ffb.natappfree.cc/gh_5500d3cbd1fe/callback

type componentTicket struct {
	AppId   string `xml:"AppId"`
	Encrypt string `xml:"Encrypt"`
}

func main() {
	http.HandleFunc("/auth", func(w http.ResponseWriter, r *http.Request) {
		data, err := io.ReadAll(r.Body)
		defer r.Body.Close()
		var ct componentTicket
		xml.Unmarshal(data, &ct)
		fmt.Println(ct, err, r.URL.Query())
	})
	http.HandleFunc("/gh_5500d3cbd1fe/callback", func(w http.ResponseWriter, r *http.Request) {
		data, err := io.ReadAll(r.Body)
		defer r.Body.Close()
		var ct componentTicket
		xml.Unmarshal(data, &ct)
		fmt.Println(ct, err, r.URL.Query())
	})
	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("./index.html")
		if err != nil {
			log.Println("err:", err)
			return
		}
		t.Execute(w, nil)

	})
	http.ListenAndServe(":19090", nil)

	// r, dd, err := utils.DecryptMsg(appid, data, key)
	// fmt.Println(string(r), string(dd), err)
}
