package padding

import (
	"bytes"
	"fmt"
)

var _ Padding = (*pkcs7Padding)(nil)

type pkcs7Padding struct{}

func NewPKCS7Padding() Padding {
	return &pkcs7Padding{}
}

func (p *pkcs7Padding) Padding(plaintext []byte, blockSize int) ([]byte, error) {
	if blockSize < 1 || blockSize > 255 {
		return nil, fmt.Errorf("padding.PKCS7Padding blockSize is out of bounds: %d", blockSize)
	}
	padding := padSize(len(plaintext), blockSize)
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(plaintext, padText...), nil
}

func (p *pkcs7Padding) UnPadding(ciphertext []byte, blockSize int) ([]byte, error) {
	length := len(ciphertext)
	if length%blockSize != 0 {
		return nil, fmt.Errorf("padding.PKCS7UnPadding ciphertext's length isn't a multiple of blockSize")
	}
	unPadding := int(ciphertext[length-1])
	if unPadding > blockSize || unPadding <= 0 {
		return nil, fmt.Errorf("padding.PKCS7UnPadding invalid padding found: %v", unPadding)
	}
	var pad = ciphertext[length-unPadding : length-1]
	for _, v := range pad {
		if int(v) != unPadding {
			return nil, fmt.Errorf("padding.PKCS7UnPadding invalid padding found")
		}
	}
	return ciphertext[:length-unPadding], nil
}
