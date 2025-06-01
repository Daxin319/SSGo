package blocks

import (
	"regexp"
	"strings"
)

var rege = regexp.MustCompile(`^\s*$`)

func MarkdownToBlocks(s string) []string {
	split := strings.Split(s, "\n\n")
	final := []string{}

	for _, block := range split {
		lBlock := strings.TrimLeft(block, " ")
		fBlock := strings.TrimRight(lBlock, " ")
		if rege.MatchString(fBlock) {
			continue
		}
		final = append(final, fBlock)
	}
	return final
}
