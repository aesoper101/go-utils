package mode

import (
	"github.com/aesoper101/go-utils/cryptox/cipher"
)

type ecb struct {
	baseCryptoMode
}

func newECB() CryptoMode {
	return &ecb{}
}

func (c *ecb) Encrypt(plaintext []byte) ([]byte, error) {
	block := c.GetBlock()
	blockSize := block.BlockSize()

	padding := c.GetPadding()
	paddingSrc, err := padding.Padding(plaintext, blockSize)
	if err != nil {
		return nil, err
	}

	encryptData := make([]byte, len(paddingSrc))
	cipher.NewECBEncrypter(block).CryptBlocks(encryptData, paddingSrc)
	return encryptData, nil
}

func (c *ecb) Decrypt(src []byte) ([]byte, error) {
	dst := make([]byte, len(src))

	block := c.GetBlock()
	mode := cipher.NewECBDecrypter(block)
	mode.CryptBlocks(dst, src)

	return c.GetPadding().UnPadding(dst, block.BlockSize())
}
