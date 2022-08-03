package cryptox

import (
	"crypto/md5"
	"github.com/aesoper101/go-utils/encodex"
)

// MD5 returns the MD5 hash of the string
func MD5(data []byte) *encodex.Encode {
	h := md5.New()
	h.Write(data)
	return encodex.FromBytes(h.Sum(nil))
}
