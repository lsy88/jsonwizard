package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(str []byte, b ...byte) string {
	m := md5.New()
	m.Write(str)                        //哈希的到16位字节数组
	return hex.EncodeToString(m.Sum(b)) //将字节数组转换成16进制字符串
}
