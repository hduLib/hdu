package exam

import (
	"github.com/hduLib/hdu/chaoxing/request"
	"time"
)

const (
	Undo     = "待做"
	Finished = "已完成"
)

type Brief struct {
	url   string
	Title string
	// 根据剩余时间推断，可能有±1分钟的误差，对已完成考试无法获取截止时间
	// 精确数据请先打开考试
	Time time.Time
	// 待做、已完成
	Status string
	req    *request.Request
}

func (b *Brief) Open() (*Exam, error) {
	//todo: open a exam
	return nil, nil
}
