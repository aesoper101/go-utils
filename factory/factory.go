package factory

import (
	"fmt"
	"strings"
	"sync"
)

type buildFn[T any] map[string]func() T

type factory[T any] struct {
	rw sync.RWMutex
	m  buildFn[T]
}

func NewFactory[T any]() *factory[T] {
	return &factory[T]{
		m: make(buildFn[T]),
	}
}

// Register registers a new build functionx.
// The name is the key used to retrieve the build functionx. If the name is already registered, returns an error.
// The name is not case sensitive.
func (f *factory[T]) Register(name string, fn func() T) error {
	f.rw.Lock()
	defer f.rw.Unlock()

	name = strings.ToLower(name)

	if _, ok := f.m[name]; ok {
		return fmt.Errorf("factory: %s already registered", name)
	}

	f.m[name] = fn

	return nil
}

func (f *factory[T]) Unregister(name string) error {
	f.rw.Lock()
	defer f.rw.Unlock()

	name = strings.ToLower(name)

	if _, ok := f.m[name]; !ok {
		return fmt.Errorf("factory: %s not registered", name)
	}

	delete(f.m, name)
	return nil
}

func (f *factory[T]) Get(name string) (func() T, error) {
	f.rw.RLock()
	defer f.rw.RUnlock()

	name = strings.ToLower(name)
	fn, ok := f.m[name]
	if !ok {
		return nil, fmt.Errorf("factory: %s not registered", name)
	}

	return fn, nil
}
