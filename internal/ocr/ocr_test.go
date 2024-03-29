package ocr_test

import (
	"io"
	"os"
	"testing"

	ocr2 "github.com/hduLib/hdu/internal/ocr"
	_ "github.com/joho/godotenv/autoload"
)

func TestOCR(t *testing.T) {
	ocr2.SetToken(os.Getenv("TOKEN")) // you should set your yunma token first
	res, err := ocr2.Parse(readInImage())
	if err != nil {
		t.Fatal(err)
	}
	t.Log("ocr result:", res)
	if res != "vyza" {
		t.Fatalf(`ocr error, expect "vyza", found %s`, res)
	}
}

func readInImage() io.Reader {
	f, _ := os.OpenFile("testdata/4.jfif", os.O_RDONLY, 0644)
	return f
}
