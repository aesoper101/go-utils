package mode

import (
	"crypto/cipher"
)

type gcm struct {
	baseCryptoMode
}

func newGcm() CryptoMode {
	return &gcm{}
}

func (o *gcm) Encrypt(plaintext []byte) ([]byte, error) {
	block := o.GetBlock()
	blockSize := block.BlockSize()
	iv := o.GetIV()

	padding := o.GetPadding()
	paddingSrc, err := padding.Padding(plaintext, blockSize)
	if err != nil {
		return nil, err
	}

	crypto, err := cipher.NewGCMWithNonceSize(block, len(iv))
	if err != nil {
		return nil, err
	}

	cipherText := crypto.Seal(nil, iv, paddingSrc, nil)

	return cipherText, nil
}

func (o *gcm) Decrypt(cipherText []byte) ([]byte, error) {
	block := o.GetBlock()
	iv := o.GetIV()

	crypto, err := cipher.NewGCMWithNonceSize(block, len(iv))
	if err != nil {
		return nil, err
	}

	plainText, err := crypto.Open(nil, iv, cipherText, nil)
	if err != nil {
		return nil, err
	}

	return o.GetPadding().UnPadding(plainText, block.BlockSize())
}
