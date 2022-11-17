package phy

import (
	"fmt"
	"os"
	"testing"

	"github.com/hduLib/hdu/internal/ocr"
	_ "github.com/joho/godotenv/autoload"
)

var studentId, password string

func TestMain(m *testing.M) {
	ocr.SetToken(os.Getenv("TOKEN"))
	studentId = os.Getenv("STUDENTID")
	password = os.Getenv("PASSWORD")
	m.Run()
}

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
