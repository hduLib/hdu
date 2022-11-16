package schema

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type schemaNode struct {
	Begin, End int
	Weekday    time.Weekday
	WeekNum    weeks
}

type Schema []schemaNode

var schemaReg = regexp.MustCompile(`星期([一二三四五六日])第(\d)-(\d)节{(.*)}`)
var weekDayMapping = map[string]time.Weekday{
	"一": time.Monday,
	"二": time.Tuesday,
	"三": time.Wednesday,
	"四": time.Thursday,
	"五": time.Friday,
	"六": time.Saturday,
	"日": time.Sunday,
}

// Check if exist. especially, [begin,end] is a range , which means it returns true
// with input [1,12] if there is any course this day.
func (s Schema) Check(begin, end int, weekday time.Weekday, weekNum int) bool {
	for _, v := range s {
		if v.Begin <= begin && v.End >= end && v.Weekday == weekday && v.WeekNum.Check(weekNum) {
			return true
		}
	}
	return false
}

func Decode(str string) (Schema, error) {
	s := strings.Split(str, ";")
	schema := make(Schema, 0, len(s))
	for _, ss := range s {
		matches := schemaReg.FindAllStringSubmatch(ss, -1)
		if len(matches) != 1 {
			return nil, fmt.Errorf("ErrDecodingSchema:invalid schema node count")
		}
		info := matches[0]
		if len(info) != 5 {
			return nil, fmt.Errorf("ErrDecodingSchema:missing schemaNode info")
		}
		var n schemaNode
		n.Weekday = weekDayMapping[info[1]]
		var err error
		n.Begin, err = strconv.Atoi(info[2])
		if err != nil {
			return nil, fmt.Errorf("ErrDecodingSchema:%v", err)
		}
		n.End, err = strconv.Atoi(info[3])
		if err != nil {
			return nil, fmt.Errorf("ErrDecodingSchema:%v", err)
		}
		n.WeekNum, err = decodeWeeks(info[4])
		if err != nil {
			return nil, fmt.Errorf("ErrDecodingSchema:%v", err)
		}
		schema = append(schema, n)
	}
	return schema, nil
}

func decodeWeeks(str string) (weeks, error) {
	ss := strings.Split(str, ",")
	if len(ss) == 0 {
		return 0, nil
	}
	var w weeks
	for _, v := range ss {
		if strings.Contains(v, "-") {
			var begin, end int
			_, err := fmt.Sscanf(v, "%d-%d周", &begin, &end)
			if err != nil {
				return w, fmt.Errorf("ErrDecodingWeeks:%v", err)
			}
			for begin <= end {
				w |= 1 << begin
				begin++
			}
		} else {
			var t int
			_, err := fmt.Sscanf(v, "%d周", &t)
			if err != nil {
				return w, fmt.Errorf("errDecodingWeeks:%v", err)
			}
			w |= 1 << t
		}
	}
	return w, nil
}
