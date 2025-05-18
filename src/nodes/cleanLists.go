package nodes

import (
	"strings"
)

func CleanLists(block string) []TextNode {
	lines := strings.Split(block, "\n")
	var nodes []TextNode

	for _, line := range lines {
		trimmed := strings.TrimSpace(line) // trim whitespace and skip if empty after trimming
		if trimmed == "" {
			continue
		}

		content := stripListMarker(trimmed) // strip list markers, convert to child nodes, map HTML and append to nodes slice to be returned
		children := TextToChildren(content)
		children = mapToHTMLChildren(children, 0)

		nodes = append(nodes, TextNode{
			Tag:      "li",
			Children: children,
		})
	}

	return nodes
}

// FIX THIS. CAUSING PROBLEMS FOR LONG ORDERED LISTS
func stripListMarker(line string) string {
	if strings.HasPrefix(line, "- ") || strings.HasPrefix(line, "* ") {
		return line[2:]
	}
	for i := range len(line) {
		if line[i] >= '0' && line[i] <= '9' && i+1 < len(line) && line[i+1] == '.' {
			return strings.TrimSpace(line[i+2:])
		}
	}
	return line
}
