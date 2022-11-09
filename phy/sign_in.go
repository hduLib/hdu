package phy

import (
	"encoding/base64"
	"github.com/hduLib/hdu/internal/ocr"
	"log"
	"strings"

	"github.com/hduLib/hdu/utils/convert"
	"github.com/parnurzeal/gorequest"
)

var (
	isSignedIn = false
	jSessionId = ""
)

func SignIn(username, password string) error {
	// encode username, password
	encodedUname, encodedPasswd :=
		base64.StdEncoding.EncodeToString(convert.ToBytes(username)),
		base64.StdEncoding.EncodeToString(convert.ToBytes(password))

	// construct payload
	var builder strings.Builder
	builder.WriteString(`returnUrl=%2F&username=`)
	builder.WriteString(encodedUname)
	builder.WriteString(`&password=`)
	builder.WriteString(encodedPasswd)
	builder.WriteString(`&captcha=`)
	// parse captcha
	captchaRes, err := ocr.RecognizeWithType(ocr.Common, getCaptcha())
	if err != nil {
		return err
	}
	builder.WriteString(captchaRes)
	builder.WriteString(`&x=0&y=0`)
	payload := builder.String()

	// send request
	log.Println("sending login request with payload...")
	requestWithPayload(payload)
	log.Println("login success!")

	isSignedIn = true

	return nil
}

// payload: `returnUrl=%2F&username=<username>&password=<password>&captcha=<ocr_result>&x=0&y=0`
func requestWithPayload(payload string) {
	agent := gorequest.New()

	agent.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	agent.Header.Set("Accept-Language", "en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7")
	agent.Header.Set("Cache-Control", "max-age=0")
	agent.Header.Set("Connection", "keep-alive")
	agent.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	agent.Header.Set("Origin", "http://phy.hdu.edu.cn")
	agent.Header.Set("Referer", "http://phy.hdu.edu.cn/login.jspx?returnUrl=/")
	agent.Header.Set("Upgrade-Insecure-Requests", "1")
	agent.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36")

	resp, _, errs :=
		agent.Post("http://phy.hdu.edu.cn/login.jspx").SendString(payload).End()
	if len(errs) != 0 {
		log.Fatal(errs)
	}

	// handle cookies
	cookies := resp.Header["Set-Cookie"]
	setJSessionId(cookies[1])
}

// set JSESSIONID
func setJSessionId(cookie string) {
	start, end := strings.Index(cookie, "="), strings.Index(cookie, ";")
	jSessionId = cookie[start+1 : end]
}
