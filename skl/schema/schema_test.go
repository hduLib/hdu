package schema

import (
	"testing"
	"time"
)

func TestSchemaDecode(t *testing.T) {
	s, err := Decode("星期三第6-7节{8-10周(双),14周};星期日第10-11节{12周}")
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Log(s)
	for _, v := range s {
		if v.Weekday == time.Wednesday {
			if !(v.WeekNum.Check(8) && v.WeekNum.Check(9) && !v.WeekNum.Check(10)) {
				t.Failed()
			}
		} else if v.Weekday == time.Sunday {
			if !v.WeekNum.Check(12) {
				t.Failed()
			}
		}
	}
	if !s.Check(6, 7, time.Wednesday, 8) {
		t.Failed()
	}
}
