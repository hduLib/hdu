package phy

import (
	"bytes"
	"io"

	"github.com/parnurzeal/gorequest"
)

func getCaptcha() io.Reader {
	agent := gorequest.New()

	agent.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	agent.Header.Set("Accept-Language", "en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7")
	agent.Header.Set("Cache-Control", "max-age=0")
	agent.Header.Set("Connection", "keep-alive")
	agent.Header.Set("Upgrade-Insecure-Requests", "1")
	agent.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36")

	_, captcha, _ := agent.Get("http://phy.hdu.edu.cn/captcha.svl").EndBytes()
	return bytes.NewReader(captcha)
}
