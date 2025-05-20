package nodes

import (
	"strings"
)

const punct = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"

type token struct {
	kind  string // delimiter type or "text"
	value string // raw text
	file  string
	line  int
	col   int
}

func tokenizeInline(input string) []token {
	var out []token
	runes := []rune(input)
	n := len(runes)
	for i := 0; i < n; {
		if runes[i] == '\\' { // check for escape character before anything else
			if i+1 < n && strings.ContainsRune(punct, runes[i+1]) {
				out = append(out, token{kind: "text", value: string(runes[i+1])}) // emit escapable rune literal
				i += 2                                                            // advance past the escaped character
			} else if i+1 == n {
				out = append(out, token{kind: "text", value: "<br />"})
				i += 6
			} else {
				out = append(out, token{kind: "text", value: "\\"}) // emit \ literal
				i++
			}
			continue
		}

		r := runes[i] // set and check rune, create token depending on what we found
		if r == '`' { // inline code
			j := i + 1 // variable for following rune
			for j < n && runes[j] != '`' {
				j++
			}
			if j < n {
				out = append(out, token{kind: "code", value: string(runes[i+1 : j])}) // create code token and append
				i = j + 1                                                             // advance to next rune
			} else {
				out = append(out, token{kind: "text", value: "`"}) // default to plaintext
				i++                                                // advance one rune
			}
			continue
		}

		if i+1 < n && runes[i] == '=' && runes[i+1] == '=' { //highlight
			out = append(out, token{kind: "==", value: "=="})
			i += 2
			continue
		}
		if r == '!' && i+1 < n && runes[i+1] == '[' { // image open
			out = append(out, token{kind: "![", value: "![]"}) // create image opening token and append
			i += 2                                             // advance two runes
			continue
		}

		if r == '[' || r == ']' || r == '(' || r == ')' { // link/list delimiter
			out = append(out, token{kind: string(r), value: string(r)})
			i++ // advance one rune
			continue
		}

		if r == '*' || r == '_' || r == '~' || r == '^' { // all other delimiters, some are single and some can repeat
			j := i
			for j < n && runes[j] == r {
				j++
			}
			runLen := j - i                            // run length is current pos minus start of run
			if (r == '*' || r == '_') && runLen >= 3 { // triple for bolditalic
				m := strings.Repeat(string(r), 3)
				out = append(out, token{kind: m, value: m})
				runLen -= 3
			}
			if runLen >= 2 && r != '^' { // double for bold/strikethrough/subscript
				m := strings.Repeat(string(r), 2)
				out = append(out, token{kind: m, value: m})
				runLen -= 2
			}
			if runLen == 1 { // single for italic/subscript
				m := string(r)
				out = append(out, token{kind: m, value: m})
			}
			i = j // reset position to end of run
			continue
		}

		// plaintext run
		j := i
		for j < n {
			c := runes[j] // current rune
			if c == '\\' || c == '`' || c == '!' || strings.ContainsAny(string(c), "[]()*_~^") || (c == '=' && j+1 < n && runes[j+1] == '=') {
				break // delimiter
			}
			j++
		}
		if j > i {
			out = append(out, token{kind: "text", value: string(runes[i:j])}) // create text token and append
		}
		i = j // advance original position
	}
	return out // return tokens slice
}
