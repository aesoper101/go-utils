package golangx

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go/parser"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

var (
	testSrc = `
package foo

import (
	"fmt"
	"time"
    _ "github.com/stretchr/testify/require"
)

func bar() {
	fmt.Println(time.Now())
}`
)

func TestCollectImports(t *testing.T) {
	p := MustParser(NewParser("", testSrc, parser.ParseComments))

	imports := CollectImports(p)

	require.Equal(t, 3, len(imports))

	for _, v := range imports {
		t.Logf("%s %s \n", v.Path, v.Name)
	}
}

func TestRemoveUnusedImports(t *testing.T) {
	p := MustParser(NewParser("", testSrc, parser.ParseComments))

	RemoveUnusedImports(p)

}

func TestIsInvalidPackageName(t *testing.T) {
	type args struct {
		pkg string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "invalid package name",
			args: args{
				pkg: "foo.bar",
			},
			want: true,
		},
		{
			name: "valid package name",
			args: args{
				pkg: "foo",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsInvalidPackageName(tt.args.pkg); got != tt.want {
				t.Errorf("IsInvalidPackageName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSanitizePackageName(t *testing.T) {
	type args struct {
		pkg string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "invalid package name",
			args: args{
				pkg: "foo.bar",
			},
			want: "foo_bar",
		},
		{
			name: "valid package name",
			args: args{
				pkg: "foo",
			},
			want: "foo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SanitizePackageName(tt.args.pkg); got != tt.want {
				t.Errorf("SanitizePackageName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNameForDir(t *testing.T) {
	wd, err := os.Getwd()
	require.NoError(t, err)

	require.Equal(t, "tmp", NameForDir("/tmp"))
	require.Equal(t, "golangx", NameForDir(wd))
	require.Equal(t, "go_utils", NameForDir(wd+"/.."))
	require.Equal(t, "foo", NameForDir("bar/foo"))
	require.Equal(t, "bar", NameForDir("bar/foo/.."))
	require.Equal(t, "bar_foo", NameForDir("bar.foo"))
}

func TestImportPathForDir(t *testing.T) {
	wd, err := os.Getwd()

	require.NoError(t, err)

	assert.Equal(t, "github.com/aesoper101/go-utils/golangx", ImportPathForDir(wd))
	assert.Equal(t, "github.com/aesoper101/go-utils/filex", ImportPathForDir(filepath.Join(wd, "..", "filex")))
	//
	//// doesnt contain go code, but should still give a valid import path
	assert.Equal(t, "github.com/aesoper101/go-utils", ImportPathForDir(filepath.Join(wd, "..")))
	//
	//// directory does not exist
	assert.Equal(t, "github.com/aesoper101/go-utils/dos", ImportPathForDir(filepath.Join(wd, "..", "dos")))
	//
	//// out of module
	assert.Equal(t, "", ImportPathForDir(filepath.Join(wd, "..", "..", "..")))
	//
	if runtime.GOOS == "windows" {
		assert.Equal(t, "", ImportPathForDir("C:/doesnotexist"))
	} else {
		assert.Equal(t, "", ImportPathForDir("/doesnotexist"))
	}
}
