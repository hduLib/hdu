package chaoxing

import (
	"fmt"
	"github.com/hduLib/hdu/chaoxing/course"
	"github.com/hduLib/hdu/chaoxing/course/work"
	"github.com/hduLib/hdu/chaoxing/request"
	"github.com/hduLib/hdu/client"
	"io"
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

func todoListURL() string {
	return fmt.Sprintf("https://home-yd.chaoxing.com/proxy/gopage?cxanalyzetag=hp&type=mywork")
}

// WorkList 是手机api的作业列表
func (cx *Cx) WorkList() (*work.List, error) {
	req, err := cx.req.NewGet(todoListURL())
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 16_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 (schild:0adc7adb7f466d4df73716e19c87efaa) (device:iPhone14,5) Language/zh-Hans com.ssreader.ChaoXingStudy/ChaoXingStudy_3_6.0.6_ios_phone_202304130930_102 (@Kalimdor)_9407410973156787895")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return nil, &client.ErrNotOk{
			StatusCode: resp.StatusCode,
			Body:       string(body),
		}
	}
	return work.NewListPhoneAPI(body, cx.req)
}
