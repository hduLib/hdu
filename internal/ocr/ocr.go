package ocr

import (
	"encoding/base64"
	"errors"
	"io"
	"log"

	"github.com/hduLib/hdu/internal/utils/convert"
)

var DefaultType = Common

func Parse(captcha interface{}) (string, error) {
	return recognizeWithType(DefaultType, captcha)
}

// recognizeWithType takes an ocr type and an image reader, returns the string
// typed ocr result along with a possible error.
func recognizeWithType(ocrType YunmaOCRType, captcha interface{}) (string, error) {
	var bs64captcha string

	// base64 encoded
	switch captcha := captcha.(type) {
	case []byte:
		bs64captcha = base64.StdEncoding.EncodeToString(captcha)
	case string:
		bs64captcha = base64.StdEncoding.EncodeToString(convert.ToBytes(captcha))
	case io.Reader:
		data, err := io.ReadAll(captcha)
		if err != nil {
			log.Fatal(err)
		}
		bs64captcha = base64.StdEncoding.EncodeToString(data)
	default:
		return "", ErrUnsupportCaptchaType
	}

	// do ocr
	switch ocrType {
	case Common:
		return commonVerify(bs64captcha)
	case Slide:
		return slideVerify(bs64captcha)
	case SinSlide:
		return sinSlideVerify(bs64captcha)
	case TrafficSlide:
		return trafficSlideVerify(bs64captcha)
	case Click:
		return clickVerify(bs64captcha)
	case Rotate:
		return rotateVerify(bs64captcha)
	case Google:
		return googleVerify(bs64captcha)
	case Hcaptcha:
		return hcaptchaVerify(bs64captcha)
	case FunCaptcha:
		return funCaptchaVerify(bs64captcha)
	}
	return "", ErrUnsupportOCRType
}

var (
	ErrUnsupportOCRType     = errors.New("ocr type is unsupported")
	ErrUnsupportCaptchaType = errors.New("ocr parse arg type is unsupported")
)

type YunmaOCRType int

//go:generate stringer -type=YunmaOCRType
const (
	Common YunmaOCRType = iota + 1
	Slide
	SinSlide
	TrafficSlide
	Click
	Rotate
	Google
	Hcaptcha
	FunCaptcha
)
