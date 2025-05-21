package nodes

import "strings"

func SanitizeNulls(s string) string {
	return strings.Map(func(r rune) rune {
		if r == 0 {
			return '\uFFFD'
		}
		return r
	}, s)
}
