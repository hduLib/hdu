package phy

import (
	"errors"
	"io"
	"net/http"

	"github.com/hduLib/hdu/client"
)

// ExperSche 具体实验安排
type ExperSche struct {
	Id          int    // ID
	ExperName   string // 实验名称
	Teacher     string // 实验教师
	GroupId     int    // 实验组号
	Location    string // 实验地址
	TimeSection string // 起止时间
	SelectedMax string // 已选 / 最大
	Grade       int    // 成绩
}

var (
	ErrNoLogin = errors.New("you are not logined before get experiments schedules")
)

// GetExperimentSche 返回所有实验安排
func GetExperimentSche() ([]*ExperSche, error) {
	if !isLogined || JSessionId == "" {
		return nil, ErrNoLogin
	}
	return getExperimentSche()
}

func getExperimentSche() ([]*ExperSche, error) {
	req, err := http.NewRequest(http.MethodGet, "http://phy.hdu.edu.cn/phymember/expt_schedule_student.jspx", nil)
	if err != nil {
		return nil, err
	}

	{
		req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
		req.Header.Set("Accept-Language", "en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7")
		req.Header.Set("Cache-Control", "max-age=0")
		req.Header.Set("Connection", "keep-alive")
		req.Header.Set("Cookie", "clientlanguage=zh_CN; "+"JSESSIONID="+JSessionId)
		req.Header.Set("Referer", "http://phy.hdu.edu.cn/phymember/expt_schedule_student.jspx")
		req.Header.Set("Upgrade-Insecure-Requests", "1")
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36")
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return deserialize(b), nil
}

// TODO: finish me
func deserialize(b []byte) []*ExperSche {
	return nil
}
