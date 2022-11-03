package hduhelp

import (
	"fmt"
	"testing"
)

func TestTime(t *testing.T) {
	time, err := Time()
	if err != nil {
		t.Log(err)
		return
	}
	fmt.Println(time)
}
