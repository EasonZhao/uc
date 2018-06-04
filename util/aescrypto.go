//package main

package util

import (
	"crypto/aes"
	"crypto/cipher"
	b64 "encoding/base64"
	// "fmt"
	"strings"
)

// func main() {
// 	strKey := "12345678901234561_123456789"
// 	arrEncrypt, err := AesEncrypt("andyniss@126.com;1527734087", []byte(strKey))
// 	if err != nil {
// 		return
// 	}
// 	fmt.Println(string(arrEncrypt))
// 	sEnc := Enb64(string(arrEncrypt), true)
// 	fmt.Println(sEnc)
// 	sDec, _ := Unb64(sEnc, true)
// 	strMsg, err := AesDecrypt(sDec, []byte(strKey))
// 	if err != nil {
// 		fmt.Println(sDec)
// 		return
// 	}
// 	fmt.Println(strMsg)
// }

func getKey(c []byte) []byte {
	keyLen := len(c)
	if keyLen < 16 {
		panic("res key can't be less than 16 bytes")
	}
	arrKey := []byte(c)
	if keyLen >= 32 {
		//取前32个字节
		return arrKey[:32]
	}
	if keyLen >= 24 {
		//取前24个字节
		return arrKey[:24]
	}
	//取前16个字节
	return arrKey[:16]
}

//加密字符串
func AesEncrypt(strMesg string, c []byte) ([]byte, error) {
	key := getKey(c)
	var iv = []byte(key)[:aes.BlockSize]
	encrypted := make([]byte, len(strMesg))
	aesBlockEncrypter, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	aesEncrypter := cipher.NewCFBEncrypter(aesBlockEncrypter, iv)
	aesEncrypter.XORKeyStream(encrypted, []byte(strMesg))
	return encrypted, nil
}

//解密字符串
func AesDecrypt(src, c []byte) (strDesc string, err error) {
	defer func() {
		//错误处理
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	key := getKey(c)
	var iv = []byte(key)[:aes.BlockSize]
	decrypted := make([]byte, len(src))
	var aesBlockDecrypter cipher.Block
	aesBlockDecrypter, err = aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}
	aesDecrypter := cipher.NewCFBDecrypter(aesBlockDecrypter, iv)
	aesDecrypter.XORKeyStream(decrypted, src)
	return string(decrypted), nil
}

func Enb64(origs string, replace bool) string {
	encs := b64.StdEncoding.EncodeToString([]byte(origs))
	if replace {
		encs = strings.Replace(encs, "/", "@", -1)
		encs = strings.Replace(encs, "+", "$", -1)
		encs = strings.Replace(encs, "=", "!", -1)
	}
	return encs
}

func Unb64(encs string, replace bool) ([]byte, error) {
	if replace {
		encs = strings.Replace(encs, "@", "/", -1)
		encs = strings.Replace(encs, "$", "+", -1)
		encs = strings.Replace(encs, "!", "=", -1)
	}
	return b64.StdEncoding.DecodeString(encs)
}
