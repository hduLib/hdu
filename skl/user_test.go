package skl

import (
	"testing"
)

const (
	id     = "21111111"
	passwd = "11111111"
)

func TestUser_Push(t *testing.T) {
	skl, err := Login(id, passwd)
	if err != nil {
		t.Error(err)
		return
	}
	if err := skl.Push(&pushReqHDU); err != nil {
		t.Error(err)
		return
	}
}

func TestUser_My(t *testing.T) {
	skl, err := Login(id, passwd)
	if err != nil {
		t.Error(err)
		return
	}
	payload := new(MyResp)
	if err := skl.My(payload); err != nil {
		t.Error(err)
		return
	}
	t.Log(payload)
}

func TestUser_UserInfo(t *testing.T) {
	skl, err := Login(id, passwd)
	if err != nil {
		t.Error(err)
		return
	}
	payload := new(UserInfoResp)
	if err := skl.UserInfo(payload); err != nil {
		t.Error(err)
		return
	}
	t.Log(payload)
}
