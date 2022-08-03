package padding

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestZeroPadding_Padding(t *testing.T) {
	p := zeroPadding{}
	padding, err := p.Padding([]byte("123456789aasd"), 10)
	require.Zero(t, err)
	require.Equal(t, 20, len(padding))

	fmt.Println(string(padding))
}

func TestZeroPadding_UnPadding(t *testing.T) {
	p := zeroPadding{}
	padding, err := p.Padding([]byte("123456789aasd"), 10)
	require.Zero(t, err)

	data, err := p.UnPadding(padding, 10)
	require.Zero(t, err)
	require.Equal(t, []byte("123456789aasd"), data)
}
