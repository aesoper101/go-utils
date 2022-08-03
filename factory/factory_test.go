package factory

import (
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

type testdata struct {
	name string
	data interface{}
}

func TestNewFactory(t *testing.T) {
	f := NewFactory[testdata]()
	require.NotNil(t, f)
	require.Equal(t, "*factory.factory[factory.testdata]", reflect.TypeOf(f).String())
}

func Test_factory_Get(t *testing.T) {
	f := NewFactory[testdata]()
	require.NotNil(t, f)

	err := f.Register("test", func() testdata {
		return testdata{}
	})
	require.Zero(t, err)

	r, err := f.Get("test")
	require.Zero(t, err)

	data := r()
	require.Equal(t, "factory.testdata", reflect.TypeOf(data).String())
}

func Test_factory_Register(t *testing.T) {
	f := NewFactory[testdata]()
	require.NotNil(t, f)

	err := f.Register("test", func() testdata {
		return testdata{}
	})
	require.Zero(t, err)
}

func Test_factory_Unregister(t *testing.T) {
	f := NewFactory[testdata]()
	require.NotNil(t, f)

	err := f.Register("test", func() testdata {
		return testdata{}
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
