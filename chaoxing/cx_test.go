package chaoxing

import (
	"fmt"
	"github.com/hduLib/hdu/net"
	"os"
	"testing"
)

var phone = os.Getenv("phone")
var passwd = os.Getenv("passwd")
var id = os.Getenv("id")
var casPasswd = os.Getenv("casPasswd")

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

func TestLoginWithCas(t *testing.T) {
	user, err := LoginWithCas(id, casPasswd)
	if err != nil {
		if err, ok := err.(*net.ErrNotOk); ok {
			fmt.Println(err.Body)
		} else {
			t.Error(err)
		}

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
