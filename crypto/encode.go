package crypto

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base32"
	"encoding/base64"
	"encoding/hex"
	"golang.org/x/crypto/sha3"
)

// md5哈希加密
func EncodeMd5(data []byte, sum []byte) (string, error) {
	md := md5.New()
	if _, err := md.Write(data); err != nil {
		return "", err
	}
	content := md.Sum(sum)
	return hex.EncodeToString(content), nil
}

// sha1哈希加密
func EncodeSha1(data []byte, sum []byte) (string, error) {
	sha := sha1.New()
	if _, err := sha.Write(data); err != nil {
		return "", err
	}
	content := sha.Sum(sum)
	return hex.EncodeToString(content), nil
}

// sha256哈希加密
func EncodeSha256(data []byte, sum []byte) (string, error) {
	sha := sha256.New()
	if _, err := sha.Write(data); err != nil {
		return "", err
	}
	content := sha.Sum(sum)
	return hex.EncodeToString(content), nil
}

// sha512哈希加密
func EncodeSha512(data []byte, sum []byte) (string, error) {
	sha := sha512.New()
	if _, err := sha.Write(data); err != nil {
		return "", err
	}
	content := sha.Sum(sum)
	return hex.EncodeToString(content), nil
}

// sha3-224哈希加密
func EncodeSha3_224(data []byte, sum []byte) (string, error) {
	sha := sha3.New224()
	if _, err := sha.Write(data); err != nil {
		return "", err
	}
	var buf []byte
	content := sha.Sum(sum)
	for _, b := range content {
		buf = append(buf, b)
	}
	return hex.EncodeToString(buf), nil
}

// sha3-256哈希加密
func EncodeSha3_256(data []byte, sum []byte) (string, error) {
	sha := sha3.New256()
	if _, err := sha.Write(data); err != nil {
		return "", err
	}
	var buf []byte
	content := sha.Sum(sum)
	for _, b := range content {
		buf = append(buf, b)
	}
	return hex.EncodeToString(buf), nil
}

// sha3-384哈希加密
func EncodeSha3_384(data []byte, sum []byte) (string, error) {
	sha := sha3.New384()
	if _, err := sha.Write(data); err != nil {
		return "", err
	}
	var buf []byte
	content := sha.Sum(sum)
	for _, b := range content {
		buf = append(buf, b)
	}
	return hex.EncodeToString(buf), nil
}

// sha3-512哈希加密
func EncodeSha3_512(data []byte, sum []byte) (string, error) {
	sha := sha3.New512()
	if _, err := sha.Write(data); err != nil {
		return "", err
	}
	var buf []byte
	content := sha.Sum(sum)
	for _, b := range content {
		buf = append(buf, b)
	}
	return hex.EncodeToString(buf), nil
}

// base32加密
func EncodeBase32(data []byte) string {
	return base32.StdEncoding.EncodeToString(data)
}

// base64加密
func EncodeBase64(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}
