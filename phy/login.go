package phy

import (
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/hduLib/hdu/client"
	"github.com/hduLib/hdu/internal/utils/convert"
)

var (
	isLogined  = false
	jSessionId = ""
)

// Login 用于登录省物理实验平台，入参为学号和物理实验平台的密码
func Login(studentId, password string) error {
	// encode studentId, password
	encodedStuId, encodedPasswd :=
		base64.StdEncoding.EncodeToString(convert.ToBytes(studentId)),
		base64.StdEncoding.EncodeToString(convert.ToBytes(password))

	// construct payload
	var builder strings.Builder
	builder.WriteString(`returnUrl=%2F&username=`)
	builder.WriteString(encodedStuId)
	builder.WriteString(`&password=`)
	builder.WriteString(encodedPasswd)
	builder.WriteString(`&captcha=`)
	builder.WriteString(getCaptchaContent())
	builder.WriteString(`&x=0&y=0`)
	payload := builder.String()

	// send request
	log.Println("sending login request with payload...")
	requestWithPayload(payload)
	log.Println("login success!")

	isLogined = true

	return nil
}

// payload: `returnUrl=%2F&username=<username>&password=<password>&captcha=<ocr_result>&x=0&y=0`
func requestWithPayload(payload string) {
	req, _ := http.NewRequest(http.MethodPost, "http://phy.hdu.edu.cn/login.jspx", strings.NewReader(payload))

	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Origin", "http://phy.hdu.edu.cn")
	req.Header.Set("Referer", "http://phy.hdu.edu.cn/login.jspx?returnUrl=/")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v", resp.Cookies())
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// handle cookies
	setJSessionId(convert.ToString(data))
}

// set JSESSIONID
func setJSessionId(resp string) {
	_ = resp // TODO: finish me
}
