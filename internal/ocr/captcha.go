//go:build export
package ocr

import (
	"bytes"
	"io"
	"log"
	"net/http"

	"github.com/hduLib/hdu/client"
)

func setHeaders(req *http.Request, url string) {
	// TODO: set headers
}

// -*- concurrent safe -*-

func (h *HandlerSync) ReadCaptchaFrom(url string) *HandlerSync {
	h.mu.Lock()
	defer h.mu.Unlock()

	req, _ := http.NewRequest(http.MethodGet, url, nil)

	setHeaders(req, url)

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

	h.captcha = b
	return h
}

func (h *HandlerSync) Parse() (string, error) {
	h.mu.Lock()
	defer h.mu.Unlock()
	captcha := bytes.NewReader(h.captcha)
	return RecognizeWithType(h.typ, captcha)
}

// -*- concurrent unsafe -*-

func (h *Handler) ReadCaptchaFrom(url string) *Handler {
	req, _ := http.NewRequest(http.MethodGet, url, nil)

	setHeaders(req, url)

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

	h.captcha = b
	return h
}

func (h *Handler) Parse() (string, error) {
	captcha := bytes.NewReader(h.captcha)
	return RecognizeWithType(h.typ, captcha)
}
