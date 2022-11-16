package skl

import (
	"github.com/hduLib/hdu/skl/schema"
	"time"
)

// HasPush check if HasPushed on the day defined by t
// notice while push logs have multi pages,it only checks one.
func (r *PushLogResp) HasPush(t time.Time) bool {
	unix := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local).UnixMilli()
	for _, v := range r.List {
		if unix == v.CreateDate {
			return true
		}
	}
	return false
}

func (c *Course) DecodeSchema() (schema.Schema, error) {
	return schema.Decode(c.CourseSchema)
}
