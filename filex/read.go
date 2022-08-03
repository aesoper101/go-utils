package filex

import (
	"os"
)

// WhenFileRead is a function that executes a function when a file is read.
func WhenFileRead(filename string, callback func([]byte) error) error {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	return callback(bytes)
}
