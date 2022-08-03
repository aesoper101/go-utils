package randomx

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRandomString(t *testing.T) {
	r := RandomString(10)
	require.Equal(t, 10, len(r))
	fmt.Println(r)

	r1 := RandomString(20)
	require.Equal(t, 20, len(r1))
	fmt.Println(r1)
}
