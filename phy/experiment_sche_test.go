package phy

import (
	"testing"
)

func TestExperSche(t *testing.T) {
	_ = Login(studentId, password)
	_, err := GetExperimentSche()
	if err != nil {
		t.Fatal(err)
	}
}
