package padding

import (
	"fmt"
)

var _ Padding = (*pkcs5Padding)(nil)

type pkcs5Padding struct{}

func NewPKCS5Padding() Padding {
	return &pkcs5Padding{}
}

func (p *pkcs5Padding) Padding(plaintext []byte, blockSize int) ([]byte, error) {
	if blockSize != 8 {
		return nil, fmt.Errorf("padding.PKCS5Padding blockSize is not equal: %d", 8)
	}
	padding := &pkcs7Padding{}
	return padding.Padding(plaintext, blockSize)
}

func (p *pkcs5Padding) UnPadding(ciphertext []byte, blockSize int) ([]byte, error) {
	if blockSize != 8 {
		return nil, fmt.Errorf("padding.PKCS5Padding blockSize is not equal: %d", 8)
	}
	padding := &pkcs7Padding{}
	return padding.UnPadding(ciphertext, blockSize)
}
