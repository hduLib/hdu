package chaoxing

import (
	"fmt"
	"github.com/hduLib/hdu/chaoxing/course"
	"github.com/hduLib/hdu/chaoxing/request"
	"net/http"
)

type Cx struct {
	req *request.Request
}

func newUser(ck []*http.Cookie) *Cx {
	return &Cx{req: &request.Request{Cookies: ck}}
}

func (cx *Cx) CourseList() (*course.List, error) {
	resp, err := cx.req.Get(courseListURL())
	if err != nil {
		return nil, err
	}
	list, err := course.NewCourseList(resp, cx.req)
	if err != nil {
		return nil, fmt.Errorf("fail to parse courselist:%v", err)
	}
	return list, nil
}
