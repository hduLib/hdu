package phy

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/hduLib/hdu/client"
	"github.com/hduLib/hdu/internal/utils/convert"
)

type User struct {
	SessionId string
}

var (
	ErrBeforeLogin  = errors.New("before login")
	ErrNoJSessionId = errors.New("JSessionId is needed")
)

// Login 用于登录省物理实验平台，入参为学号和物理实验平台的密码
func Login(studentId, password string) (*User, error) {
	user := new(User)
	payload, err := user.buildLoginPayload(studentId, password)
	if err != nil {
		return nil, err
	}

	// send request
	err = user.requestWithPayload(payload)
	if err != nil {
		return nil, err
	}

	return user, err
}

func (u *User) requestWithPayload(payload string) error {
	req, err := http.NewRequest(http.MethodPost, "https://phy.hdu.edu.cn/login.jspx", strings.NewReader(payload))
	if err != nil {
		return err
	}

	{
		req.Header.Set("authority", "phy.hdu.edu.cn")
		req.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
		req.Header.Set("accept-language", "en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7")
		req.Header.Set("cache-control", "max-age=0")
		req.Header.Set("content-type", "application/x-www-form-urlencoded")
		req.Header.Set("cookie", "clientlanguage=zh_CN; JSESSIONID="+u.SessionId)
		req.Header.Set("origin", "https://phy.hdu.edu.cn")
		req.Header.Set("referer", "https://phy.hdu.edu.cn/login.jspx?returnUrl=/")
		req.Header.Set("sec-ch-ua", `"Google Chrome";v="107", "Chromium";v="107", "Not=A?Brand";v="24"`)
		req.Header.Set("sec-ch-ua-mobile", "?0")
		req.Header.Set("sec-ch-ua-platform", `"Windows"`)
		req.Header.Set("sec-fetch-dest", "document")
		req.Header.Set("sec-fetch-mode", "navigate")
		req.Header.Set("sec-fetch-site", "same-origin")
		req.Header.Set("sec-fetch-user", "?1")
		req.Header.Set("upgrade-insecure-requests", "1")
		req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36")
	}

	// send login request
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// check login status
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = checkLoginStatus(bodyText)
	if err != nil {
		return err
	}

	// refresh JSessionId
	u.setJSessionId(resp.Cookies())

	return nil
}

// payload: `returnUrl=%2F&username=<username>&password=<password>&captcha=<ocr_result>&x=0&y=0`
func (u *User) buildLoginPayload(stuId, passwd string) (string, error) {
	payload := make(url.Values, 6)
	payload.Add("returnUrl", "/")
	payload.Add("username", stuId)
	payload.Add("password", passwd)
	captchaContent, err := getCaptchaContent(u)
	if err != nil {
		return "", err
	}
	payload.Add("captcha", captchaContent)
	payload.Add("x", "0")
	payload.Add("y", "0")
	return payload.Encode(), nil
}

func checkLoginStatus(respBody []byte) error {
	if bytes.Contains(respBody, convert.ToBytes("您还没有登录")) {
		return ErrBeforeLogin
	}
	return nil
}
