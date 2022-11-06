package phy

import (
	"os"
	"testing"
)

const Kib = 1024

func TestGetCaptcha(t *testing.T) {
	buf := make([]byte, 12*Kib)
	n, _ := getCaptcha().Read(buf)
	buf = buf[:n]
	err := os.WriteFile("test.svl", buf, 0655)
	if err != nil {
		t.FailNow()
	}
}
