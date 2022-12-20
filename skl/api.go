package skl

import (
	"encoding/json"
	"github.com/hduLib/hdu/client"
	"time"
)

func (user *User) Push(payload *PushReq) error {
	_, err := user.post(pushURL, payload)
	if err == nil {
		return nil
	}
	if err, ok := err.(*client.ErrNotOk); ok {
		if err.StatusCode == 400 {
			var msg errorMsg
			if err1 := json.Unmarshal([]byte(err.Body), &msg); err1 != nil {
				return err
			}
			if msg.Code == 0 && msg.Msg == "今日已经打卡" {
				return ErrAlreadyPushed
			}
		}
	}
	return err
}

func (user *User) PushLogs() (*PushLogResp, error) {
	resp := new(PushLogResp)
	return resp, user.get(pushLogURL, resp)
}

func (user *User) My() (*MyResp, error) {
	resp := new(MyResp)
	return resp, user.get(myURL, resp)
}

func (user *User) UserInfo() (*UserInfoResp, error) {
	resp := new(UserInfoResp)
	return resp, user.get(userInfoURL, resp)
}

func (user *User) Leave(payload *LeaveReq) error {
	_, err := user.post(leaveURL, payload)
	return err
}

// Course needs a startTime, which determined the semester
// that the returned course list belongs to. So you may simply
// use time.Now() to get the current course list.
func (user *User) Course(startTime time.Time) (*CourseResp, error) {
	resp := new(CourseResp)
	return resp, user.get(courseURL+"?startTime="+startTime.Format("2006-01-02"), resp)
}
