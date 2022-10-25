package chaoxing

import (
	"fmt"
	"time"
)

const (
	fanyaLoginURL = "http://passport2.chaoxing.com/fanyalogin"
)

func courseListURL() string {
	return fmt.Sprintf("http://mooc2-ans.chaoxing.com/mooc2-ans/visit/courses/list?v=%d&rss=1&start=0&size=500&catalogId=0&superstarClass=0&searchname=", time.Now().UnixMilli())
}
