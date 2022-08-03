package filex

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
)

// CopyFile copies a file from src to dst. returns an error if something goes wrong.
// if overwrite slice is not empty, the first one will be used to determine if the file should be overwritten.
func CopyFile(src, dst string, overwrite ...bool) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	canOverwrite := false
	if len(overwrite) > 0 {
		canOverwrite = overwrite[0]
	}

	return CreateFileFromReader(dst, canOverwrite, in)
}

// CopyDir copies a directory from src to dst. returns an error if something goes wrong.
// if overwrite slice is not empty, the first one will be used to determine if the file should be overwritten.
func CopyDir(src, dst string, overwrite ...bool) error {
	if IsNotExists(src) {
		return errors.New("src dir not exists")
	}
	if !IsDir(src) {
		return errors.New("src is not a directory")
	}

	if IsSubDir(src, dst) {
		return errors.New("dst is a subdir of src")
	}

	err := filepath.WalkDir(src, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel("testdata", path)
		if err != nil {
			return err
		}

		dstPath := filepath.Join(dst, relPath)
		if d.IsDir() {
			return MkdirIfNotExist(dstPath)
		}

		if IsSymlink(path) {
			symlinkPath, err := filepath.EvalSymlinks(path)
			if err != nil {
				return err
			}

			info, err := os.Lstat(symlinkPath)
			if err != nil {
				return err
			}

			if info.IsDir() {
				return CopyDir(symlinkPath, dstPath)
			}

			return CopyFile(symlinkPath, dstPath, overwrite...)
		}

		return CopyFile(path, dstPath, overwrite...)
	})

	return err
}
