package nodes

import (
	"strings"
)

// TextToTextNodes converts a markdown string into a slice of TextNode representing inline structure.
func TextToTextNodes(s string) []TextNode {
	tokens := tokenizeInline(s)
	nodes, _ := parseTokens(tokens, 0, "")
	return nodes
}

// token represents an inline delimiter or text span.
type token struct {
	kind  string // delimiter type or "text"
	value string // raw text
}

// tokenizeInline splits the input into tokens for inline parsing.
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
		// link start
		if ch == '[' {
			tokens = append(tokens, token{kind: "[", value: "["})
			i++
			continue
		}
		// closing bracket
		if ch == ']' {
			tokens = append(tokens, token{kind: "]", value: "]"})
			i++
			continue
		}
		// parentheses
		if ch == '(' {
			tokens = append(tokens, token{kind: "(", value: "("})
			i++
			continue
		}
		if ch == ')' {
			tokens = append(tokens, token{kind: ")", value: ")"})
			i++
			continue
		}
		// emphasis delimiter '*' or '_'
		if ch == '*' || ch == '_' {
			runChar := ch
			j := i
			for j < length && input[j] == runChar {
				j++
			}
			runLen := j - i
			// triple delimiter
			if runLen >= 3 {
				tokens = append(tokens, token{kind: strings.Repeat(string(runChar), 3), value: strings.Repeat(string(runChar), 3)})
				runLen -= 3
			}
			// double delimiters
			for runLen >= 2 {
				tokens = append(tokens, token{kind: strings.Repeat(string(runChar), 2), value: strings.Repeat(string(runChar), 2)})
				runLen -= 2
			}
			// single delimiter
			if runLen == 1 {
				tokens = append(tokens, token{kind: string(runChar), value: string(runChar)})
			}
			i = j
			continue
		}
		// text span
		j := i
		for j < length {
			c := input[j]
			if c == '`' || c == '!' || c == '[' || c == ']' || c == '(' || c == ')' || c == '*' || c == '_' {
				break
			}
			j++
		}
		tokens = append(tokens, token{kind: "text", value: input[i:j]})
		i = j
	}
	return tokens
}

// parseTokens parses tokens starting at pos until a matching closing delimiter (if any).
// delim is the delimiter kind to close (e.g., "*", "**", "***", "_", "__", "___", "]", "").
func parseTokens(tokens []token, pos int, delim string) ([]TextNode, int) {
	var nodes []TextNode
	for i := pos; i < len(tokens); {
		tok := tokens[i]
		// closing delimiter?
		if delim != "" && tok.kind == delim {
			return nodes, i
		}
		switch tok.kind {
		case strings.Repeat("*", 3), strings.Repeat("_", 3):
			// bold + italic
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
