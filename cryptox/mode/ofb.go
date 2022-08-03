package mode

import (
	"crypto/cipher"
)

type ofb struct {
	baseCryptoMode
}

func newOFB() CryptoMode {
	return &ofb{}
}

func (o *ofb) Encrypt(plaintext []byte) ([]byte, error) {
	block := o.GetBlock()
	blockSize := block.BlockSize()
	iv := o.GetIV()

	padding := o.GetPadding()
	paddingSrc, err := padding.Padding(plaintext, blockSize)
	if err != nil {
		return nil, err
	}

	cipherText := make([]byte, len(paddingSrc))
	crypto := cipher.NewOFB(block, iv)
	crypto.XORKeyStream(cipherText, paddingSrc)

	return cipherText, nil
}

func (o *ofb) Decrypt(cipherText []byte) ([]byte, error) {
	plainText := make([]byte, len(cipherText))

	block := o.GetBlock()
	iv := o.GetIV()

	crypto := cipher.NewOFB(block, iv)
	crypto.XORKeyStream(plainText, cipherText)

	return o.GetPadding().UnPadding(plainText, block.BlockSize())
}
