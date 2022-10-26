package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"github.com/PuerkitoBio/goquery"
	"log"
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
