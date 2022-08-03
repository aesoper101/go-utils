package binaryx

import (
	"github.com/aesoper101/go-utils/filex"
	"github.com/aesoper101/go-utils/golangx"
	"path/filepath"
	"runtime"
)

// Install installs the binary package in the current directory. returns the path to the binary package.
// If the binary package is not installed, it returns an error.
func Install(binCachePath string, binaryName string, installFn func(dest string) (string, error)) (string, error) {
	goBin := golangx.GoBin()
	cacheFile := filepath.Join(binCachePath, binaryName)
	destBinFile := filepath.Join(goBin, binaryName)

	switch runtime.GOOS {
	case "windows":
		destBinFile += ".exe"
		cacheFile += ".exe"
	}

	// install binary from cache
	err := filex.CopyFile(cacheFile, destBinFile, true)
	if err == nil {
		return destBinFile, nil
	}

	// install binary from install functionx
	dest, err := installFn(destBinFile)
	if err != nil {
		return "", err
	}

	// copy binary to cache
	_ = filex.CopyFile(dest, cacheFile, true)

	return destBinFile, nil
}
