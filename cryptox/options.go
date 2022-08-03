package cryptox

import (
	"crypto/cipher"
	"github.com/aesoper101/go-utils/cryptox/mode"
	"github.com/aesoper101/go-utils/cryptox/padding"
)

type options struct {
	mode    mode.Mode
	key     []byte
	iv      []byte
	padding padding.Padding
	keySize KeySize
	method  Method
}

type OptionFn func(*options)

func (o *options) GetBlock() (cipher.Block, error) {
	return newBlock(o.method, o.key)
}

func (o *options) Check(fns ...func(*options) error) error {
	if o.method.Is(AES) {
		if NotAesValidKeySize(o.keySize) {
			return ErrAesKeySize
		}
		if o.mode.Is(mode.GCM) {
			return ErrAesNotSupportedGcm
		}
	} else if o.method.Is(DES) {
		if o.keySize.Not(KeySize64) {
			return ErrDesKeySize
		}
		if o.mode.Is(mode.GCM) {
			return ErrDesNotSupportedGcm
		}
	} else if o.method.Is(DES3) {
		if o.keySize.Not(KeySize192) {
			return ErrDesKeySize
		}
		if o.mode.Is(mode.GCM) {
			return ErrDes3NotSupportedGcm
		}
	}
	for _, fn := range fns {
		if err := fn(o); err != nil {
			return err
		}
	}
	return nil
}

func WithMode(m mode.Mode) OptionFn {
	return func(o *options) {
		o.mode = m
	}
}

func WithKey(key []byte) OptionFn {
	return func(o *options) {
		o.key = key
	}
}

func WithIV(iv []byte) OptionFn {
	return func(o *options) {
		o.iv = iv
	}
}

func WithPadding(p padding.Mode) OptionFn {
	return func(o *options) {
		if p2, err := padding.New(p.String()); err == nil {
			o.padding = p2
			return
		}

		o.padding = padding.NewNoPadding()
	}
}

func WithKeySize(keySize KeySize) OptionFn {
	return func(o *options) {
		o.keySize = keySize
	}
}

func WithMethod(method Method) OptionFn {
	return func(o *options) {
		o.method = method
	}
}
