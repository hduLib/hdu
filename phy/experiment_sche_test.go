package phy

import (
	"fmt"
	"testing"
)

func TestExperSche(t *testing.T) {
	err := Login(studentId, password)
	if err != nil {
		t.Fatal(err)
	}
	expers, err := GetExperimentSche()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%#v", expers)
}
