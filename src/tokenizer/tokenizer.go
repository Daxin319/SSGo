package tokenizer

import (
	"strings"
)

const punct = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"

type Token struct {
	Kind  string // delimiter type or "text"
	Value string // raw text
	// file  string
	// line  int    -> uncomment when ready to add this to diagnostics
	// col   int
}

func TokenizeInline(input string) []Token {
	var out []Token
	runes := []rune(input)
	n := len(runes)
	for i := 0; i < n; {

		r := runes[i] // set and check rune, create Token depending on what we found
		if r == '`' {
			// count opening ticks
			j := i
			for j < n && runes[j] == '`' {
				j++
			}
			delimLen := j - i

			// scan for a matching run of exactly delimLen backticks
			k := j
			for {
				// find next backtick
				for k < n && runes[k] != '`' {
					k++
				}
				if k >= n {
					break
				}
				// count ticks at k
				l := k
				for l < n && runes[l] == '`' {
					l++
				}
				if l-k == delimLen {
					// **skip** if this run is escaped (preceded by a backslash)
					if k > 0 && runes[k-1] == '\\' {
						k = l
						continue
					}
					// found closer—capture literal content
					content := string(runes[j:k])
					// trim a single leading/trailing space per spec
					if len(content) > 1 && content[0] == ' ' && content[len(content)-1] == ' ' {
						content = content[1 : len(content)-1]
					}
					out = append(out, Token{Kind: "code", Value: content})
					i = l
					goto nextToken
				}
				k = l
			}

			// no closer: emit each backtick literally
			for range delimLen {
				out = append(out, Token{Kind: "text", Value: "`"})
			}
			i = j
			continue

		nextToken:
			continue
		}

		if r == '\\' {
			if i+1 < n && strings.ContainsRune(punct, runes[i+1]) {
				out = append(out, Token{Kind: "text", Value: string(runes[i+1])})
				i += 2
				continue
			}
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
