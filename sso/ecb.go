package sso

import (
	"bytes"
	"crypto/des"
	"encoding/base64"
	"fmt"
)

func EncryptPasswd(key []byte, password string) (string, error) {
	var keyBytes [8]byte
	if _, err := base64.StdEncoding.Decode(keyBytes[:], key); err != nil {
		return "", fmt.Errorf("decode key: %v", err)
	}
	cipher, err := des.NewCipher(keyBytes[:])
	if err != nil {
		return "", fmt.Errorf("des cipher: %v", err)
	}
	// padding
	text := pkcs7Padding([]byte(password), cipher.BlockSize())
	// ecb mode
	for i := 0; i < len(text); i += cipher.BlockSize() {
		cipher.Encrypt(text[i:], text[i:])
	}
	return base64.StdEncoding.EncodeToString(text), nil
}

func pkcs7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	return append(data, bytes.Repeat([]byte{byte(padding)}, padding)...)
}
