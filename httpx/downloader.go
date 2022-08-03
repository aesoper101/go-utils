package httpx

import (
	"github.com/aesoper101/go-utils/filex"
	"github.com/schollz/progressbar/v3"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// Download downloads a file from a url and saves it to a file, an error is returned if something goes wrong
func Download(url, filename string, requestFns ...func(request *http.Request)) error {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	for _, fn := range requestFns {
		fn(request)
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	return filex.CreateFileFromBytes(filename, false, body)
}

// DownloadWithProgress downloads a file from a url and saves it to a file, an error is returned if something goes wrong.
func DownloadWithProgress(url, filename string, requestFns ...func(request *http.Request)) error {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	for _, fn := range requestFns {
		fn(request)
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	_ = filex.MkdirIfNotExist(filepath.Dir(filename))

	file, _ := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0600)
	defer func() {
		_ = file.Close()
	}()

	bar := progressbar.DefaultBytes(
		response.ContentLength,
		"downloading "+strings.Split(filepath.Base(filename), ".")[0],
	)

	_, err = io.Copy(io.MultiWriter(file, bar), response.Body)

	return err
}
