package mode

import (
	"crypto/cipher"
	"github.com/aesoper101/go-utils/cryptox/padding"
)

type CryptoMode interface {
	SetPadding(padding padding.Padding)
	SetIV(iv []byte)
	SetKey(key []byte)
	SetBlock(block cipher.Block)
	GetPadding() padding.Padding
	GetIV() []byte
	GetKey() []byte
	GetBlock() cipher.Block
	Encrypt(plaintext []byte) ([]byte, error)
	Decrypt(ciphertext []byte) ([]byte, error)
}

type baseCryptoMode struct {
	padding padding.Padding
	iv      []byte
	key     []byte
	block   cipher.Block
}

func (b *baseCryptoMode) GetPadding() padding.Padding {
	return b.padding
}

func (b *baseCryptoMode) GetIV() []byte {
	return b.iv
}

func (b *baseCryptoMode) GetKey() []byte {
	return b.key
}

func (b *baseCryptoMode) GetBlock() cipher.Block {
	return b.block
}

func (b *baseCryptoMode) SetPadding(padding padding.Padding) {
	b.padding = padding
}

func (b *baseCryptoMode) SetIV(iv []byte) {
	b.iv = iv
}

func (b *baseCryptoMode) SetKey(key []byte) {
	b.key = key
}

func (b *baseCryptoMode) SetBlock(block cipher.Block) {
	b.block = block
}

func (b *baseCryptoMode) Encrypt(plaintext []byte) ([]byte, error) {
	return nil, nil
}

func (b *baseCryptoMode) Decrypt(ciphertext []byte) ([]byte, error) {
	return nil, nil
}

var _ CryptoMode = &baseCryptoMode{}
