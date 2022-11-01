package cas

import (
	"bytes"
	_ "embed"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

var ltRegexp = regexp.MustCompile("<input type=\"hidden\" id=\"lt\" name=\"lt\" value=\"(.*)\" />")
var executionRegexp = regexp.MustCompile("<input type=\"hidden\" name=\"execution\" value=\"(.*)\" />")

func GenLoginReq(URL, user, passwd string) (*http.Request, error) {
	var lt, execution []byte

	//获取lt和execution
	req, err := http.NewRequest(http.MethodGet, URL, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.5112.81 Safari/537.36")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		reason, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("fail to read body: %v", err)
		}
		return nil, fmt.Errorf("fail to get lt and excution: %s", string(reason))
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	tmp := ltRegexp.FindSubmatch(body)
	if len(tmp) != 2 {
		return nil, errors.New("提取lt错误")
	}
	lt = tmp[1]
	tmp = executionRegexp.FindSubmatch(body)
	if len(tmp) != 2 {
		return nil, errors.New("提取execution错误")
	}
	execution = tmp[1]

	//获取rsa
	rsa := getRsa(user + passwd + string(lt))

	postData := url.Values{}
	postData.Add("rsa", rsa)
	postData.Add("ul", strconv.Itoa(len(user)))
	postData.Add("pl", strconv.Itoa(len(passwd)))
	postData.Add("lt", string(lt))
	postData.Add("execution", string(execution))
	postData.Add("_eventId", "submit")
	req, err = http.NewRequest(http.MethodPost, URL, bytes.NewBufferString(postData.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.5112.81 Safari/537.36")
	req.Header.Add("Referer", URL)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	for _, c := range resp.Cookies() {
		req.AddCookie(c)
	}
	if err != nil {
		return nil, err
	}
	return req, nil
}
