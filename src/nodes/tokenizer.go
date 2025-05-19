package nodes

import (
	"strings"
)

const punct = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"

type token struct {
	kind  string // delimiter type or "text"
	value string // raw text
}

func tokenizeInline(input string) []token {
	var out []token
	runes := []rune(input)
	n := len(runes)
	for i := 0; i < n; {
		if runes[i] == '\\' { // check for escape character before anything else
			if i+1 < n && strings.ContainsRune(punct, runes[i+1]) {
				out = append(out, token{kind: "text", value: string(runes[i+1])}) // if next rune is escapable then emit escapable rune literal
				i += 2                                                            // advance past the escaped character and allow the \ to fade into the ether
			} else {
				out = append(out, token{kind: "text", value: "\\"}) // non-escapable character, emit \ literal
				i++
			}
			continue
		}

		r := runes[i] // set and check rune, create token depending on what we found
		if r == '`' { // inline code
			j := i + 1                     // variable for following rune
			for j < n && runes[j] != '`' { // as long as it's not the end of the code literal
				j++ // advance until closing backtick or end of string
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

		if r == '!' && i+1 < n && runes[i+1] == '[' { // if rune is '!' and next is '['
			out = append(out, token{kind: "![", value: "![]"}) // create image opening token and append
			i += 2                                             // advance two runes
			continue
		}

		if r == '[' || r == ']' || r == '(' || r == ')' { // if rune is square bracket or paren
			out = append(out, token{kind: string(r), value: string(r)})
			i++ // advance one rune
			continue
		}

		if r == '*' || r == '_' || r == '~' || r == '^' { // all other delimiters, some are single and some can repeat
			j := i
			for j < n && runes[j] == r {
				j++ // as long as we find the same rune keep going
			}
			runLen := j - i                            // run length is current pos minus start of run
			if (r == '*' || r == '_') && runLen >= 3 { // triple for bolditalic
				m := strings.Repeat(string(r), 3)
				out = append(out, token{kind: m, value: m}) // create token and append
				runLen -= 3                                 // subtract three from run length
			}
			if runLen >= 2 && r != '^' { // double for bold/strikethrough/subscript
				m := strings.Repeat(string(r), 2)
				out = append(out, token{kind: m, value: m}) // create token and append
				runLen -= 2                                 // subtract two from run length
			}
			if runLen == 1 { // single for italic/subscript
				m := string(r)
				out = append(out, token{kind: m, value: m})
			}
			i = j // reset position to end of run
			continue
		}

		j := i      // set current position to first non-delimiter rune
		for j < n { // as long as it's not the end of the string
			c := runes[j]                                                                        // current rune
			if c == '\\' || c == '`' || c == '!' || strings.ContainsAny(string(c), "[]()*_~^") { // delimiter?
				break // yes, stop
			}
			j++ // advance one rune
		}
		if j > i { // plaintext run
			out = append(out, token{kind: "text", value: string(runes[i:j])}) // create text token and append
		}
		i = j // advance original position
	}
	return out // return tokens slice
}
