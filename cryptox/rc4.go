package cryptox

import (
	"crypto/rc4"
	"github.com/aesoper101/go-utils/encodex"
)

type RC4 struct {
	key    []byte
	cipher *rc4.Cipher
}

func NewRC4(key []byte) (Crypto, error) {
	cipher, err := rc4.NewCipher(key)
	if err != nil {
		return nil, err
	}

	return &RC4{
		key:    key,
		cipher: cipher,
	}, nil
}

func (r *RC4) Encrypt(plaintext []byte) (*encodex.Encode, error) {
	ciphertext := make([]byte, len(plaintext))
	r.cipher.XORKeyStream(ciphertext, plaintext)

	return encodex.FromBytes(ciphertext), nil
}

func (r *RC4) Decrypt(ciphertext []byte) (plaintext []byte, err error) {
	plaintext = make([]byte, len(ciphertext))
	r.cipher.XORKeyStream(plaintext, ciphertext)

	return plaintext, nil
}
