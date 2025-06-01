package tokenizer

import (
	"strings"
)

const punct = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"

type Token struct {
	Kind  string // delimiter type or "text"
	Value string // raw text
	// File  string
	// Line  int
	// Col   int
}

func TokenizeInline(input string) []Token {
	var out []Token
	runes := []rune(input)
	n := len(runes)
	for i := 0; i < n; {
		if runes[i] == '\\' {
			if i+1 < n && strings.ContainsRune(punct, runes[i+1]) {
				out = append(out, Token{Kind: "text", Value: string(runes[i+1])})
				i += 2
				continue
			}
		}

		r := runes[i] // set and check rune, create Token depending on what we found
		if r == '`' { // inline code
			j := i + 1 // variable for following rune
			for j < n && runes[j] != '`' {
				if runes[j] != '\\' {
					j++
				} else {
					j += 2
				}
			}
			if j < n {
				out = append(out, Token{Kind: "code", Value: string(runes[i+1 : j])}) // create code Token and append
				i = j + 1                                                             // advance to next rune
			} else {
				out = append(out, Token{Kind: "text", Value: "`"}) // default to plaintext
				i++                                                // advance one rune
			}
			continue
		}

		if i+1 < n && runes[i] == '=' && runes[i+1] == '=' { //highlight
			out = append(out, Token{Kind: "==", Value: "=="})
			i += 2
			continue
		}
		if r == '!' && i+1 < n && runes[i+1] == '[' { // image open
			out = append(out, Token{Kind: "![", Value: "![]"}) // create image opening Token and append
			i += 2                                             // advance two runes
			continue
		}

		if r == '[' || r == ']' || r == '(' || r == ')' { // link/list delimiter
			out = append(out, Token{Kind: string(r), Value: string(r)})
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
				out = append(out, Token{Kind: m, Value: m})
				runLen -= 3
			}
			if runLen >= 2 && r != '^' { // double for bold/strikethrough/subscript
				m := strings.Repeat(string(r), 2)
				out = append(out, Token{Kind: m, Value: m})
				runLen -= 2
			}
			if runLen == 1 { // single for italic/subscript
				m := string(r)
				out = append(out, Token{Kind: m, Value: m})
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
			out = append(out, Token{Kind: "text", Value: string(runes[i:j])}) // create text Token and append
		}
		i = j // advance original position
	}
	return out // return Tokens slice
}
