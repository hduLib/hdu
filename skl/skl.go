package skl

import (
	"bytes"
	"encoding/json"
	"github.com/hduLib/hdu/client"
	"net/http"
)

type User struct {
	xAuthToken string
}

func (user *User) addHeaderToReq(req *http.Request) {
	if req == nil {
		return
	}
	req.Header.Add("X-Auth-Token", user.xAuthToken)
	req.Header.Add("skl-Ticket", GenTicket())
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 16_0_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/20A380 AliApp(DingTalk/6.5.55) com.laiwang.DingTalk/26960829 Channel/201200 language/zh-Hans-CN UT4Aplus/0.0.6 WK")
	req.Header.Add("DingTalk-Flag", "1")
}

func (user *User) newGet(url string) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	user.addHeaderToReq(req)
	return req, nil
}

func (user *User) newPost(url string, body []byte) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	user.addHeaderToReq(req)
	return req, nil
}

func (user *User) get(url string, data interface{}) error {
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
