package skl

import "testing"

func TestSkl_Push(t *testing.T) {
	skl, err := Login("2111111", "11111111")
	if err != nil {
		t.Error(err)
		return
	}
	if err := skl.Push(); err != nil {
		t.Error(err)
		return
	}
}
