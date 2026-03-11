package blocks

import (
	"strings"
)

// strip quote block markdown from each line and replace with <br> tag

func CleanQuotes(s string) string {
	split := strings.Split(s, "\n")
	var fixed []string

	for _, item := range split {
		trimmed := item
		if strings.HasPrefix(item, "> ") {
			trimmed = item[2:]
		} else if strings.HasPrefix(item, ">") {
			trimmed = item[1:]
		}
		fixed = append(fixed, trimmed)
	}
	return strings.Join(fixed, "\n")
}
