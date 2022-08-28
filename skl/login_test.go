package skl

import "testing"

func TestLogin(t *testing.T) {
	skl, err := Login(id, passwd)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(skl.xAuthToken)
	return
}
