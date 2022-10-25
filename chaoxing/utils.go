package chaoxing

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

const key = "u2oh6Vu^HWe4_AES"

func encryptByAES(msg string) string {
	block, _ := aes.NewCipher([]byte(key))
	cbc := cipher.NewCBCEncrypter(block, []byte(key))
	padding := aes.BlockSize - len(msg)%aes.BlockSize
	src := append([]byte(msg), bytes.Repeat([]byte{byte(padding)}, padding)...)
	dst := make([]byte, len(src))
	cbc.CryptBlocks(dst, src)
	return base64.StdEncoding.EncodeToString(dst)
}
