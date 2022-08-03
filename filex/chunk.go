package filex

import (
	"errors"
	"os"
)

type Chunk struct {
	Number uint   // Chunk number
	Offset uint64 // Chunk offset
	Size   uint64 // Chunk size.
}

type SplitFile struct {
	size uint64
}

// NewSplitFile creates a new SplitFile.
func NewSplitFile(path string) (*SplitFile, error) {
	if IsNotExists(path) {
		return nil, errors.New("file not exists")
	}
	if !IsFile(path) {
		return nil, errors.New("not a file")
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func(fp *os.File) {
		_ = fp.Close()
	}(file)

	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}

	size := uint64(stat.Size())
	return NewSplitFileBySize(size)
}

// NewSplitFileBySize create a new SplitFile by size.
func NewSplitFileBySize(size uint64) (*SplitFile, error) {
	c := &SplitFile{}
	c.size = size
	return c, nil
}

// SplitFileByChunkNum splits a file into chunks by chunk number.
func (c *SplitFile) SplitFileByChunkNum(chunkNum uint) ([]Chunk, error) {
	if c.size == 0 {
		return nil, errors.New("file size is zero")
	}
	if chunkNum == 0 {
		return nil, errors.New("chunk num is zero")
	}
	chunks := make([]Chunk, chunkNum)
	chunkSize := c.size / uint64(chunkNum)
	for i := uint(0); i < chunkNum; i++ {
		chunks[i] = Chunk{
			Number: i,
			Offset: uint64(i) * chunkSize,
			Size:   chunkSize,
		}
	}
	return chunks, nil
}

// SplitFileByChunkSize splits a file into chunks by chunk size.
func (c *SplitFile) SplitFileByChunkSize(chunkSize uint) ([]Chunk, error) {
	if c.size == 0 {
		return nil, errors.New("file size is zero")
	}

	if chunkSize == 0 {
		return nil, errors.New("chunk size is zero")
	}

	if c.size < uint64(chunkSize) {
		return nil, errors.New("file size is less than chunk size")
	}

	chunks := make([]Chunk, 0)
	for i := uint64(0); i < c.size; i += uint64(chunkSize) {
		chunks = append(chunks, Chunk{
			Number: uint(i / uint64(chunkSize)),
			Offset: i,
			Size:   uint64(chunkSize),
		})
	}

	return chunks, nil
}
