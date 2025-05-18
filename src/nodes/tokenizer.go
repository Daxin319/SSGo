package nodes

import "strings"

type delimRun struct {
	marker string
	pos    int
}
type token struct {
	kind  string // delimiter type or "text"
	value string // raw text
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
