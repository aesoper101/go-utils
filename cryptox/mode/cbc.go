package mode

import (
	"crypto/cipher"
)

type cbc struct {
	baseCryptoMode
}

func newCBC() CryptoMode {
	return &cbc{}
}

func (c *cbc) Encrypt(plaintext []byte) ([]byte, error) {
	block := c.GetBlock()
	blockSize := block.BlockSize()

	padding := c.GetPadding()
	paddingSrc, err := padding.Padding(plaintext, blockSize)
	if err != nil {
		return nil, err
	}

	encryptData := make([]byte, len(paddingSrc))
	cipher.NewCBCEncrypter(block, c.GetIV()).CryptBlocks(encryptData, paddingSrc)

	return encryptData, nil
}

func (c *cbc) Decrypt(src []byte) ([]byte, error) {
	dst := make([]byte, len(src))

	block := c.GetBlock()
	cipher.NewCBCDecrypter(block, c.GetIV()).CryptBlocks(dst, src)

	return c.GetPadding().UnPadding(dst, block.BlockSize())
}
