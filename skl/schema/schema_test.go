package schema

import (
	"testing"
	"time"
)

func TestSchemaDecode(t *testing.T) {
	s, err := Decode("星期一第3-5节{1-5周,7-17周};星期六第3-5节{6周}")
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Log(s)
	for _, v := range s {
		if v.Weekday == time.Monday {
			if !(v.WeekNum.Check(1) && v.WeekNum.Check(2) && !v.WeekNum.Check(6) && v.WeekNum.Check(17)) {
				t.Failed()
			}
		} else if v.Weekday == time.Saturday {
			if !(v.WeekNum.Check(6) && !v.WeekNum.Check(7)) {
				t.Failed()
			}
		}
	}
	if !s.Check(3, 5, time.Monday, 5) {
		t.Failed()
	}
	if s.Check(3, 5, time.Tuesday, 5) {
		t.Failed()
	}
	if !s.Check(1, 12, time.Monday, 5) {
		t.Failed()
	}
}
