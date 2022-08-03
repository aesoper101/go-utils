package padding

var _ Padding = (*iso97971Padding)(nil)

type iso97971Padding struct{}

func NewISO97971Padding() Padding {
	return &iso97971Padding{}
}

func (i *iso97971Padding) Padding(plaintext []byte, blockSize int) ([]byte, error) {
	p := &zeroPadding{}
	return p.Padding(append(plaintext, 0x80), blockSize)
}

func (i *iso97971Padding) UnPadding(ciphertext []byte, blockSize int) ([]byte, error) {
	p := &zeroPadding{}
	data, err := p.UnPadding(ciphertext, blockSize)
	if err != nil {
		return nil, err
	}
	return data[:len(data)-1], nil
}
