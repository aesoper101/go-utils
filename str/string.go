package str

import (
	"fmt"
	"strings"
)

// ToString converts any value to a string.
func ToString(v interface{}) string {
	switch v := v.(type) {
	case string:
		return v
	case []byte:
		return string(v)
	case error:
		return v.Error()
	case fmt.Stringer:
		return v.String()
	default:
		return fmt.Sprint(v)
	}
}

// LastFromSplit returns the last part of a string split by a separator.
func LastFromSplit(input, split string) string {
	rel := strings.Split(input, split)
	return rel[len(rel)-1]
}

// InArray checks if a string is in a slice of haystack.
func InArray[T comparable](needle T, haystack []T) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}
	return false
}

// Substring returns a substring of a string.
func Substring(s string, start int, end int) (string, error) {
	runes := []rune(s)
	if start < 0 || start > len(s) {
		return "", fmt.Errorf("start index out of range")
	}

	if end < 0 || end > len(s) {
		return "", fmt.Errorf("end index out of range")
	}

	return string(runes[start:end]), nil
}

// IsEmpty checks if a string is empty.
func IsEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

// IsNotEmpty checks if a string is not empty.
func IsNotEmpty(s string) bool {
	return !IsEmpty(s)
}
