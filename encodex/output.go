package encodex

import (
	"encoding/base32"
	"encoding/base64"
	"encoding/hex"
)

type Encode struct {
	input []byte
}

func FromBytes(input []byte) *Encode {
	return &Encode{input: input}
}

func FromString(input string) *Encode {
	return &Encode{input: []byte(input)}
}

func (o *Encode) ToHex() string {
	return hex.EncodeToString(o.input)
}

func (o *Encode) ToBase64() string {
	return base64.StdEncoding.EncodeToString(o.input)
}

func (o *Encode) ToBase32() string {
	return base32.StdEncoding.EncodeToString(o.input)
}

func (o *Encode) ToBytes() []byte {
	return o.input
}
