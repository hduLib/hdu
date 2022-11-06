package ocr

import (
	"encoding/base64"
	"io"
	"log"
)

// RecognizeWithType takes an ocr type and an image reader, returns the string
// typed ocr result along with a possible error.
func RecognizeWithType(ocrType YunmaOCRType, image io.Reader) (string, error) {
	imgdata, err := io.ReadAll(image)
	if err != nil {
		log.Fatal(err)
	}
	imgbase64 := base64.StdEncoding.EncodeToString(imgdata)
	switch ocrType {
	case Common:
		return commonVerify(imgbase64)
	case Slide:
		return slideVerify(imgbase64)
	case SinSlide:
		return sinSlideVerify(imgbase64)
	case TrafficSlide:
		return trafficSlideVerify(imgbase64)
	case Click:
		return clickVerify(imgbase64)
	case Rotate:
		return rotateVerify(imgbase64)
	case Google:
		return googleVerify(imgbase64)
	case Hcaptcha:
		return hcaptchaVerify(imgbase64)
	case FunCaptcha:
		return funCaptchaVerify(imgbase64)
	}
	return "", ErrUnsupportOCRType
}

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
