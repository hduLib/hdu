package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
	"time"
)

const key = "u2oh6Vu^HWe4_AES"

type toFind interface {
	Find(selector string) *goquery.Selection
}

func EncryptByAES(msg string) string {
	block, _ := aes.NewCipher([]byte(key))
	cbc := cipher.NewCBCEncrypter(block, []byte(key))
	padding := aes.BlockSize - len(msg)%aes.BlockSize
	src := append([]byte(msg), bytes.Repeat([]byte{byte(padding)}, padding)...)
	dst := make([]byte, len(src))
	cbc.CryptBlocks(dst, src)
	return base64.StdEncoding.EncodeToString(dst)
}

func GetValueAttrBySelector(doc toFind, sel string) string {
	val, exist := doc.Find(sel).Attr("value")
	if !exist {
		log.Printf("%s not existed\n", val)
	}
	return val
}

func ParseLeftTime2Deadline(t string) time.Time {
	n := time.Now()
	if len(t) < 2 {
		return time.Time{}
	}
	if strings.Contains(t, "小时") {
		var hour, minute int
		fmt.Sscanf(t, "剩余%d小时%d分钟", &hour, &minute)
		n = n.Add(time.Duration(hour)*time.Hour + time.Duration(minute)*time.Minute)
	} else {
		var minute int
		fmt.Sscanf(t, "剩余%d分钟", &minute)
		n = n.Add(time.Duration(minute) * time.Minute)
	}
	if n.Second() != 0 {
		n = n.Add(time.Minute - time.Duration(n.Second())*time.Second)
	}
	return n.Round(time.Second)
}
