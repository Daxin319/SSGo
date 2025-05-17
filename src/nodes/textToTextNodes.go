// textToTextNodes.go
package nodes

import (
	"strings"
)

// TextToTextNodes converts a markdown string into inline TextNodes.
func TextToTextNodes(s string) []TextNode {
	tokens := tokenizeInline(s)
	nodes, _ := parseTokens(tokens, 0, "")
	return nodes
}

type token struct {
	kind  string // delimiter or "text"
	value string
}

func tokenizeInline(input string) []token {
	var tokens []token
	length := len(input)
	for i := 0; i < length; {
		ch := input[i]
		// code delimiter
		if ch == '`' {
			tokens = append(tokens, token{kind: "`", value: "`"})
			i++
			continue
		}
		// image start
		if i+1 < length && input[i] == '!' && input[i+1] == '[' {
			tokens = append(tokens, token{kind: "![", value: "!["})
			i += 2
			continue
		}
		// link start or bracket
		if ch == '[' || ch == ']' || ch == '(' || ch == ')' {
			tokens = append(tokens, token{kind: string(ch), value: string(ch)})
			i++
			continue
		}
		// emphasis / strikethrough / subscript delimiter
		if ch == '*' || ch == '_' || ch == '~' {
			runChar := ch
			j := i
			for j < length && input[j] == runChar {
				j++
			}
			runLen := j - i
			// only '*' and '_' get triple-delimiter
			if (runChar == '*' || runChar == '_') && runLen >= 3 {
				tokens = append(tokens, token{
					kind:  strings.Repeat(string(runChar), 3),
					value: strings.Repeat(string(runChar), 3),
				})
				runLen -= 3
			}
			// doubles
			for runLen >= 2 {
				tokens = append(tokens, token{
					kind:  strings.Repeat(string(runChar), 2),
					value: strings.Repeat(string(runChar), 2),
				})
				runLen -= 2
			}
			// singles
			if runLen == 1 {
				tokens = append(tokens, token{kind: string(runChar), value: string(runChar)})
			}
			i = j
			continue
		}
		// plain text
		j := i
		for j < length {
			c := input[j]
			if c == '`' || c == '!' || c == '[' || c == ']' ||
				c == '(' || c == ')' || c == '*' || c == '_' || c == '~' {
				break
			}
			j++
		}
		tokens = append(tokens, token{kind: "text", value: input[i:j]})
		i = j
	}
	return tokens
}

func parseTokens(tokens []token, pos int, delim string) ([]TextNode, int) {
	var nodes []TextNode
	for i := pos; i < len(tokens); {
		tok := tokens[i]
		if delim != "" && tok.kind == delim {
			return nodes, i
		}
		switch tok.kind {
		case strings.Repeat("~", 2):
			children, j := parseTokens(tokens, i+1, tok.kind)
			nodes = append(nodes, TextNode{TextType: Strikethrough, Children: children})
			i = j + 1

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

		case "~":
			children, j := parseTokens(tokens, i+1, tok.kind)
			nodes = append(nodes, TextNode{TextType: Subscript, Children: children})
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
				t := Link
				if isImage {
					t = Image
				}
				nodes = append(nodes, TextNode{TextType: t, Url: url, Children: children})
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
