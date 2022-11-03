package course

import (
	"fmt"
	"time"
)

func (c *Course) workListURL() string {
	return fmt.Sprintf("https://mooc1.chaoxing.com/mooc2/work/list?courseId=%s&classId=%s&cpi=%s&ut=%s&enc=%s", c.CourseId, c.ClazzId, c.cpi, c.heardUt, c.workEnc)
}

func (c *Course) examListURL() string {
	return fmt.Sprintf("https://mooc1.chaoxing.com/exam-ans/mooc2/exam/exam-list?courseid=%s&clazzid=%s&cpi=%s&ut=%s&t=%s&enc=%s&openc=%s", c.CourseId, c.ClazzId, c.cpi, c.heardUt, c.t, c.enc, c.opEnc)
}

func (c *Course) chapterListURL() string {
	return fmt.Sprintf("https://mooc2-ans.chaoxing.com/mooc2-ans/mycourse/studentcourse?courseid=%s&clazzid=%s&cpi=%s&ut=%s&t=%d", c.CourseId, c.ClazzId, c.cpi, c.heardUt, time.Now().UnixMilli())
}
