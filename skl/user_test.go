package skl

import (
	"testing"
	"time"
)

const (
	id     = "11111111"
	passwd = "11111111"
)

func TestUser_Push(t *testing.T) {
	skl, err := Login(id, passwd)
	if err != nil {
		t.Error(err)
		return
	}
	if err := skl.Push(&PushReqHDU); err != nil {
		t.Error(err)
		return
	}
}

func TestUser_PushLogs(t *testing.T) {
	skl, err := Login(id, passwd)
	if err != nil {
		t.Error(err)
		return
	}
	resp, err := skl.PushLogs()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(resp)
}

func TestUser_My(t *testing.T) {
	skl, err := Login(id, passwd)
	if err != nil {
		t.Error(err)
		return
	}
	resp, err := skl.My()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(resp)
}

func TestUser_Course(t *testing.T) {
	skl, err := Login(id, passwd)
	if err != nil {
		t.Error(err)
		return
	}
	resp, err := skl.Course(time.Now())
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(resp)
}

func TestUser_UserInfo(t *testing.T) {
	skl, err := Login(id, passwd)
	if err != nil {
		t.Error(err)
		return
	}
	resp, err := skl.UserInfo()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(resp)
}
