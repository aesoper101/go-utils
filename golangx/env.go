package golangx

import (
	"github.com/aesoper101/go-utils/filex"
	"github.com/aesoper101/go-utils/str"
	"go/build"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var defaultGoProxy = "https://goproxy.cn,direct"

// GoBin returns the path to the go binary.
func GoBin() string {
	bin := os.Getenv("GOBIN")
	if bin != "" {
		if filex.IsExists(bin) {
			return bin
		}
	}

	gopath := GoPath()
	bin = filepath.Join(gopath, "bin")
	if filex.IsExists(bin) {
		return bin
	}

	goroot := GoRoot()
	bin = filepath.Join(goroot, "bin")
	if filex.IsExists(bin) {
		return bin
	}

	ctx := build.Default

	return filepath.Join(ctx.GOPATH, "bin")
}

// GoProxy returns the proxy setting for go. if not set,
// return a default proxy setting: "https://goproxy.cn,direct".
func GoProxy() string {
	proxy := os.Getenv("GOPROXY")
	if proxy != "" {
		return proxy
	}

	if output, _ := exec.Command("go", "env", "GOPROXY").Output(); str.IsNotEmpty(string(output)) {
		return string(output)
	}

	return defaultGoProxy
}

// IsGO111ModuleOn returns true if go111module is set to on.
func IsGO111ModuleOn() bool {
	v := os.Getenv("GO111MODULE")
	if str.IsNotEmpty(v) {
		return strings.ToLower(v) == "on"
	}

	if output, _ := exec.Command("go", "env", "GO111MODULE").Output(); str.IsNotEmpty(string(output)) {
		return strings.ToLower(string(output)) == "on"
	}

	return true
}

// GoPath returns the GOPATH.
func GoPath() string {
	path := os.Getenv("GOPATH")
	if str.IsNotEmpty(path) {
		return path
	}

	if output, _ := exec.Command("go", "env", "GOPATH").Output(); str.IsNotEmpty(string(output)) {
		return string(output)
	}

	return build.Default.GOPATH
}

// GoRoot returns the GOROOT.
func GoRoot() string {
	path := os.Getenv("GOROOT")
	if str.IsNotEmpty(path) {
		return path
	}

	if output, _ := exec.Command("go", "env", "GOROOT").Output(); str.IsNotEmpty(string(output)) {
		return string(output)
	}

	return build.Default.GOROOT
}

// GoVersion returns the go version.
func GoVersion() string {
	if output, _ := exec.Command("go", "env", "GOVERSION").Output(); str.IsNotEmpty(string(output)) {
		return strings.ReplaceAll(strings.Replace(string(output), "\n", "", -1), "go", "")
	}

	return ""
}
