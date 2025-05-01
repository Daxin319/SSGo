package nodes

import "regexp"

func TextToTextNodes(s string) []TextNode {

	node := TextNode{
		Text:     s,
		TextType: Text,
		Url:      "",
	}

	pass1 := splitTextHelper([]TextNode{node})
	pass2 := splitTextHelper(pass1)
	pass3 := splitTextHelper(pass2)
	images, _ := SplitNodesImage(pass3)
	final, _ := SplitNodesLink(images)

	return final
}

var reg = regexp.MustCompile("`|\\*\\*|_")

func findFirstDelimiter(s string) string {
	match := reg.FindStringIndex(s)
	if len(match) == 0 {
		return s
	}
	if len(match) > 1 {
		return string(s[match[0]:match[len(match)-1]])
	}
	return string(s[match[0]])
}

func splitTextHelper(oldNodes []TextNode) []TextNode {
	var results []TextNode

	for _, node := range oldNodes {
		if node.TextType != Text {
			results = append(results, node)
		}

		delim := findFirstDelimiter(node.Text)

		switch delim {
		case "`":
			code, _ := SplitNodesDelimiter(oldNodes, delim, Code)
			return code
		case "**":
			bold, _ := SplitNodesDelimiter(oldNodes, delim, Bold)
			return bold
		case "_":
			italic, _ := SplitNodesDelimiter(oldNodes, delim, Italic)
			return italic
		default:
			results = append(results, node)
		}
	}
	return results
}
