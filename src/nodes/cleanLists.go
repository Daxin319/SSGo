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
		trim1 := strings.TrimLeft(item, "-* ")

		if re.MatchString(item) {
			wsIdx := bytes.Index([]byte(trim1), []byte(" "))
			trimmed = trim1[wsIdx+1:]
		} else {
			trimmed = strings.TrimLeft(item, "-* ")
		}

		child := HTMLNode{
			Tag:   "li",
			Value: trimmed,
		}

		children = append(children, child)
	}

	return children
}
