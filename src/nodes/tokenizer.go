package nodes

import "strings"

type token struct {
	kind  string // delimiter type or "text"
	value string // raw text
}

func tokenizeInline(input string) []token {
	var tokens []token
	length := len(input)
	for i := 0; i < length; {
		ch := input[i]
		if ch == '`' {
			tokens = append(tokens, token{kind: "`", value: "`"})
			i++
			continue
		}
		if i+1 < length && input[i] == '!' && input[i+1] == '[' {
			tokens = append(tokens, token{kind: "![", value: "!["})
			i += 2
			continue
		}
		if ch == '[' {
			tokens = append(tokens, token{kind: "[", value: "["})
			i++
			continue
		}
		if ch == ']' {
			tokens = append(tokens, token{kind: "]", value: "]"})
			i++
			continue
		}
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
		if ch == '*' || ch == '_' {
			runChar := ch
			j := i
			for j < length && input[j] == runChar {
				j++
			}
			runLen := j - i
			if runLen >= 3 {
				tokens = append(tokens, token{kind: strings.Repeat(string(runChar), 3), value: strings.Repeat(string(runChar), 3)})
				runLen -= 3
			}
			for runLen >= 2 {
				tokens = append(tokens, token{kind: strings.Repeat(string(runChar), 2), value: strings.Repeat(string(runChar), 2)})
				runLen -= 2
			}
			if runLen == 1 {
				tokens = append(tokens, token{kind: string(runChar), value: string(runChar)})
			}
			i = j
			continue
		}
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
