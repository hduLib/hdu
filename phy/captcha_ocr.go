package phy

import (
	"bytes"
	"io"
	"log"
	"net/http"

	"github.com/hduLib/hdu/client"
	"github.com/hduLib/hdu/internal/ocr"
)

func getCaptchaContent() string {
	req, _ := http.NewRequest(http.MethodGet, "http://phy.hdu.edu.cn/captcha.svl", nil)

	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%v", resp)

	// read in image
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// ocr
	rd := bytes.NewReader(b)
	captchaContent, err := ocr.Parse(rd)
	if err != nil {
		log.Fatal(err)
	}

	return captchaContent
}
