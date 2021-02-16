package crypto

import (
	"encoding/base32"
	"encoding/base64"
)

// base32解密
func DecodeBase32(data string) ([]byte, error) {
	return base32.StdEncoding.DecodeString(data)
}

// base64解密
func DecodeBase64(data string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(data)
}
