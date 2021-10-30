package cas

import (
	"testing"
)

func TestAuth(t *testing.T) {
	token, err := Login("https://api.hduhelp.com/login/direct/cas?clientID=healthcheckin&redirect=https://healthcheckin.hduhelp.com/#/auth", "21111111", "666666")
	if err != nil {
		t.Error(err)
	}
	t.Log(token)
}
