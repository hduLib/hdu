package phy

import (
	"github.com/hduLib/hdu/internal/ocr"
	"log"
	"os"
	"testing"

	_ "github.com/joho/godotenv/autoload"
)

func TestSignIn(t *testing.T) {
	// set your own token and you username & password
	ocr.SetToken(os.Getenv("TOKEN"))
	err := SignIn(os.Getenv("USERNAME"), os.Getenv("PASSWORD"))
	if err != nil {
		log.Println(err)
		t.FailNow()
	}
	log.Println(jSessionId)
}
