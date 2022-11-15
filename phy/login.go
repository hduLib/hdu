package phy

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/hduLib/hdu/client"
	"github.com/hduLib/hdu/internal/utils/convert"
)

var (
	isLogined  = false
	JSessionId = ""
)

// Login 用于登录省物理实验平台，入参为学号和物理实验平台的密码
func Login(studentId, password string) error {
	// encode studentId, password
	encodedStuId, encodedPasswd :=
		base64.StdEncoding.EncodeToString(convert.ToBytes(studentId)),
		base64.StdEncoding.EncodeToString(convert.ToBytes(password))

	payload, err := buildPayload(encodedStuId, encodedPasswd)
	if err != nil {
		return err
	}

	// send request
	err = requestWithPayload(payload)
	if err != nil {
		return err
	}

	isLogined, err = getLoginStatus()
	if err != nil {
		return err
	}

	return nil
}

func buildPayload(stuId, passwd string) (string, error) {
	var builder strings.Builder
	builder.WriteString(`returnUrl=%2F&username=`)
	builder.WriteString(stuId)
	builder.WriteString(`&password=`)
	builder.WriteString(passwd)
	builder.WriteString(`&captcha=`)
	captchaContent, err := getCaptchaContent()
	if err != nil {
		return "", err
	}
	builder.WriteString(captchaContent)
	builder.WriteString(`&x=0&y=0`)
	return builder.String(), nil
}

type LoginResult struct {
	Result bool  `json:"result"`
	Count  int64 `json:"count"`
}

func getLoginStatus() (bool, error) {
	if JSessionId == "" {
		return false, nil
	}

	req, err := http.NewRequest("POST", "http://phy.hdu.edu.cn/phymember/message_countUnreadMsg.jspx", nil)
	if err != nil {
		return false, err
	}

	{
		req.Header.Set("Accept", "application/json, text/javascript, */*")
		req.Header.Set("Accept-Encoding", "gzip, deflate, br")
		req.Header.Set("Accept-Language", "en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7")
		req.Header.Set("Connection", "keep-alive")
		req.Header.Set("Cache-Control", "max-age=0")
		req.Header.Set("Content-langth", "0")
		req.Header.Set("Content-Type", "text/plain;charset=UTF-8")
		req.Header.Set("Cookie", "clientlanguage=zh_CN; "+"JSESSIONID="+JSessionId)
		req.Header.Set("Host", "phy.hdu.edu.cn")
		req.Header.Set("Origin", "https://phy.hdu.edu.cn")
		req.Header.Set("Referer", "https://phy.hdu.edu.cn/phymember/index.jspx?locale=zh_CN")
		req.Header.Set("sec-ch-ua", `"Google Chrome";v="107", "Chromium";v="107", "Not=A?Brand";v="24"`)
		req.Header.Set("sec-ch-ua-mobile", "?0")
		req.Header.Set("sec-ch-ua-platform", `"Windows"`)
		req.Header.Set("Sec-Fetch-Dest", "empty")
		req.Header.Set("Sec-Fetch-Mode", "cors")
		req.Header.Set("Sec-Fetch-Site", "same-origin")
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36")
		req.Header.Set("X-Requested-With", "XMLHttpRequest")
	}
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	// check login result
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}
	loginResult := new(LoginResult)
	err = json.Unmarshal(data, loginResult)
	if err != nil {
		return false, err
	}

	return loginResult.Result, nil
}

// payload: `returnUrl=%2F&username=<username>&password=<password>&captcha=<ocr_result>&x=0&y=0`
func requestWithPayload(payload string) error {
	req, err := http.NewRequest(http.MethodPost, "http://phy.hdu.edu.cn/login.jspx", strings.NewReader(payload))
	if err != nil {
		return err
	}

	{
		req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*;q=0.8,application/signed-exchange;v=b3;q=0.9")
		req.Header.Set("Accept-Language", "en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7")
		req.Header.Set("Cache-Control", "max-age=0")
		req.Header.Set("Connection", "keep-alive")
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Set("Origin", "https://phy.hdu.edu.cn")
		req.Header.Set("Referer", "https://phy.hdu.edu.cn/login.jspx")
		req.Header.Set("Sec-Fetch-Dest", "document")
		req.Header.Set("Sec-Fetch-Mode", "navigate")
		req.Header.Set("Sec-Fetch-Site", "same-origin")
		req.Header.Set("Sec-Fetch-User", "?1")
		req.Header.Set("Upgrade-Insecure-Requests", "1")
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36")
		req.Header.Set("sec-ch-ua", `"Google Chrome";v="107", "Chromium";v="107", "Not=A?Brand";v="24"`)
		req.Header.Set("sec-ch-ua-mobile", "?0")
		req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	}

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
