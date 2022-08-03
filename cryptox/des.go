package cryptox

import (
	"crypto/des"
	"github.com/aesoper101/go-utils/cryptox/mode"
	"github.com/aesoper101/go-utils/cryptox/padding"
)

func NewDES(fns ...OptionFn) (Crypto, error) {
	opts := options{
		method:  DES,
		mode:    mode.CBC,
		keySize: KeySize64,
		padding: padding.NewNoPadding(),
	}
	for _, fn := range fns {
		fn(&opts)
	}

	if err := opts.Check(); err != nil {
		return nil, err
	}

	opts.key = PadZeroKey(opts.key, opts.keySize.Int())
	opts.iv = PadZeroKey(opts.iv, des.BlockSize)

	return &baseCrypto{opts: opts}, nil
}
