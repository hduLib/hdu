package sso

import (
	"testing"
)

func TestLogin(t *testing.T) {
	req, err := GenLoginReq("https://sso.hdu.edu.cn/login?service=https%3A%2F%2Fi.hdu.edu.cn%2Fsopcb%2F", "21111111", "11111111")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(req)
}
