package work

import (
	"github.com/hduLib/hdu/chaoxing/request"
	"time"
)

type Brief struct {
	url   string
	Title string
	// 根据剩余时间推断，可能有±1分钟的误差，对已完成作业无法获取截止时间
	// 精确数据请先打开作业
	Time time.Time
	// 未交（手机作业列表API->未完成）、已完成、待批阅
	Status  string
	ClazzId string
	req     *request.Request
}

func (b *Brief) Detail() *Work {
	_, err := b.req.Get(b.url)
	if err != nil {
		return nil
	}
	// todo: course detail, 其实也不一定会做，因为这个东西做出来也很难应用，brief就够了
	return nil
}
