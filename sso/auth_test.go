package sso

import (
	"testing"
)

func TestLogin(t *testing.T) {
	cookies, err := Login("https://sso.hdu.edu.cn/login?service=https%3A%2F%2Fi.hdu.edu.cn%2Fsopcb%2F", "211111111", "123123123")
	if err != nil {
		t.Error(err)
	}
	t.Log(cookies)
}
