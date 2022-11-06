package phy

import (
	"bytes"
	"io"
	"log"
	"net/http"
)

func getCaptcha() io.Reader {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://phy.hdu.edu.cn/captcha.svl", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cookie", "clientlanguage=zh_CN; JSESSIONID=66BAAEB5CDA440199BC160EE76CAE8B0")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	captcha, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return bytes.NewReader(captcha)
}
