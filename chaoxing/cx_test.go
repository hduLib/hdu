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

func TestCourseAndExam(t *testing.T) {
	user, err := LoginByPhoneAndPwd(phone, passwd)
	if err != nil {
		t.Error(err)
		return
	}
	list, err := user.CourseList()
	if err != nil {
		t.Error(err)
		return
	}
	c, err := list.FindByName("数字电路设计").Detail()
	if err != nil {
		t.Error(err)
		return
	}
	workList, err := c.WorkList()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(workList)
	examList, err := c.ExamList()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(examList)
}
