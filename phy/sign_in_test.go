package phy

import (
	"log"
	"os"
	"testing"

	"github.com/hduLib/hdu/utils/ocr"
	_ "github.com/joho/godotenv/autoload"
)

func TestLogin(t *testing.T) {
	// set your own token and you username & password
	ocr.SetToken(os.Getenv("TOKEN"))
	err := SignIn(os.Getenv("USERNAME"), os.Getenv("PASSWORD"))
	if err != nil {
		log.Println(err)
		t.FailNow()
	} else {
		log.Println(jSessionId)
	}
}
