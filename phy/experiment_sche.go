package phy

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
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

var (
	ErrNoJSessionId = errors.New("JSessionId is needed")
)

// GetExperimentSche 返回所有实验安排
func GetExperimentSche() (*ExperSche, error) {
	if JSessionId == "" {
		return nil, ErrNoJSessionId
	}
	return getExperSches()
}

var (
	CourseId   = "325"
	SemesterNo = "202220231"
)

func getExperSches() (*ExperSche, error) {
	data := strings.NewReader("queryCourseId=" + CourseId + "&semesterNo=" + SemesterNo + "&queryExperimentId=-1")
	req, err := http.NewRequest("POST", "http://phy.hdu.edu.cn/phymember/v_mycourse_changed.jspx", data)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json, text/javascript, */*")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "clientlanguage=zh_CN; JSESSIONID=FA33E279A1BF84C0AC18F22EE0DAF6B3")
	req.Header.Set("Origin", "http://phy.hdu.edu.cn")
	req.Header.Set("Referer", "http://phy.hdu.edu.cn/phymember/expt_schedule_student.jspx")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")

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
