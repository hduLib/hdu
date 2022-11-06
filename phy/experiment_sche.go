package phy

import (
	"github.com/parnurzeal/gorequest"
)

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

// TODO: finish me
func GetExperimentSche() {
}

func getContent() {
	agent := gorequest.New()

	agent.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	agent.Header.Set("Accept-Language", "en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7")
	agent.Header.Set("Cache-Control", "max-age=0")
	agent.Header.Set("Connection", "keep-alive")
	agent.Header.Set("Cookie", "clientlanguage=zh_CN; " /*TODO: set JSESSIONID*/)
	agent.Header.Set("Referer", "http://phy.hdu.edu.cn/phymember/expt_schedule_student.jspx")
	agent.Header.Set("Upgrade-Insecure-Requests", "1")
	agent.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36")

	agent.Get("http://phy.hdu.edu.cn/phymember/expt_schedule_student.jspx").EndBytes()
}
