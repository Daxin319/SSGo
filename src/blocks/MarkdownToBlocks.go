package blocks

import (
	"regexp"
	"strings"
)

var rege = regexp.MustCompile(`^\s*$`)

// Unified autolink regex: matches protocol autolinks, email autolinks (with optional user:pass:), and GFM-style domain autolinks
var autolinkBlockRe = regexp.MustCompile(`^<((?:https?|ftp|ftps|sftp|ws|wss)://[^ >]+|(?:[A-Za-z0-9._%+\-]+:)?[A-Za-z0-9._%+\-]+@[A-Za-z0-9.\-]+\.[A-Za-z]{2,}|(?:[a-zA-Z0-9](?:[a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])?\.)+[a-zA-Z]{2,6}(?::[0-9]+)?(?:/[a-zA-Z0-9\-._~:/?#\[\]@!$&'()*+,;=%]*)?)>$`)

// Regex for HTML tags and comments
var htmlTagOrCommentRe = regexp.MustCompile(`^<(?:!--[\s\S]*?--|/?[a-zA-Z][a-zA-Z0-9-]*(?:\s+[^<>]*)?)>$`)

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
		if autolinkBlockRe.MatchString(trimmed) || htmlTagOrCommentRe.MatchString(trimmed) {
			flush()
			blocks = append(blocks, trimmed)
			continue
		}
		current = append(current, line)
	}
	flush()
	return blocks
}
