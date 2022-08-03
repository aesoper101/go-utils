package cryptox

import (
	"crypto/hmac"
	"crypto/sha256"
)

func HmacSha256(data []byte, key string) string {
	hashed := hmac.New(sha256.New, []byte(key))
	hashed.Write(data)

	return string(hashed.Sum(nil))
}

func HmacSha256String(data string, key string) string {
	return HmacSha256([]byte(data), key)
}
