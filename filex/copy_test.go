package filex

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCopyDir(t *testing.T) {
	err := CopyDir("testdata", "testdata_copy")
	require.Zero(t, err)

	require.Equal(t, true, IsExists("testdata"))
}
