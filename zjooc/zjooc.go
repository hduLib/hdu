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

func (user *User) newPost(url string, body []byte) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Add("openid", user.openid)
	return req, nil
}

func (user *User) newGet(url string) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("openid", user.openid)
	return req, nil
}

func (user *User) get(url string, data *interface{}) error {
	req, err := user.newGet(url)
	if err != nil {
		return err
	}
	return client.Get(req, data)
}

func (user *User) post(url string, data interface{}) ([]byte, error) {
	reqBody, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	req, err := user.newPost(url, reqBody)
	if err != nil {
		return nil, err
	}
	return client.Post(req)
}
