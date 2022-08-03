package cryptox

type Method string

const (
	AES  = "AES"
	DES  = "DES"
	DES3 = "DES3"
)

func (m Method) String() string {
	return string(m)
}

func (m Method) Is(m1 Method) bool {
	return m == m1
}
