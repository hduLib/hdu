package skl

import (
	"fmt"
	"os"
	"testing"
	"time"
)

var (
	id     = os.Getenv("id")
	passwd = os.Getenv("casPasswd")
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
	t.Log(resp.Count)
	if err != nil {
		t.Error(err)
		return
	}
	if !resp.HasPush(time.Now()) {
		t.Error("not push today")
	}
	if resp.HasPush(time.Now().Add(time.Hour * 24)) {
		t.Error("why tomorrow pushed")
	}
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
	fmt.Println(resp)
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
	fmt.Println(resp.Week)
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
	fmt.Println(resp)
}
