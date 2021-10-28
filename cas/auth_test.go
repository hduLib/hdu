package cas

import (
	"testing"
)

func TestAuth(t *testing.T) {
	token, err := GetToken("21111111", "66666666")
	if err != nil {
		t.Error(err)
	}
	t.Log(token)
}
