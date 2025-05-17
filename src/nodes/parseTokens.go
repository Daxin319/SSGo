package nodes

import "strings"

func parseTokens(tokens []token, pos int, delim string) ([]TextNode, int) {
	var nodes []TextNode
	for i := pos; i < len(tokens); {
		tok := tokens[i]
		if delim != "" && tok.kind == delim {
			return nodes, i
		}
		switch tok.kind {
		case strings.Repeat("*", 3), strings.Repeat("_", 3):
			children, j := parseTokens(tokens, i+1, tok.kind)
			italicNode := TextNode{TextType: Italic, Children: children}
			nodes = append(nodes, TextNode{TextType: Bold, Children: []TextNode{italicNode}})
			i = j + 1
		case strings.Repeat("*", 2), strings.Repeat("_", 2):
			children, j := parseTokens(tokens, i+1, tok.kind)
			nodes = append(nodes, TextNode{TextType: Bold, Children: children})
			i = j + 1
		case "*", "_":
			children, j := parseTokens(tokens, i+1, tok.kind)
			nodes = append(nodes, TextNode{TextType: Italic, Children: children})
			i = j + 1
		case "`":
			var sb strings.Builder
			j := i + 1
			for j < len(tokens) && tokens[j].kind != "`" {
				sb.WriteString(tokens[j].value)
				j++
			}
			nodes = append(nodes, TextNode{TextType: Code, Text: sb.String()})
			if j < len(tokens) && tokens[j].kind == "`" {
				i = j + 1
			} else {
				i = j
			}
		case "![", "[":
			isImage := tok.kind == "!["
			children, j1 := parseTokens(tokens, i+1, "]")
			var altOrText string
			for _, c := range children {
				altOrText += c.Text
			}
			if j1+3 < len(tokens) && tokens[j1+1].kind == "(" && tokens[j1+3].kind == ")" {
				url := tokens[j1+2].value
				if isImage {
					nodes = append(nodes, TextNode{TextType: Image, Url: url, Children: children})
				} else {
					nodes = append(nodes, TextNode{TextType: Link, Url: url, Children: children})
				}
				i = j1 + 4
			} else {
				nodes = append(nodes, TextNode{TextType: Text, Text: tok.value})
				i++
			}
		default:
			nodes = append(nodes, TextNode{TextType: Text, Text: tok.value})
			i++
		}
	}
	return nodes, len(tokens)
}
