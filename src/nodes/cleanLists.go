package nodes

import (
	"bytes"
	"regexp"
	"strings"
)

var re = regexp.MustCompile(`^\d+\. `)

func CleanLists(s string) []HTMLNode {
	var children []HTMLNode
	var trimmed string

	split := strings.Split(s, "\n")

	for _, item := range split {
		trim1 := strings.TrimLeft(item, " ")

		if re.MatchString(item) {
			wsIdx := bytes.Index([]byte(trim1), []byte(" "))
			trimmed = trim1[wsIdx+1:]
		} else {
			trim2 := strings.TrimPrefix(item, "* ")
			trimmed = strings.TrimPrefix(trim2, "- ")
		}
		oldNodes := MarkdownToHTMLNode(trimmed)

		child := HTMLNode{
			Tag:   "li",
			Value: oldNodes.ToHTML(),
		}

		children = append(children, child)
	}

	return children
}
