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

	require.Equal(t, "*file.SplitFile", reflect.TypeOf(f).String())
}

func TestNewSplitFileBySize(t *testing.T) {
	f, err := NewSplitFileBySize(1000)
	require.Zero(t, err)

	require.Equal(t, reflect.TypeOf(f).String(), "*file.SplitFile")
}

func TestSplitFile_SplitFileByChunkNum1(t *testing.T) {
	type fields struct {
		size uint64
	}
	type args struct {
		chunkNum uint
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []Chunk
		wantErr bool
	}{
		//  TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &SplitFile{
				size: tt.fields.size,
			}
			got, err := c.SplitFileByChunkNum(tt.args.chunkNum)
			if (err != nil) != tt.wantErr {
				t.Errorf("SplitFileByChunkNum() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SplitFileByChunkNum() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSplitFile_SplitFileByChunkSize1(t *testing.T) {
	type fields struct {
		size uint64
	}
	type args struct {
		chunkSize uint
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []Chunk
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &SplitFile{
				size: tt.fields.size,
			}
			got, err := c.SplitFileByChunkSize(tt.args.chunkSize)
			if (err != nil) != tt.wantErr {
				t.Errorf("SplitFileByChunkSize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SplitFileByChunkSize() got = %v, want %v", got, tt.want)
			}
		})
	}
}
