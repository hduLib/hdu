package skl

func (user *User) Push(payload *PushReq) error {
	_, err := user.Post(payload)
	return err
}

func (user *User) My(payload *MyResp) error {
	return user.Get(payload)
}

func (user *User) UserInfo(payload *UserInfoResp) error {
	return user.Get(payload)
}

func (user *User) Leave(payload *LeaveReq) error {
	_, err := user.Post(payload)
	return err
}
