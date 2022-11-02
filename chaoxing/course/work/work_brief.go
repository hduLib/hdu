package work

import "github.com/hduLib/hdu/chaoxing/request"

type Brief struct {
	url   string
	Title string
	//todo: 解析到time.Duration 或解析为截止时间time.Time
	Time string
	// 未交、已完成、待批阅
	Status string
	req    *request.Request
}
