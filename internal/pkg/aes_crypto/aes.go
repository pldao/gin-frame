package aes_crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"io"
)

// key不能泄露
var PwdKey = []byte("ABCDEFGHIJKLMNO1")

// 加密函数
func AesEncryptCFB(origData []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	encrypted := make([]byte, aes.BlockSize+len(origData))
	iv := encrypted[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(encrypted[aes.BlockSize:], origData)
	return encrypted, nil
}

// 解密函数
func AesDecryptCFB(encrypted []byte, key []byte) ([]byte, error) {
	block, _ := aes.NewCipher(key)
	if len(encrypted) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	iv := encrypted[:aes.BlockSize]
	encrypted = encrypted[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(encrypted, encrypted)
	return encrypted, nil
}

// 加密
func EnPwdCode(pwd string) (string, error) {
	encryptCode, err := AesEncryptCFB([]byte(pwd), PwdKey)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(encryptCode), nil
}

// 解密
func DePwdCode(pwd string) (string, error) {
	edpass, err := hex.DecodeString(pwd)
	if err != nil {
		return "", err
	}
	decryptCode, err := AesDecryptCFB(edpass, PwdKey)
	if err != nil {
		return "", err
	}
	return string(decryptCode), nil

}
