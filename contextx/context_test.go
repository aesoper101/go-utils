package contextx

import (
	"context"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

type testValue struct {
	Name string
}

type testKey struct {
	Name1 string
}

func TestFromContext(t *testing.T) {
	key := testKey{Name1: "test"}
	ctx := context.WithValue(context.Background(), key, testValue{Name: "test"})
	v, ok := FromContext[testKey, testValue](ctx, key)

	require.Equal(t, true, ok)
	require.Equal(t, "contextx.testValue", reflect.TypeOf(v).String())
	require.Equal(t, "test", v.Name)
}

func TestNewContext(t *testing.T) {
	ctx := NewContext[testKey, testValue](context.Background(), testKey{Name1: "test"}, testValue{Name: "test"})

	key := testKey{Name1: "test"}
	v, ok := FromContext[testKey, testValue](ctx, key)

	require.Equal(t, true, ok)
	require.Equal(t, "contextx.testValue", reflect.TypeOf(v).String())
	require.Equal(t, "test", v.Name)
}
