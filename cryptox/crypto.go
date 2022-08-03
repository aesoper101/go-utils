package cryptox

import "github.com/aesoper101/go-utils/encodex"

type baseCrypto struct {
	opts options
}

func (a *baseCrypto) Encrypt(plaintext []byte) (*encodex.Encode, error) {
	crypto, err := createCryptoMode(a.opts)
	if err != nil {
		return encodex.FromBytes([]byte{}), err
	}

	resultBytes, err := crypto.Encrypt(plaintext)
	if err != nil {
		return encodex.FromBytes([]byte{}), err
	}
	return encodex.FromBytes(resultBytes), err
}

func (a *baseCrypto) Decrypt(ciphertext []byte) ([]byte, error) {
	crypto, err := createCryptoMode(a.opts)
	if err != nil {
		return nil, err
	}
	return crypto.Decrypt(ciphertext)
}

var _ Crypto = (*baseCrypto)(nil)
