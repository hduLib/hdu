package course

import (
	"github.com/hduLib/hdu/chaoxing/request"
)

type Brief struct {
	ClazzId  string
	CourseId string
	CoverURL string
	Title    string
	// Duration 意义不大，不同老师填的不同，不推荐使用
	Duration    string
	TeacherName string
	// CourseNum 可能是上课地点可能是课程号，感觉全看教师自己填了什么，不推荐使用
	CourseNum string
	url       string
	// point to cx for getting further information
	req *request.Request
	//本来应该有一个名为cpi的字段，但是仅仅出现在url内并且没有摸清他的意义，暂时不予理会
}

// Detail returns detailed Course for further request
func (br *Brief) Detail() (*Course, error) {
	resp, err := br.req.Get(br.url)
	if err != nil {
		return nil, err
	}
	return NewCourse(resp, br)
}
