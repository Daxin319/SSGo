package nodes

import (
	"strings"
)

type token struct {
	kind  string
	value string
}

type delimRun struct {
	marker string
	pos    int
}

func TextToTextNodes(s string) []TextNode {
	toks := tokenizeInline(s)
	return parseInlineStack(toks)
}

func tokenizeInline(input string) []token {
	var out []token
	n := len(input)
	for i := 0; i < n; {
		r := input[i]
		if r == '`' {
			j := i + 1
			for j < n && input[j] != '`' {
				j++
			}
			if j < n {
				out = append(out, token{kind: "code", value: input[i+1 : j]})
				i = j + 1
			} else {
				out = append(out, token{kind: "text", value: "`"})
				i++
			}
			continue
		}
		if r == '!' && i+1 < n && input[i+1] == '[' {
			out = append(out, token{kind: "![", value: "!["})
			i += 2
			continue
		}
		if r == '[' || r == ']' || r == '(' || r == ')' {
			out = append(out, token{kind: string(r), value: string(r)})
			i++
			continue
		}
		if r == '*' || r == '_' || r == '~' {
			j := i
			for j < n && input[j] == r {
				j++
			}
			runLen := j - i
			if (r == '*' || r == '_') && runLen >= 3 {
				m := strings.Repeat(string(r), 3)
				out = append(out, token{kind: m, value: m})
				runLen -= 3
			}
			if runLen >= 2 {
				m := strings.Repeat(string(r), 2)
				out = append(out, token{kind: m, value: m})
				runLen -= 2
			}
			if runLen == 1 {
				m := string(r)
				out = append(out, token{kind: m, value: m})
			}
			i = j
			continue
		}
		j := i
		for j < n {
			c := input[j]
			if c == '`' || c == '!' || strings.ContainsAny(string(c), "[]()*_~") {
				break
			}
			j++
		}
		if j > i {
			out = append(out, token{kind: "text", value: input[i:j]})
		}
		i = j
	}
	return out
}

func parseInlineStack(tokens []token) []TextNode {
	var nodes []TextNode
	var stack []delimRun

	wrap := func(marker string, children []TextNode) TextNode {
		switch marker {
		case "*", "_":
			return TextNode{TextType: Italic, Children: children}
		case "**", "__":
			return TextNode{TextType: Bold, Children: children}
		case "***", "___":
			return TextNode{TextType: BoldItalic, Children: children}
		case "~~":
			return TextNode{TextType: Strikethrough, Children: children}
		case "~":
			return TextNode{TextType: Subscript, Children: children}
		}
		return TextNode{TextType: Text, Text: marker}
	}

	processAsterisk := func(m string) {
		length := len(m)
		char := m[0]
		if length <= 2 {
			for i := len(stack) - 1; i >= 0; i-- {
				if stack[i].marker == m {
					op := stack[i].pos
					content := append([]TextNode{}, nodes[op:]...)
					nodes = nodes[:op]
					nodes = append(nodes, wrap(m, content))
					stack = append(stack[:i], stack[i+1:]...)
					return
				}
			}
			stack = append(stack, delimRun{marker: m, pos: len(nodes)})
			return
		}
		remaining := length
		for remaining > 0 {
			idx := -1
			for j := len(stack) - 1; j >= 0; j-- {
				if stack[j].marker[0] == char && len(stack[j].marker) <= remaining {
					idx = j
					break
				}
			}
			if idx >= 0 {
				mrk := stack[idx].marker
				op := stack[idx].pos
				content := append([]TextNode{}, nodes[op:]...)
				nodes = nodes[:op]
				nodes = append(nodes, wrap(mrk, content))
				stack = append(stack[:idx], stack[idx+1:]...)
				remaining -= len(mrk)
			} else {
				marker := strings.Repeat(string(char), remaining)
				stack = append(stack, delimRun{marker: marker, pos: len(nodes)})
				break
			}
		}
	}

	for i := 0; i < len(tokens); {
		t := tokens[i]
		switch t.kind {
		case "code":
			nodes = append(nodes, TextNode{TextType: Code, Text: t.value})
			i++
		case "![[", "[":
			isImage := t.kind == "!["
			j := i + 1
			for j < len(tokens) && tokens[j].kind != "]" {
				j++
			}
			altNodes := parseInlineStack(tokens[i+1 : j])
			if j+3 < len(tokens) && tokens[j+1].kind == "(" && tokens[j+3].kind == ")" {
				url := tokens[j+2].value
				typeEnum := Link
				if isImage {
					typeEnum = Image
				}
				nodes = append(nodes, TextNode{TextType: typeEnum, Url: url, Children: altNodes})
				i = j + 4
			} else {
				nodes = append(nodes, TextNode{TextType: Text, Text: t.value})
				i++
			}
		case "*", "**", "***", "_", "__", "___":
			processAsterisk(t.kind)
			i++
		case "~~", "~":
			m := t.kind
			closed := false
			for j := len(stack) - 1; j >= 0; j-- {
				if stack[j].marker == m {
					op := stack[j].pos
					content := append([]TextNode{}, nodes[op:]...)
					nodes = nodes[:op]
					nodes = append(nodes, wrap(m, content))
					stack = append(stack[:j], stack[j+1:]...)
					closed = true
					break
				}
			}
			if !closed {
				stack = append(stack, delimRun{marker: m, pos: len(nodes)})
			}
			i++
		case "]", "(", ")":
			nodes = append(nodes, TextNode{TextType: Text, Text: t.value})
			i++
		default:
			nodes = append(nodes, TextNode{TextType: Text, Text: t.value})
			i++
		}
	}

	for _, op := range stack {
		nodes = append(nodes, TextNode{TextType: Text, Text: op.marker})
	}

	return nodes
}
