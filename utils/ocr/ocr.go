package ocr

import (
	"encoding/base64"
	"errors"
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
	case SinSlide:
	case TrafficSlide:
	case Click:
	case Rotate:
	case Google:
	case Hcaptcha:
	case FunCaptcha:
	}
	return "", errors.New("unsupported ocr type")
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