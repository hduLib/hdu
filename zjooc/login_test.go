package zjooc

import "testing"

const (
	account  = "11111111"
	password = "11111111"
)

func TestLogin(t *testing.T) {
	user, err := Login(account, password)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(user.openid)
}
