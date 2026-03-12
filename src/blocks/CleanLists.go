package blocks

import (
	"strings"

	"github.com/Daxin319/SSGo/src/inline"
	"github.com/Daxin319/SSGo/src/nodes"
)

func CleanLists(block string, blockParser func(string) []nodes.TextNode) []nodes.TextNode {
	lines := strings.Split(block, "\n")
	var newNodes []nodes.TextNode
	var currentItem []string
	var sublines []string
	var subIndent int = -1
	isLoose := false

	flushItem := func() {
		if len(currentItem) > 0 {
			content := strings.Join(currentItem, " ")
			children := inline.TextToChildren(content)
			children = nodes.MapToHTMLChildren(children, 0)

			if isLoose && len(children) > 0 {
				children = []nodes.TextNode{{
					Tag:      "p",
					Children: children,
				}}
			}

			if len(sublines) > 0 {
				subBlock := strings.Join(sublines, "\n")
				subNodes := blockParser(subBlock)
				children = append(children, subNodes...)
			}

			newNodes = append(newNodes, nodes.TextNode{
				Tag:      "li",
				Children: children,
			})
			currentItem = nil
			sublines = nil
			subIndent = -1
			isLoose = false
		}
	}

	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		if trimmedLine == "" {
			isLoose = true
			if len(sublines) > 0 {
				sublines = append(sublines, "")
			}
			continue
		}
		trimmed := strings.TrimLeft(line, " ")
		indent := len(line) - len(trimmed)

		if indent == 0 && isListItemStart(trimmed) {
			flushItem()
			currentItem = append(currentItem, stripListMarker(trimmed))
		} else if indent >= 2 {
			// Use the indentation level of the first sub-line
			if subIndent == -1 {
				subIndent = indent
			}
			subline := line
			if len(subline) >= subIndent {
				subline = subline[subIndent:]
			} else {
				subline = strings.TrimLeft(subline, " ")
			}
			sublines = append(sublines, subline)
		} else {
			// This is a continuation of the current item's text
			currentItem = append(currentItem, strings.TrimSpace(line))
		}
	}
	flushItem()

	return newNodes
}

func isListItemStart(line string) bool {
	if strings.HasPrefix(line, "- ") || strings.HasPrefix(line, "* ") {
		return true
	}
	// Match 1. 2. etc.
	for i, r := range line {
		if r >= '0' && r <= '9' {
			continue
		}
		if r == '.' && i > 0 && i+1 < len(line) && line[i+1] == ' ' {
			return true
		}
		break
	}
	return false
}

func stripListMarker(line string) string {
	if strings.HasPrefix(line, "- ") || strings.HasPrefix(line, "* ") {
		return line[2:]
	}
	dotIdx := strings.Index(line, ". ")
	if dotIdx != -1 {
		// Check if everything before dot is digits
		isDigits := true
		for i := 0; i < dotIdx; i++ {
			if line[i] < '0' || line[i] > '9' {
				isDigits = false
				break
			}
		}
		if isDigits {
			return strings.TrimSpace(line[dotIdx+1:])
		}
	}
	return line
}
