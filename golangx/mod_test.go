package golangx

import "testing"

func TestGoModName(t *testing.T) {
	t.Logf("%s", GoModName())
}

func TestGoModPath(t *testing.T) {
	t.Logf("%s", GoModPath())
}

func TestGoModCachePath(t *testing.T) {
	t.Logf("%s", GoModCachePath())
}

func TestGoProxy(t *testing.T) {
	t.Logf("%s", GoProxy())
}
