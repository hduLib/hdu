package zjooc

import (
	"bytes"
	"encoding/json"
	"github.com/hduLib/hdu/internal/client"
	"net/http"
)

type User struct {
	openid string
}

func (u *User) addHeaderToReq(req *http.Request) {
	if req == nil {
		return
	}
	req.Header.Add("openid", u.openid)
	req.Header.Set("User-Agent", "%E5%9C%A8%E6%B5%99%E5%AD%A6/2 CFNetwork/1390 Darwin/22.0.0")
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
}

func (u *User) newPost(url string, body []byte) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	u.addHeaderToReq(req)
	return req, nil
}

func (u *User) newGet(url string) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	u.addHeaderToReq(req)
	return req, nil
}

func (u *User) get(url string, data interface{}) error {
	req, err := u.newGet(url)
	if err != nil {
		return err
	}
	return client.Get(req, data)
}

func (u *User) post(url string, data interface{}) ([]byte, error) {
	reqBody, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	req, err := u.newPost(url, reqBody)
	if err != nil {
		return nil, err
	}
	return client.Post(req)
}
