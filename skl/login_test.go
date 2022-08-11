package skl

import "testing"

func TestLogin(t *testing.T) {
	skl, err := Login("21111111", "11111111")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(skl.xAuthToken)
	return
}
