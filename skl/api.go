package skl

import "time"

func (user *User) Push(payload *PushReq) error {
	_, err := user.post(pushURL, payload)
	return err
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
func (user *User) Course(startTime time.Time) ([]course, error) {
	resp := new(courseResp)
	return resp.List, user.get(courseURL+"?startTime="+startTime.Format("2006-01-02"), resp)
}
