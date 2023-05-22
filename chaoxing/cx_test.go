package chaoxing

import (
	"fmt"
	"github.com/hduLib/hdu/client"
	"os"
	"testing"
)

var phone = os.Getenv("phone")
var passwd = os.Getenv("passwd")
var id = os.Getenv("id")
var casPasswd = os.Getenv("casPasswd")

func TestLogin(t *testing.T) {
	user, err := LoginWithPhoneAndPwd(phone, passwd)
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
		if err, ok := err.(*client.ErrNotOk); ok {
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
	user, err := LoginWithPhoneAndPwd(phone, passwd)
	if err != nil {
		t.Error(err)
		return
	}
	list, err := user.CourseList()
	if err != nil {
		t.Error(err)
		return
	}
	c, err := list.FindByName("计算机平面动画设计与制作").Detail()
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

func TestWork_Detail(t *testing.T) {
	user, err := LoginWithPhoneAndPwd(phone, passwd)
	if err != nil {
		t.Error(err)
		return
	}
	list, err := user.CourseList()
	if err != nil {
		t.Error(err)
		return
	}
	c, err := list.FindByName("创新实践B").Detail()
	if err != nil {
		t.Error(err)
		return
	}
	workList, err := c.WorkList()
	if err != nil {
		t.Error(err)
		return
	}
	wk := workList.Works[0].Detail()
	fmt.Println(wk)
}

func TestCourseChapter_NewList(t *testing.T) {
	user, err := LoginWithPhoneAndPwd(phone, passwd)
	if err != nil {
		t.Error(err)
		return
	}
	list, err := user.CourseList()
	if err != nil {
		t.Error(err)
		return
	}
	c, err := list.FindByName("计算机平面动画设计与制作").Detail()
	if err != nil {
		t.Error(err)
		return
	}
	chapter, err := c.ChapterList()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(chapter)
}

func TestCx_WorkList(t *testing.T) {
	user, err := LoginWithPhoneAndPwd(phone, passwd)
	if err != nil {
		t.Error(err)
		return
	}
	list, err := user.WorkList()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(list)
}
