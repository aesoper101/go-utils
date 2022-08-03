package padding

import (
	"bytes"
	"fmt"
)

var _ Padding = (*zeroPadding)(nil)

type zeroPadding struct{}

func NewZeroPadding() Padding {
	return &zeroPadding{}
}

func (z *zeroPadding) Padding(plaintext []byte, blockSize int) ([]byte, error) {
	if blockSize < 1 || blockSize > 255 {
		return nil, fmt.Errorf("padding.ZeroPadding blockSize is out of bounds: %d", blockSize)
	}
	padding := padSize(len(plaintext), blockSize)
	padText := bytes.Repeat([]byte{0}, padding)
	return append(plaintext, padText...), nil
}

func (z *zeroPadding) UnPadding(ciphertext []byte, _ int) ([]byte, error) {
	return bytes.TrimRightFunc(ciphertext, func(r rune) bool { return r == rune(0) }), nil
}
