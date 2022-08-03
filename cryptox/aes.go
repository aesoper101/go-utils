package cryptox

import (
	"crypto/aes"
	"github.com/aesoper101/go-utils/cryptox/mode"
	"github.com/aesoper101/go-utils/cryptox/padding"
)

func NewAES(fns ...OptionFn) (Crypto, error) {
	opts := options{
		method:  AES,
		mode:    mode.CBC,
		keySize: KeySize128,
		padding: padding.NewNoPadding(),
	}
	for _, fn := range fns {
		fn(&opts)
	}

	if err := opts.Check(); err != nil {
		return nil, err
	}

	opts.key = PadZeroKey(opts.key, opts.keySize.Int())
	opts.iv = PadZeroKey(opts.iv, aes.BlockSize)

	return &baseCrypto{opts: opts}, nil
}
