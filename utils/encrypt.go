package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

//生成随机的16字节的初始化向量 (IV)
func init() {
	iv = make([]byte, 16)
	_, err := io.ReadFull(rand.Reader, iv)
	if err != nil {
		panic(err)
	}
}

var iv []byte

// EncryptJSON 加密使用 AES-256 算法,CBC 模式加密
// 填充秘钥key的16位，24,32分别对应AES-128, AES-192, or AES-256.
// data: 加密目标字符串
// key: 加密Key
func EncryptJSON(data string, key string) (string, error) {
	encrypted, err := aesCBCEncrypt([]byte(data), []byte(key), iv)
	if err != nil {
		return "", err
	}
	
	return base64.StdEncoding.EncodeToString(encrypted), nil
}

// DecryptJSON 解密
// encrypted: 解密目标字符串
// key: 加密Key
func DecryptJSON(encrypted string, key string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return "", err
	}
	//解密的iv要和加密的iv一致
	decrypted, err := aesCBCDecrypt(data, []byte(key), iv)
	if err != nil {
		return "", err
	}
	
	return string(decrypted), nil
}

// aesCBCEncrypt AES/CBC/PKCS7Padding 加密
func aesCBCEncrypt(plaintext []byte, key []byte, iv []byte) ([]byte, error) {
	// AES,创建加密器,返回AES算法的Block接口对象
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	// 填充
	plaintext = paddingPKCS7(plaintext, aes.BlockSize)
	
	// 选择分组加密模式--CBC加密
	// iv的长度必须和block的长度相同
	cbc := cipher.NewCBCEncrypter(block, iv)
	
	cbc.CryptBlocks(plaintext, plaintext)
	
	return plaintext, nil
}

// aesCBCDecrypt AES/CBC/PKCS7Padding 解密
func aesCBCDecrypt(ciphertext []byte, key []byte, iv []byte) ([]byte, error) {
	// AES
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	
	if len(ciphertext)%aes.BlockSize != 0 {
		return nil, errors.New("ciphertext is not a multiple of the block size")
	}
	
	// CBC 解密
	cbc := cipher.NewCBCDecrypter(block, iv)
	cbc.CryptBlocks(ciphertext, ciphertext)
	
	// PKCS7 反填充
	result := unPaddingPKCS7(ciphertext)
	return result, nil
}

// PKCS7 填充
func paddingPKCS7(plaintext []byte, blockSize int) []byte {
	paddingSize := blockSize - len(plaintext)%blockSize
	//对原来的明文填充paddingSize个[]byte{byte(paddingSize)}
	paddingText := bytes.Repeat([]byte{byte(paddingSize)}, paddingSize)
	//fmt.Println(paddingText)
	return append(plaintext, paddingText...)
}

// PKCS7 反填充
func unPaddingPKCS7(s []byte) []byte {
	length := len(s)
	if length == 0 {
		return s
	}
	//取出密文的最后一个字节
	unPadding := int(s[length-1])
	//删除填充
	return s[:(length - unPadding)]
}
