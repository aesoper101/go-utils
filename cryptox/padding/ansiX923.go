package padding

import (
	"bytes"
	"fmt"
)

var _ Padding = (*ansiX923Padding)(nil)

type ansiX923Padding struct{}

func NewAnsiX923Padding() Padding {
	return &ansiX923Padding{}
}

func (a *ansiX923Padding) Padding(plaintext []byte, blockSize int) ([]byte, error) {
	if blockSize < 1 || blockSize > 255 {
		return nil, fmt.Errorf("padding.AnsiX923Padding blockSize is out of bounds: %d", blockSize)
	}
	padding := padSize(len(plaintext), blockSize)
	padText := append(bytes.Repeat([]byte{byte(0)}, padding-1), byte(padding))
	return append(plaintext, padText...), nil
}

func (a *ansiX923Padding) UnPadding(ciphertext []byte, blockSize int) ([]byte, error) {
	length := len(ciphertext)
	if length%blockSize != 0 {
		return nil, fmt.Errorf("padding.AnsiX923UnPadding ciphertext's length isn't a multiple of blockSize")
	}
	unPadding := int(ciphertext[length-1])
	if unPadding > blockSize || unPadding < 1 {
		return nil, fmt.Errorf("padding.AnsiX923UnPadding invalid padding found: %d", unPadding)
	}
	if length-unPadding < length-2 {
		pad := ciphertext[length-unPadding : length-2]
		for _, v := range pad {
			if int(v) != 0 {
				return nil, fmt.Errorf("padding.AnsiX923UnPadding invalid padding found")
			}
		}
	}
	return ciphertext[0 : length-unPadding], nil
}
