package padding

import "fmt"

var _ Padding = (*iso10126Padding)(nil)

type iso10126Padding struct{}

func NewISO10126Padding() Padding {
	return &iso10126Padding{}
}

func (i *iso10126Padding) Padding(plaintext []byte, blockSize int) ([]byte, error) {
	if blockSize < 1 || blockSize > 256 {
		return nil, fmt.Errorf("padding.ISO10126Padding blockSize is out of bounds: %d", blockSize)
	}
	padding := padSize(len(plaintext), blockSize)
	padText := append(randBytes(padding-1), byte(padding))
	return append(plaintext, padText...), nil
}

func (i *iso10126Padding) UnPadding(ciphertext []byte, blockSize int) ([]byte, error) {
	length := len(ciphertext)
	if length%blockSize != 0 {
		return nil, fmt.Errorf("padding.ISO10126UnPadding ciphertext's length isn't a multiple of blockSize")
	}
	unPadding := int(ciphertext[length-1])
	if unPadding > blockSize || unPadding < 1 {
		return nil, fmt.Errorf("padding.ISO10126UnPadding invalid padding found: %v", unPadding)
	}
	return ciphertext[:length-unPadding], nil
}
