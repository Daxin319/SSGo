package nodes

import (
	"strings"
)

func CleanLists(block string) []TextNode {
	lines := strings.Split(block, "\n")
	var nodes []TextNode

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" {
			continue
		}

		content := stripListMarker(trimmed)
		children := TextToChildren(content)
		children = mapToHTMLChildren(children, 0)

		nodes = append(nodes, TextNode{
			Tag:      "li",
			Children: children,
		})
	}

	return nodes
}

func stripListMarker(line string) string {
	if strings.HasPrefix(line, "- ") || strings.HasPrefix(line, "* ") {
		return line[2:]
	}
	for i := 0; i < len(line); i++ {
		if line[i] >= '0' && line[i] <= '9' && i+1 < len(line) && line[i+1] == '.' {
			return strings.TrimSpace(line[i+2:])
		}
	}
	return line
}
