package phy

import (
	"fmt"
	"testing"
)

func TestExperSche(t *testing.T) {
	u, _ := Login(studentId, password)
	experiments, err := u.GetExperimentSche()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v", experiments)
}
