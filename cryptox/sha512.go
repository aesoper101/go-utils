package cryptox

import (
	"crypto/sha512"
	"github.com/aesoper101/go-utils/encodex"
)

func SHA512(data []byte) *encodex.Encode {
	b := sha512.Sum512(data)
	return encodex.FromBytes(b[:])
}

func SHA384(data []byte) *encodex.Encode {
	b := sha512.Sum384(data)
	return encodex.FromBytes(b[:])
}
