package nodes

import (
	"regexp"
	"strings"
)

func TextToTextNodes(s string) []TextNode {

	node := TextNode{
		Text:     s,
		TextType: Text,
		Url:      "",
	}
	pass1 := SplitTextHelper([]TextNode{node})
	pass2 := SplitTextHelper(pass1)
	pass3 := SplitTextHelper(pass2)
	pass4 := SplitTextHelper(pass3)
	images, _ := SplitNodesImage(pass4)
	final, _ := SplitNodesLink(images)

	return final
}

var reg1 = regexp.MustCompile("`|\\*\\*\\*")
var boldReg = regexp.MustCompile("\\*\\*")
var italicReg = regexp.MustCompile("\\*|_")

func findFirstDelimiter(s string) string {
	var match []int
	match = reg1.FindStringIndex(s)
	if len(match) == 0 {
		match = boldReg.FindStringIndex(s)
		if len(match) == 0 {
			match = italicReg.FindStringIndex(s)
			if len(match) == 0 {
				return s
			}
		}
	}
	if len(match) > 1 {
		return string(s[match[0]:match[len(match)-1]])
	}
	return string(s[match[0]])
}

func SplitTextHelper(oldNodes []TextNode) []TextNode {
	var results []TextNode

	for _, node := range oldNodes {
		if node.TextType == Code {
			results = append(results, node)
			continue
		}

		delim := findFirstDelimiter(node.Text)

		switch delim {
		case "`":
			code, _ := SplitNodesDelimiter(oldNodes, delim, Code)
			return code
		case `***`:
			boldtalic, _ := SplitNodesDelimiter(oldNodes, delim, Boldtalic)
			return boldtalic
		case `**`:
			bold, _ := SplitNodesDelimiter(oldNodes, delim, Bold)
			return bold
		case "_", `*`:
			italic, _ := SplitNodesDelimiter(oldNodes, delim, Italic)
			return italic
		default:
			if len(node.Text) == 0 {
				continue
			}
			if string(node.Text[0]) != ">" {
				results = append(results, node)
				continue
			}
			split := strings.Split(node.Text, "\n")
			var fixed []string

			for _, item := range split {
				if len(item) == 0 || item == " " || item == "\n" {
					continue
				}
				trimmed := strings.TrimLeft(item, "> ")
				fixed = append(fixed, trimmed)
			}
			joined := strings.Join(fixed, "\n")

			newNode := TextNode{
				Text:     joined,
				TextType: Text,
				Url:      "",
			}
			results = append(results, newNode)
		}
	}
	return results
}
