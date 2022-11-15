package phy

import (
	"testing"
)

func TestGetCaptchaContent(t *testing.T) {
	_, err := getCaptchaContent()
	if err != nil {
		t.Fatal(err)
	}
}
