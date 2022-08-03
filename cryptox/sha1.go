package cryptox

import (
	"crypto/sha1"
	"github.com/aesoper101/go-utils/encodex"
)

// SHA1 returns the SHA1 hash of the string
func SHA1(data []byte) *encodex.Encode {
	h := sha1.New()
	h.Write(data)
	return encodex.FromBytes(h.Sum(nil))
}
