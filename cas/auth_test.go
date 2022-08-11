package cas

import (
	"testing"
)

func TestGenLoginReq(t *testing.T) {
	req, err := GenLoginReq("https://api.hduhelp.com/login/direct/cas?clientID=healthcheckin&redirect=https://healthcheckin.hduhelp.com/#/auth", "21111111", "666666")
	if err != nil {
		t.Error(err)
	}
	t.Log(req)
}
