package skl

import (
	"testing"
)

func TestUser_Upload(t *testing.T) {
	skl, err := Login(id, passwd)
	if err != nil {
		t.Error(err)
		return
	}
	file, err := skl.Upload("C:\\Users\\a\\Pictures\\zh.png")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(file)
}
