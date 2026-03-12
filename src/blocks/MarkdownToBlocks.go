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

	for i, line := range lines {
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
			// Peek ahead to see if the next non-empty line is indented.
			// This allows list items to span multiple blocks.
			isContinuation := false
			for j := i + 1; j < len(lines); j++ {
				nextTrimmed := strings.TrimSpace(lines[j])
				if nextTrimmed == "" {
					continue
				}
				// If the next non-empty line is indented by at least 2 spaces,
				// it's a continuation of the current block.
				if strings.HasPrefix(lines[j], "  ") {
					isContinuation = true
				}
				break
			}
			if !isContinuation {
				flush()
				continue
			}
		}

		current = append(current, line)
	}
	flush()
	return blocks
}
