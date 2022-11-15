package phy

import (
	"os"
	"testing"

	"github.com/hduLib/hdu/internal/ocr"
)

func TestGetCaptchaContent(t *testing.T) {
	ocr.SetToken(os.Getenv("TOKEN"))
	_, err := getCaptchaContent()
	if err != nil {
		t.Fatal(err)
	}
}
