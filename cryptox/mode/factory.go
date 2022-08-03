package mode

import (
	"github.com/aesoper101/go-utils/factory"
)

var fi = factory.NewFactory[CryptoMode]()

func init() {
	_ = fi.Register(ECB.String(), newECB)
	_ = fi.Register(CBC.String(), newCBC)
	_ = fi.Register(CTR.String(), newCTR)
	_ = fi.Register(CFB.String(), newCFB)
	_ = fi.Register(OFB.String(), newOFB)
	_ = fi.Register(GCM.String(), newGcm)
}

func GetCryptoMode(name string) (CryptoMode, error) {
	fn, err := fi.Get(name)
	if err != nil {
		return nil, err
	}
	return fn(), nil
}
