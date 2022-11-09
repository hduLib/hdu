package zjooc

import (
	"os"
	"testing"
)

func TestLogin(t *testing.T) {
	user, err := Login(os.Getenv("zjooc_account"), os.Getenv("zjooc_passwd"))
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(user.openid)
}
