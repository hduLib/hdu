package request

import (
	"bytes"
	"fmt"
	"github.com/hduLib/hdu/net"
	"io"
	"net/http"
)

type Request struct {
	Cookies []*http.Cookie
}

func (r *Request) AddCookieAndHeader2Req(req *http.Request) {
	for _, v := range r.Cookies {
		req.AddCookie(v)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36 Edg/106.0.1370.52")
}

func (r *Request) NewGet(url string) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	r.AddCookieAndHeader2Req(req)
	return req, nil
}

func (r *Request) NewPost(url string, body []byte) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	r.AddCookieAndHeader2Req(req)
	return req, nil
}

func (r *Request) Get(url string) ([]byte, error) {
	req, err := r.NewGet(url)
	if err != nil {
		return nil, err
	}
	resp, err := net.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request fail:http status code is %d", resp.StatusCode)
	}
	return io.ReadAll(resp.Body)
}

func (r *Request) Post(url string, data []byte) ([]byte, error) {
	req, err := r.NewPost(url, data)
	if err != nil {
		return nil, err
	}
	return net.Post(req)
}
