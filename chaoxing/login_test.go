package chaoxing

import "testing"

const phone = "111111111"
const passwd = "111111111"

func TestLogin(t *testing.T) {
	user, err := loginByPhoneAndPwd(phone, passwd)
	if err != nil {
		t.Error(err)
		return
	}
	for _, v := range user.cookie {
		t.Log(v.String())
	}
}
