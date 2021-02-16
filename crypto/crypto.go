package crypto

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
)

// sha1哈希加密
func EncodeSha1(data []byte) string {
	sha := sha1.New()
	sha.Write(data)
	sum := sha.Sum([]byte(""))
	return hex.EncodeToString(sum)
}

// md5哈希加密
func EncodeMd5(data []byte) string {
	md := md5.New()
	md.Write(data)
	sum := md.Sum([]byte(""))
	return hex.EncodeToString(sum)
}
