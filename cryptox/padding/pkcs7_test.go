package padding

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPKCS7_Padding(t *testing.T) {
	p := pkcs7Padding{}
	padding, err := p.Padding([]byte("123456789aasd"), 10)
	require.Zero(t, err)
	require.Equal(t, 20, len(padding))
}

func TestPKCS7_UnPadding(t *testing.T) {
	p := pkcs7Padding{}
	padding, err := p.Padding([]byte("123456789aasd"), 10)
	require.Zero(t, err)

	data, err := p.UnPadding(padding, 10)
	require.Zero(t, err)
	require.Equal(t, []byte("123456789aasd"), data)
}
