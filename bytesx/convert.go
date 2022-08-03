package bytesx

import "unsafe"

// ToString converts a byte slice to a string.
func ToString(s []byte) string {
	return *(*string)(unsafe.Pointer(&s))
}
