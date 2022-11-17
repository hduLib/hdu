package phy

import (
	"encoding/base64"
	"net/http"
	"net/url"
	"strings"

	"github.com/hduLib/hdu/client"
	"github.com/hduLib/hdu/internal/utils/convert"
)

var (
	IsLogined  = false
	JSessionId = ""
)

// Login 用于登录省物理实验平台，入参为学号和物理实验平台的密码
func Login(studentId, password string) error {
	// encode studentId, password
	encodedStuId := base64.StdEncoding.EncodeToString(convert.ToBytes(studentId))
	encodedPasswd := base64.StdEncoding.EncodeToString(convert.ToBytes(password))

	payload, err := buildLoginPayload(encodedStuId, encodedPasswd)
	if err != nil {
		return err
	}

	// send request
	err = requestWithPayload(payload)
	if err != nil {
		return err
	}

	IsLogined = true

	return nil
}

// payload: `returnUrl=%2F&username=<username>&password=<password>&captcha=<ocr_result>&x=0&y=0`
func buildLoginPayload(stuId, passwd string) (string, error) {
	payload := make(url.Values, 6)
	payload.Add("returnUrl", "/")
	payload.Add("username", stuId)
	payload.Add("password", passwd)
	captchaContent, err := getCaptchaContent()
	if err != nil {
		return "", err
	}
	payload.Add("captcha", captchaContent)
	payload.Add("x", "0")
	payload.Add("y", "0")
	return payload.Encode(), nil
}

func requestWithPayload(payload string) error {
	req, err := http.NewRequest(http.MethodPost, "http://phy.hdu.edu.cn/login.jspx", strings.NewReader(payload))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36")
	req.Header.Set("sec-ch-ua", `"Google Chrome";v="107", "Chromium";v="107", "Not=A?Brand";v="24"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	setJSessionId(resp)

	return nil
}

// set JSESSIONID
func setJSessionId(resp *http.Response) {
	for _, cookie := range resp.Cookies() {
		if cookie.Name != "JSESSIONID" {
			continue
		}
		JSessionId = cookie.Value
		break
	}
}
