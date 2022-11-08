package phy

import (
	"os"
	"testing"
)

func TestGetCaptcha(t *testing.T) {
	buf := make([]byte, 2*4096)
	n, _ := getCaptcha().Read(buf)
	buf = buf[:n]
	err := os.WriteFile("captcha.svl", buf, 0655)
	if err != nil {
		t.FailNow()
	}
}
