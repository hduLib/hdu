package sso

import (
	"bytes"
	"crypto/aes"
	"encoding/base64"
)

// AesEncrypt 使用 AES/ECB/PKCS7Padding 模式加密文本
func AesEncrypt(key, plainText string) (string, error) {
	keyBytes, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}

	paddedText := PKCS7Padding([]byte(plainText), block.BlockSize())
	cipherText := make([]byte, len(paddedText))

	// ECB 模式是分块加密
	for bs, be := 0, block.BlockSize(); bs < len(paddedText); bs, be = bs+block.BlockSize(), be+block.BlockSize() {
		block.Encrypt(cipherText[bs:be], paddedText[bs:be])
	}

	encodedCipherText := base64.StdEncoding.EncodeToString(cipherText)
	return encodedCipherText, nil
}

// PKCS7Padding 将明文填充到块大小的倍数
func PKCS7Padding(plainText []byte, blockSize int) []byte {
	padding := blockSize - len(plainText)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(plainText, padtext...)
}
