package phy

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/hduLib/hdu/client"
)

type ExperSche struct {
	QueryWeekDayFlag int64        `json:"queryWeekDayFlag"`
	Experiments      []Experiment `json:"experiments"`
}

type Experiment struct {
	Selected bool   `json:"selected"`
	Value    int64  `json:"value"`
	Text     string `json:"text"`
}

// GetExperimentSche 返回所有实验安排
func (u *User) GetExperimentSche() (*ExperSche, error) {
	if u.SessionId == "" {
		return nil, ErrNoJSessionId
	}
	return u.getExperSches()
}

var (
	CourseId          = "325"
	SemesterNo        = "202220231"
	QueryExperimentId = "-1"
)

func (u *User) getExperSches() (*ExperSche, error) {
	payload := buildExprSchePayload(CourseId, SemesterNo, QueryExperimentId)

	req, err := http.NewRequest(http.MethodPost, "http://phy.hdu.edu.cn/phymember/v_mycourse_changed.jspx", strings.NewReader(payload))
	if err != nil {
		return nil, err
	}

	{
		req.Header.Set("authority", "phy.hdu.edu.cn")
		req.Header.Set("accept", "application/json, text/javascript, */*")
		req.Header.Set("accept-language", "en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7")
		req.Header.Set("content-type", "application/x-www-form-urlencoded")
		req.Header.Set("cookie", "clientlanguage=zh_CN; JSESSIONID="+u.SessionId)
		req.Header.Set("origin", "https://phy.hdu.edu.cn")
		req.Header.Set("referer", "https://phy.hdu.edu.cn/phymember/expt_schedule_student.jspx")
		req.Header.Set("sec-ch-ua", `"Google Chrome";v="107", "Chromium";v="107", "Not=A?Brand";v="24"`)
		req.Header.Set("sec-ch-ua-mobile", "?0")
		req.Header.Set("sec-ch-ua-platform", `"Windows"`)
		req.Header.Set("sec-fetch-dest", "empty")
		req.Header.Set("sec-fetch-mode", "cors")
		req.Header.Set("sec-fetch-site", "same-origin")
		req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36")
		req.Header.Set("x-requested-with", "XMLHttpRequest")
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	res := new(ExperSche)
	err = json.Unmarshal(b, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func buildExprSchePayload(courseId, semesterNo, queryExperimentId string) string {
	payload := make(url.Values, 3)
	payload.Add("queryCourseId", courseId)
	payload.Add("semesterNo", semesterNo)
	payload.Add("queryExperimentId", queryExperimentId)
	return payload.Encode()
}
