package str

import (
	"github.com/iancoleman/strcase"
	"unicode"
)

// UcFirst returns a copy of s with each word capitalized.
func UcFirst(s string) string {
	if s == "" {
		return ""
	}
	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}

// LcFirst returns a copy of s with each word lowercased.
func LcFirst(s string) string {
	if s == "" {
		return ""
	}

	r := []rune(s)
	r[0] = unicode.ToLower(r[0])
	return string(r)
}

func ToCamel(s string) string {
	return strcase.ToCamel(s)
}

func ToLowerCamel(s string) string {
	return strcase.ToLowerCamel(s)
}

func ToSnake(s string) string {
	return strcase.ToSnake(s)
}

func ToKebab(s string) string {
	return strcase.ToKebab(s)
}

func ToScreamingKebab(s string) string {
	return strcase.ToScreamingKebab(s)
}

func ToDelimited(s string, delim uint8) string {
	return strcase.ToDelimited(s, delim)
}

func ToScreamingDelimited(s string, delimiter uint8, ignore string, screaming bool) string {
	return strcase.ToScreamingDelimited(s, delimiter, ignore, screaming)
}

func ToSnakeWithIgnore(s string, ignore string) string {
	return strcase.ToSnakeWithIgnore(s, ignore)
}
