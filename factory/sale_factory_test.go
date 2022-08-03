package factory

import (
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func TestNewSafeFactory(t *testing.T) {
	f := NewSafeFactory[testdata]()
	require.NotNil(t, f)
	require.Equal(t, "*factory.safeFactory[factory.testdata]", reflect.TypeOf(f).String())
}

func Test_safeFactory_Get(t *testing.T) {
	f := NewSafeFactory[testdata]()
	require.NotNil(t, f)

	err := f.Register("test", func() (testdata, error) {
		return testdata{}, nil
	})
	require.Zero(t, err)

	r, err := f.Get("test")
	require.Zero(t, err)

	data, err := r()
	require.Zero(t, err)
	require.Equal(t, "factory.testdata", reflect.TypeOf(data).String())
}

func Test_safeFactory_Register(t *testing.T) {
	f := NewSafeFactory[testdata]()
	require.NotNil(t, f)

	err := f.Register("test", func() (testdata, error) {
		return testdata{}, nil
	})
	require.Zero(t, err)
}

func Test_safeFactory_Unregister(t *testing.T) {
	f := NewSafeFactory[testdata]()
	require.NotNil(t, f)

	err := f.Register("test", func() (testdata, error) {
		return testdata{}, nil
	})
	require.Zero(t, err)

	r, err := f.Get("test")
	require.Zero(t, err)
	require.NotNil(t, r)

	err = f.Unregister("test")
	require.Zero(t, err)

	r, err = f.Get("test")
	require.NotZero(t, err)
	require.Nil(t, r)
}
