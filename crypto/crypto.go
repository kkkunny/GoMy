package crypto

import (
	"crypto/md5"
	"crypto/sha1"
)

// sha1哈希加密
func EncodeSha1(data []byte)[]byte{
	sha := sha1.New()
	sha.Write(data)
	return sha.Sum([]byte(""))
}

// md5哈希加密
func EncodeMd5(data []byte)[]byte{
	md := md5.New()
	md.Write(data)
	return md.Sum([]byte(""))
}