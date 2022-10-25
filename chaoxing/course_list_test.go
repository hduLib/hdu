package chaoxing

import (
	"testing"
)

func TestCx_CourseList(t *testing.T) {
	cx, err := loginByPhoneAndPwd(phone, passwd)
	if err != nil {
		t.Log(err)
	}
	list := cx.CourseList()
	t.Log(list)
}
