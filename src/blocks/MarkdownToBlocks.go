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
	inCodeBlock := false
	var codeFence string

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

		// Handle fenced code blocks
		if !inCodeBlock {
			if strings.HasPrefix(trimmed, "```") || strings.HasPrefix(trimmed, "~~~") {
				flush()
				inCodeBlock = true
				codeFence = trimmed[:3]
				current = append(current, line)
				continue
			}
		} else {
			current = append(current, line)
			if strings.HasPrefix(trimmed, codeFence) {
				inCodeBlock = false
				flush()
			}
			continue
		}

		if trimmed == "" {
			flush()
			continue
		}

		current = append(current, line)
	}
	flush()
	return blocks
}
