package filex

import "path/filepath"

// MatchPattern returns true if the filepath matches any of the patterns.
func MatchPattern(patterns []string, name string) bool {
	for _, pat := range patterns {
		matched, _ := filepath.Match(pat, name)
		if matched {
			return true
		}
	}
	return false
}
