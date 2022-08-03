package filex

import (
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func TestSplitFile_SplitFileByChunkNum(t *testing.T) {
	f, err := NewSplitFileBySize(1000)
	require.Zero(t, err)

	r, err := f.SplitFileByChunkNum(20)
	require.Zero(t, err)

	require.Equal(t, 20, len(r))
}

func TestSplitFile_SplitFileByChunkSize(t *testing.T) {
	f, err := NewSplitFileBySize(1000)
	require.Zero(t, err)

	r, err := f.SplitFileByChunkSize(20)
	require.Zero(t, err)

	require.Equal(t, 50, len(r))
}

func TestNewSplitFile(t *testing.T) {
	f, err := NewSplitFile("./testdata/test.txt")
	require.Zero(t, err)

	require.Equal(t, "*filex.SplitFile", reflect.TypeOf(f).String())
}

func TestNewSplitFileBySize(t *testing.T) {
	f, err := NewSplitFileBySize(1000)
	require.Zero(t, err)

	require.Equal(t, reflect.TypeOf(f).String(), "*filex.SplitFile")
}
