package httpx

import (
	"github.com/aesoper101/go-utils/filex"
	"github.com/stretchr/testify/require"
	"testing"
)

var testUrl = "https://github.com/protocolbuffers/protobuf/releases/download/v3.19.4/protoc-3.19.4-win64.zip"

func TestDownload(t *testing.T) {
	err := Download(testUrl, "testdata/protoc.zip")

	require.Zero(t, err)
	require.Equal(t, true, filex.IsExists("testdata/protoc.zip"))
}

func TestDownloadWithProgress(t *testing.T) {
	err := DownloadWithProgress(testUrl, "testdata/protoc.zip")

	require.Zero(t, err)
	require.Equal(t, true, filex.IsExists("testdata/protoc.zip"))
}
