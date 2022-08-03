package cryptox

import (
	"fmt"
	"github.com/aesoper101/go-utils/cryptox/mode"
	"github.com/aesoper101/go-utils/cryptox/padding"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	c, err := NewAES(
		WithKey([]byte("0123456789abcdef")),
		WithIV([]byte("0123456789abcdef")),
		WithMode(mode.ECB),
		WithPadding(padding.NO),
	)
	require.Zero(t, err)
	require.Equal(t, "*cryptox.baseCrypto", reflect.TypeOf(c).String())
}

func Test_aesCrypto_Encrypt(t *testing.T) {
	c, err := NewAES(
		WithKey([]byte("0123456789abcdef")),
		WithIV([]byte("0123456789abcdef")),
		WithMode(mode.ECB),
		WithPadding(padding.NO),
	)
	require.Zero(t, err)

	_, err = c.Encrypt([]byte("0123456789abcdef"))
	require.Zero(t, err)
}

func Test_aesCrypto_Decrypt(t *testing.T) {
	c, err := NewAES(
		WithKey([]byte("0123456789abcdef")),
		WithIV([]byte("0123456789abcdef")),
		WithMode(mode.ECB),
		WithPadding(padding.NO),
	)
	require.Zero(t, err)

	result, err := c.Encrypt([]byte("0123456789abcdef"))
	require.Zero(t, err)

	fmt.Println(result.ToBase64())

	r2, err := c.Decrypt(result.ToBytes())
	require.Zero(t, err)
	require.Equal(t, "0123456789abcdef", string(r2))
}
