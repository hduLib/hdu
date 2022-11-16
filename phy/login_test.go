package phy

import (
	"os"
	"testing"

	"github.com/hduLib/hdu/internal/ocr"

	_ "github.com/joho/godotenv/autoload"
)

var studentId, password string

func TestMain(m *testing.M) {
	ocr.SetToken(os.Getenv("TOKEN"))
	studentId = os.Getenv("USERNAME")
	password = os.Getenv("PASSWORD")
	m.Run()
}

func TestLogin(t *testing.T) {
	err := Login(studentId, password)
	if err != nil {
		t.Fatal(err)
	}
	if JSessionId == "" {
		t.Fatal("no JSESSIONID")
	}
}
