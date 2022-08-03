package unicodex

import "unicode"

// IsChinese returns true if the rune is a Chinese character.
func IsChinese(args string) bool {
	set := []*unicode.RangeTable{unicode.Han, unicode.P}
	for _, r := range args {
		if unicode.IsOneOf(set, r) {
			return true
		}
	}
	return false
}
