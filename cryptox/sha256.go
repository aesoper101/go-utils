package cryptox

import (
	"crypto/sha256"
	"github.com/aesoper101/go-utils/encodex"
)

func SHA256(data []byte) *encodex.Encode {
	b := sha256.Sum256(data)
	return encodex.FromBytes(b[:])
}

func SHA224(data []byte) *encodex.Encode {
	b := sha256.Sum224(data)
	return encodex.FromBytes(b[:])
}
