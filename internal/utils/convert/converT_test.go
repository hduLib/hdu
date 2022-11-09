package convert_test

import (
	"encoding/base64"
	"fmt"
	"testing"

	"github.com/hduLib/hdu/utils/convert"
)

func TestToString(t *testing.T) {
	b := []byte{65}
	fmt.Println("b:", b)

	s := convert.ToString(b)
	fmt.Println("after convert:", s)
}

func TestToBytes(t *testing.T) {
	s := "Hello, world"
	fmt.Println("s:", s)

	b := convert.ToBytes(s)
	fmt.Println("after convert", b)
}

func TestBase64Encoding(t *testing.T) {
	username := "limuluteenpzsite"
	bUname := convert.ToBytes(username)
	obRes := base64.StdEncoding.EncodeToString([]byte(username))
	cbRes := base64.StdEncoding.EncodeToString(bUname)
	if cbRes != obRes {
		t.FailNow()
	}
	fmt.Println("It works!")
	fmt.Println("obRes:", obRes)
	fmt.Println("cbRes:", cbRes)
}
