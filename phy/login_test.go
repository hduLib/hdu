package phy_test

import (
	"log"
	"testing"

	"github.com/hduLib/hdu/phy"
	"github.com/hduLib/hdu/utils/ocr"
)

func TestLogin(t *testing.T) {
	// set your own token and you username & password
	ocr.SetToken(``)
	if err := phy.Login("", ""); err != nil {
		log.Println(err)
		t.FailNow()
	}
}
