package filex

import (
	"os"
	"path/filepath"
	"strings"
)

// IsDir checks if the given path is a directory.
func IsDir(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}

// IfIsDir executes the given functionx if the given path is a directory.
// an error is returned if execution fails.
func IfIsDir(path string, f func() error) error {
	if IsDir(path) {
		return f()
	}
	return nil
}

func IsEmptyDir(path string) bool {
	if !IsDir(path) {
		return false
	}
	files, err := os.ReadDir(path)
	if err != nil {
		return false
	}

	return len(files) == 0
}

// DeleteDir deletes a directory.
func DeleteDir(dir string) error {
	return os.RemoveAll(dir)
}

// DeleteDirFn executes the functionx if the directory is deleted.
func DeleteDirFn(dir string, fn func() error) error {
	if err := DeleteDir(dir); err != nil {
		return err
	}
	return fn()
}

// IsSubDir checks if the given path is a subdirectory of the given parent.
func IsSubDir(parent, child string) bool {
	relPath, err := filepath.Rel(parent, child)
	if err != nil {
		return false
	}

	if strings.Contains(parent, child) {
		n := strings.ReplaceAll(child, parent, "")
		return filepath.Join(parent, relPath) == filepath.Join(parent, n)
	}

	return false
}
