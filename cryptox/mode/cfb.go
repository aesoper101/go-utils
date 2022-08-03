package mode

import (
	"crypto/cipher"
)

type cfb struct {
	baseCryptoMode
}

func newCFB() CryptoMode {
	return &cfb{}
}

func (c *cfb) Encrypt(plaintext []byte) ([]byte, error) {
	block := c.GetBlock()
	blockSize := block.BlockSize()
	iv := c.GetIV()

	padding := c.GetPadding()
	paddingSrc, err := padding.Padding(plaintext, blockSize)
	if err != nil {
		return nil, err
	}

	cipherText := make([]byte, len(paddingSrc))
	crypto := cipher.NewCFBEncrypter(block, iv)
	crypto.XORKeyStream(cipherText, paddingSrc)

	return cipherText, nil
}

func (c *cfb) Decrypt(cipherText []byte) ([]byte, error) {
	plainText := make([]byte, len(cipherText))

	block := c.GetBlock()
	iv := c.GetIV()

	crypto := cipher.NewCFBDecrypter(block, iv)
	crypto.XORKeyStream(plainText, cipherText)

	return c.GetPadding().UnPadding(plainText, block.BlockSize())
}
