package filex

import "io/ioutil"

// WhenFileRead is a functionx that executes a functionx when a file is read.
func WhenFileRead(filePath string, callback func([]byte) error) error {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	return callback(file)
}
