package sso

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/hduLib/hdu/client"
	"io"
	"net/http"
	"net/url"
	"regexp"
)

var keyRegexp = regexp.MustCompile("<p id=\"login-croypto\">(.*?)</p>")
var executionRegexp = regexp.MustCompile("<p id=\"login-page-flowkey\">(.*?)</p>")

func GenLoginReq(URL, user, passwd string) (*http.Request, error) {
	var key, execution []byte
	req, err := http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %v", err)
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36 Edg/124.0.0.0")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		reason, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("read body: %v", err)
		}
		return nil, fmt.Errorf("get key lt and excution: %s", string(reason))
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	tmp := keyRegexp.FindSubmatch(body)
	if len(tmp) != 2 {
		return nil, errors.New("match key")
	}
	key = tmp[1]
	tmp = executionRegexp.FindSubmatch(body)
	if len(tmp) != 2 {
		return nil, errors.New("match execution")
	}
	execution = tmp[1]
	bytes.Trim(key, " \"\r\n")

	//获取password
	encryptedPasswd, err := EncryptPasswd(key, passwd)
	if err != nil {
		return nil, fmt.Errorf("encrypt password: %v", err)
	}

	postData := url.Values{}
	postData.Set("username", user)
	postData.Set("passwordPre", passwd)
	postData.Set("password", encryptedPasswd)
	postData.Set("type", "UsernamePassword")
	postData.Set("_eventId", "submit")
	postData.Set("geolocation", "")
	postData.Set("execution", string(execution))
	// missing spelling from hdu, so what can I say?
	postData.Set("croypto", string(key))

	req, err = http.NewRequest(http.MethodPost, URL, bytes.NewBufferString(postData.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36 Edg/124.0.0.0")
	req.Header.Add("Referer", URL)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	for _, c := range resp.Cookies() {
		req.AddCookie(c)
	}

	var nextReq *http.Request
	c := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			nextReq = req
			return http.ErrUseLastResponse
		},
	}

	resp, err = c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %v", err)
	}

	if nextReq == nil || nextReq.URL.Hostname() == "sso.hdu.edu.cn" {
		return nil, errors.New("login failed")
	}

	return nextReq, nil
}
