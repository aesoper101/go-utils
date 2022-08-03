package contextx

import "context"

// NewContext returns a new Context that carries value.
func NewContext[K comparable, V any](ctx context.Context, key K, value V) context.Context {
	return context.WithValue(ctx, key, value)
}

// FromContext returns the value associated with the given key in the context, or nil if no value is associated with the key.
func FromContext[K comparable, V any](ctx context.Context, key K) (V, bool) {
	v, ok := ctx.Value(key).(V)
	return v, ok
}
