package blocks

import (
	"strings"

	"github.com/Daxin319/SSGo/src/inline"
	"github.com/Daxin319/SSGo/src/nodes"
)

func CleanLists(block string) []nodes.TextNode {
	lines := strings.Split(block, "\n")
	var newNodes []nodes.TextNode

	for _, line := range lines {
		trimmed := strings.TrimSpace(line) // trim whitespace and skip if empty after trimming
		if trimmed == "" {
			continue
		}

		content := stripListMarker(trimmed) // strip list markers, convert to child nodes, map HTML and append to nodes slice to be returned
		children := inline.TextToChildren(content)
		children = nodes.MapToHTMLChildren(children, 0)

		newNodes = append(newNodes, nodes.TextNode{
			Tag:      "li",
			Children: children,
		})
	}

	return newNodes
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
