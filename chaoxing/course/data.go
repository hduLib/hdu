package course

import "fmt"

func workListURL(c *Course) string {
	return fmt.Sprintf("https://mooc1.chaoxing.com/mooc2/work/list?courseId=%s&classId=%s&cpi=%s&ut=%s&enc=%s", c.CourseId, c.ClazzId, c.cpi, c.heardUt, c.workEnc)
}
