package chaoxing

import (
	"os"
	"testing"
)

var phone = os.Getenv("phone")
var passwd = os.Getenv("passwd")

func TestLogin(t *testing.T) {
	user, err := LoginByPhoneAndPwd(phone, passwd)
	if err != nil {
		t.Error(err)
		return
	}
	for _, v := range user.req.Cookies {
		t.Log(v.String())
	}
}

func TestCourse(t *testing.T) {
	user, err := LoginByPhoneAndPwd(phone, passwd)
	if err != nil {
		t.Error(err)
		return
	}
	for _, v := range user.req.Cookies {
		t.Logf("%s\n", v.String())
	}
	list, err := user.CourseList()
	if err != nil {
		t.Error(err)
		return
	}
	c, err := list.Courses[0].Detail()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(c)
}
