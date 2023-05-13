package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

//data 是json串, key是32位加密密钥
//EncryptJSON 函数使用 AES-256 算法,CBC 模式加密,并在加密前生成一个随机的 16 字节的初始化向量 (IV)

func EncryptJSON(data []byte, key []byte) ([]byte, error) {
	plainText := data
	
	// Generate a random 16-byte initialization vector (IV)
	iv := make([]byte, 16)
	_, err := io.ReadFull(rand.Reader, iv)
	if err != nil {
		return nil, err
	}
	//创建加密器
	cipherBlock, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//获取加密器所采用的分组大小，把原始文本填充（padding）成分组大小的倍数
	blockSize := cipherBlock.BlockSize()
	paddedPlainText := make([]byte, len(plainText)+(blockSize-len(plainText)%blockSize))
	copy(paddedPlainText, plainText)
	
	// Encrypt the plaintext using CBC mode
	encrypted := make([]byte, len(paddedPlainText)+len(iv))
	copy(encrypted[:16], iv)
	//创建CBC模式加密器
	cbc := cipher.NewCBCEncrypter(cipherBlock, iv)
	cbc.CryptBlocks(encrypted[16:], paddedPlainText)
	
	return encrypted, nil
}
