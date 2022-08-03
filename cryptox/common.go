package cryptox

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"github.com/aesoper101/go-utils/cryptox/mode"
	"github.com/aesoper101/go-utils/encodex"
)

type Crypto interface {
	Encrypt(plaintext []byte) (*encodex.Encode, error)
	Decrypt(ciphertext []byte) (plaintext []byte, err error)
}

func PadZeroKey(key []byte, size int) []byte {
	if len(key) == size {
		return key
	}
	if len(key) > size {
		return key[:size]
	}
	return append(key, bytes.Repeat([]byte{0}, size-len(key))...)
}

func createCryptoMode(opts options) (mode.CryptoMode, error) {
	block, err := opts.GetBlock()
	if err != nil {
		return nil, err
	}

	crypto, err := mode.GetCryptoMode(opts.mode.String())
	if err != nil {
		return nil, err
	}

	crypto.SetBlock(block)
	crypto.SetIV(opts.iv)
	crypto.SetPadding(opts.padding)
	crypto.SetKey(opts.key)

	return crypto, nil
}

func newBlock(m Method, key []byte) (cipher.Block, error) {
	var block cipher.Block
	var err error
	switch m {
	case AES:
		block, err = aes.NewCipher(key)
		break
	case DES:
		block, err = des.NewCipher(key)
		break
	case DES3:
		block, err = des.NewTripleDESCipher(key)
		break
	default:
		block, err = aes.NewCipher(key)
	}

	return block, err
}
