package phy

import (
	"testing"
)

func TestGetCaptchaContent(t *testing.T) {
	if getCaptchaContent() == "" {
		t.FailNow()
	}
}
