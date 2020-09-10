package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
)

func main() {
	key := "XjFN@9^vZPWkSrzP"
	src := "momo"
	fmt.Println(src, "------------", key)
	aesEncrypt, _ := AESEncrypt(src, key)
	fmt.Println(aesEncrypt)
	aesDecrypt, _ := AesDecrypt(aesEncrypt, key)
	fmt.Println(aesDecrypt)

}

// AES-CBC 加密
func AESEncrypt(plantText, key string) (string, error) {
	en, err := encrypt([]byte(plantText), []byte(key))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(en), nil
}

// AES-CBC 解密
func AesDecrypt(plantText, key string) (string, error) {
	deData, _ := hex.DecodeString(plantText)
	de, err := decrypt(deData, []byte(key))
	if err != nil {
		return "", err
	}
	return string(de), nil
}

func encrypt(plantText, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key) //选择加密算法
	if err != nil {
		return nil, err
	}
	plantText = pKCS7Padding(plantText, block.BlockSize())

	blockModel := cipher.NewCBCEncrypter(block, key)

	cipherText := make([]byte, len(plantText))

	blockModel.CryptBlocks(cipherText, plantText)

	return cipherText, nil
}

func pKCS7Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padText...)
}

func decrypt(cipherText, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key) //选择加密算法
	if err != nil {
		return nil, err
	}
	blockModel := cipher.NewCBCDecrypter(block, key)
	plantText := make([]byte, len(cipherText))
	blockModel.CryptBlocks(plantText, cipherText)
	plantText = pKCS7UnPadding(plantText)
	return plantText, nil
}

func pKCS7UnPadding(plantText []byte) []byte {
	length := len(plantText)
	unPadding := int(plantText[length-1])
	return plantText[:(length - unPadding)]
}
