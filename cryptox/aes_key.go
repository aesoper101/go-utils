package cryptox

type KeySize int

const (
	KeySize128 KeySize = 16
	KeySize192 KeySize = 24
	KeySize256 KeySize = 32

	// KeySize64 is supported by Des, but not by AES.
	KeySize64 KeySize = 8
)

func (k KeySize) Int() int {
	return int(k)
}

func (k KeySize) Not(size KeySize) bool {
	return k != size
}

func NotAesValidKeySize(k KeySize) bool {
	return k != KeySize128 && k != KeySize192 && k != KeySize256
}
