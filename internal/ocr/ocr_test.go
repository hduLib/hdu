package ocr_test

import (
	ocr2 "github.com/hduLib/hdu/internal/ocr"
	"io"
	"os"
	"testing"
)

func TestOCR(t *testing.T) {
	ocr2.SetToken("") // you should set your yunma token first
	res, err := ocr2.RecognizeWithType(ocr2.Common, readInImage())
	if err != nil {
		t.Fatal(err)
	}
	t.Log("ocr result:", res)
}

func readInImage() io.Reader {
	f, _ := os.OpenFile("testdata/4.jfif", os.O_RDONLY, 0644)
	return f
}
