package blocks

import (
	"regexp"
	"strings"
)

var rege = regexp.MustCompile(`^\s*$`)

func MarkdownToBlocks(s string) []string {
	lines := strings.Split(s, "\n")
	var blocks []string
	var current []string

	flush := func() {
		if len(current) > 0 {
			block := strings.Join(current, "\n")
			if !rege.MatchString(block) {
				blocks = append(blocks, block)
			}
			current = nil
		}
	}

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" {
			flush()
			continue
		}
		// Only treat standalone HTML tags as blocks
		if strings.HasPrefix(trimmed, "<") && strings.HasSuffix(trimmed, ">") &&
			(strings.HasPrefix(trimmed, "<div") || strings.HasPrefix(trimmed, "<!--")) {
			flush()
			blocks = append(blocks, trimmed)
			continue
		}
		current = append(current, line)
	}
	flush()
	return blocks
}
