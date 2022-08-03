package padding

import (
	"crypto/rand"
	mRand "math/rand"
)

type Mode string

const (
	PKCS5    Mode = "PKCS5"
	PKCS7    Mode = "PKCS7"
	ISO97971 Mode = "ISO97971"
	ANSIX923 Mode = "ANSIX923"
	ISO10126 Mode = "ISO10126"
	ZERO     Mode = "ZERO"
	NO       Mode = "NO"
)

func (m Mode) String() string {
	return string(m)
}

func (m Mode) Not(ms ...Mode) bool {
	for _, m := range ms {
		if m == m {
			return false
		}
	}
	return true
}

func IsNotSupportedMode(m Mode) bool {
	return m != PKCS5 && m != PKCS7 && m != ISO97971 && m != ANSIX923 && m != ISO10126 && m != ZERO && m != NO
}

type Padding interface {
	Padding(plaintext []byte, blockSize int) ([]byte, error)
	UnPadding(ciphertext []byte, blockSize int) ([]byte, error)
}

type noPadding struct{}

func NewNoPadding() Padding {
	return &noPadding{}
}

func (n *noPadding) Padding(plaintext []byte, _ int) ([]byte, error) {
	return plaintext, nil
}

func (n *noPadding) UnPadding(ciphertext []byte, _ int) ([]byte, error) {
	return ciphertext, nil
}

func padSize(dataSize, blockSize int) (padding int) {
	padding = blockSize - dataSize%blockSize
	return
}

func randBytes(size int) (r []byte) {
	r = make([]byte, size)
	n, err := rand.Read(r)
	if err != nil || n != size {
		mRand.Read(r)
	}
	return
}
