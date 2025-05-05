package nodes

import "strings"

func CleanQuotes(s string) string {
	split := strings.Split(s, "\n")
	var fixed []string

	for _, item := range split {
		trimmed := strings.TrimLeft(item, "> ")
		fixed = append(fixed, trimmed)
	}
	joined := strings.Join(fixed, "\n")

	return joined
}
