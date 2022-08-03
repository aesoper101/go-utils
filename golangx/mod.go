package golangx

import (
	"github.com/aesoper101/go-utils/str"
	"golang.org/x/mod/modfile"
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"strings"
)

// IsGoFile checks if the given file is a go file
func IsGoFile(path string) bool {
	return filepath.Ext(path) == ".go"
}

// GoModCachePath returns the path of Go module cache.
func GoModCachePath() string {
	cacheOut, _ := exec.Command("go", "env", "GOMODCACHE").Output()
	cachePath := strings.Trim(string(cacheOut), "\n")
	if str.IsEmpty(cachePath) {
		return filepath.Join(GoPath(), "pkg", "mod")
	}
	return cachePath
}

// GoModFilePath returns the full path to the go.mod file for the current project.
func GoModFilePath() string {
	modOut, _ := exec.Command("go", "env", "GOMOD").Output()
	modPath := strings.Trim(string(modOut), "\n")
	if len(modPath) == 3 && strings.ToLower(modPath) == "nil" {
		return ""
	}
	return modPath
}

// GoModPath returns the directory path of the go.mod file for the current project.
func GoModPath() string {
	modPath := GoModFilePath()
	if str.IsEmpty(modPath) {
		return ""
	}
	return filepath.Dir(modPath)
}

// IsGoModProject returns true if the current project is a Go module project.
func IsGoModProject() bool {
	return GoModFilePath() != ""
}

// GoModName returns the name of the current project without version.
func GoModName() string {
	if !IsGoModProject() {
		return ""
	}

	filename := GoModFilePath()
	data, _ := ioutil.ReadFile(filename)

	if f, _ := modfile.Parse(filename, data, nil); f != nil {
		return f.Module.Mod.Path
	}

	return ""
}

// RefPathToGoModPath returns the relative path to the go.mod file for the given path.
func RefPathToGoModPath(path string) string {
	refPath, _ := filepath.Rel(GoModPath(), path)
	return refPath
}
