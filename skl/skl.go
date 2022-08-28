package skl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/hduLib/hdu/net"
	"io"
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

func (user *User) Get(data interface{}) error {
	var url string
	switch data.(type) {
	case *MyResp:
		url = MyURL
	case *UserInfoResp:
		url = UserInfoURL
	}
	req, err := user.newGet(url)
	if err != nil {
		return err
	}
	resp, err := net.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("status code is %d:%s", resp.StatusCode, string(body))
	}
	if err := json.Unmarshal(body, data); err != nil {
		return err
	}
	return nil
}

func (user *User) Post(data interface{}) ([]byte, error) {
	var url string
	switch data.(type) {
	case *PushReq:
		url = PushURL
	case *LeaveReq:
		url = LeaveURL
	}
	reqBody, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	req, err := user.newPost(url, reqBody)
	if err != nil {
		return nil, err
	}
	resp, err := net.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("status code is %d:%s", resp.StatusCode, string(resBody))
	}
	return resBody, err
}
