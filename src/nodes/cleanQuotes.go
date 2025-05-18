package nodes

import (
	"strings"
)

// strip quote block markdown from each line and replace with <br> tag

func CleanQuotes(s string) string {
	split := strings.Split(s, "\n")
	var fixed []string

	for _, item := range split {
		trimmed := strings.TrimLeft(item, "> ")
		if len(trimmed) == 0 {
			trimmed = "<br>"
		}
		fixed = append(fixed, trimmed)
	}
	joined := strings.Join(fixed, "<br>")

	return joined
}
