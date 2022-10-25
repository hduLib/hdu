package chaoxing

import (
	"bytes"
	"github.com/hduLib/hdu/net"
	"io"
	"net/http"
)

type Cx struct {
	cookie []*http.Cookie
}

func newUser(ck []*http.Cookie) *Cx {
	user := new(Cx)
	user.cookie = ck
	return user
}

func (cx *Cx) addCookieAndHeader2Req(req *http.Request) {
	for _, v := range cx.cookie {
		req.AddCookie(v)
	}
}

func (cx *Cx) newGet(url string) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	cx.addCookieAndHeader2Req(req)
	return req, nil
}

func (cx *Cx) newPost(url string, body []byte) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	cx.addCookieAndHeader2Req(req)
	return req, nil
}

func (cx *Cx) get(url string) ([]byte, error) {
	req, err := cx.newGet(url)
	if err != nil {
		return nil, err
	}
	resp, err := net.DefaultClient.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	return io.ReadAll(resp.Body)
}

func (cx *Cx) post(url string, data []byte) ([]byte, error) {
	req, err := cx.newPost(url, data)
	if err != nil {
		return nil, err
	}
	return net.Post(req)
}
