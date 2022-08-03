package padding

import (
	f "github.com/aesoper101/go-utils/factory"
)

// Factory is a factory for creating padding instances.
var factory = f.NewFactory[Padding]()

func init() {
	_ = factory.Register(PKCS5.String(), NewPKCS5Padding)
	_ = factory.Register(PKCS7.String(), NewPKCS7Padding)
	_ = factory.Register(ZERO.String(), NewZeroPadding)
	_ = factory.Register(ANSIX923.String(), NewAnsiX923Padding)
	_ = factory.Register(ISO10126.String(), NewISO10126Padding)
	_ = factory.Register(ISO97971.String(), NewISO97971Padding)
	_ = factory.Register(NO.String(), NewNoPadding)
}

func New(name string) (Padding, error) {
	fn, err := factory.Get(name)
	if err != nil {
		return nil, err
	}

	return fn(), nil
}
