package exam

import "github.com/hduLib/hdu/chaoxing/request"

const (
	Undo     = "待做"
	Finished = "已完成"
)

type Brief struct {
	url   string
	Title string
	//todo: 解析到time.Duration 或解析为截止时间time.Time
	Time string
	//待做、已完成
	Status string
	req    *request.Request
}

func (b *Brief) Open() (*Exam, error) {
	//todo: open a exam
	return nil, nil
}
