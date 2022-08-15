package skl

import (
	"errors"
	"github.com/hduLib/hdu/cas"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
)

const casLogin = "https://skl.hdu.edu.cn/api/userinfo?type=&index=passcard.html"

func Login(id, password string) (*User, error) {
	resp, err := http.Get(casLogin)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	url := gjson.Get(string(body), "url").String()
	req, err := cas.GenLoginReq(url, id, password)
	if err != nil {
		return nil, err
	}
	XAuthToken := ""
	c := &http.Client{CheckRedirect: func(req *http.Request, via []*http.Request) error {
		token := req.Response.Header.Get("X-Auth-Token")
		if token != "" {
			XAuthToken = token
		}
		return nil
	}}
	resp, err = c.Do(req)
	if err != nil {
		return nil, err
	}
	if XAuthToken == "" {
		return nil, errors.New("fail to get xauthtoken")
	}
	return &User{
		xAuthToken: XAuthToken,
	}, nil
}
