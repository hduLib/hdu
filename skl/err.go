package skl

import "errors"

var ErrAlreadyPushed = errors.New("今日已经打卡")

type errorMsg struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
