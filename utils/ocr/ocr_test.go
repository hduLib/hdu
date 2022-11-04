package ocr_test

import (
	"io"
	"os"
	"testing"

	"github.com/hduLib/hdu/utils/ocr"
)

func TestOCR(t *testing.T) {
	ocr.SetToken("") // you should set your yunma token first
	res, err := ocr.RecognizeWithType(ocr.Common, readInImage())
	if err != nil {
		t.Fatal(err)
	}
	t.Log("ocr result:", res)
}

func readInImage() io.Reader {
	f, _ := os.OpenFile("testdata/4.jfif", os.O_RDONLY, 0644)
	return f
}
