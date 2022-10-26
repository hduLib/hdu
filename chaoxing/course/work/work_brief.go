package work

import "github.com/hduLib/hdu/chaoxing/request"

type Brief struct {
	url   string
	Title string
	Time  string
	// 未交、已完成、待批阅
	Status string
	req    *request.Request
}
