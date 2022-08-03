package mode

import (
	"crypto/cipher"
)

type ctr struct {
	baseCryptoMode
}

func newCTR() CryptoMode {
	return &ctr{}
}

func (c *ctr) Encrypt(plaintext []byte) ([]byte, error) {
	block := c.GetBlock()
	blockSize := block.BlockSize()

	padding := c.GetPadding()
	paddingSrc, err := padding.Padding(plaintext, blockSize)
	if err != nil {
		return nil, err
	}

	iv := c.GetIV()

	cipherText := make([]byte, len(paddingSrc))
	crypto := cipher.NewCTR(block, iv)
	crypto.XORKeyStream(cipherText, paddingSrc)
	return cipherText, nil
}

func (c *ctr) Decrypt(cipherText []byte) ([]byte, error) {
	block := c.GetBlock()

	iv := c.GetIV()

	plainText := make([]byte, len(cipherText))
	crypto := cipher.NewCTR(block, iv)
	crypto.XORKeyStream(plainText, cipherText)

	return c.GetPadding().UnPadding(plainText, block.BlockSize())
}
