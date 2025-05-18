package nodes

import "strings"

type token struct {
	kind  string // delimiter type or "text"
	value string // raw text
}

func tokenizeInline(input string) []token {
	var out []token
	n := len(input)
	for i := 0; i < n; {
		r := input[i] // set and check rune, create token depending on what we found
		if r == '`' { // inline code
			j := i + 1                     // variable for following rune
			for j < n && input[j] != '`' { // as long as it's not the end of the code literal
				j++ // advance until closing backtick or end of string
			}
			if j < n {
				out = append(out, token{kind: "code", value: input[i+1 : j]}) // create code token and append
				i = j + 1                                                     // advance to next rune
			} else {
				out = append(out, token{kind: "text", value: "`"}) // default to plaintext
				i++                                                // advance one rune
			}
			continue
		}
		if r == '!' && i+1 < n && input[i+1] == '[' { // if rune is ! and next rune is [
			out = append(out, token{kind: "![", value: "![]"}) // create image opening token and append
			i += 2                                             // advance 2 runes
			continue
		}
		if r == '[' || r == ']' || r == '(' || r == ')' { // if rune is square bracket or paren
			out = append(out, token{kind: string(r), value: string(r)})
			i++ // advance one rune
			continue
		}
		if r == '*' || r == '_' || r == '~' || r == '^' { // all other delimiters, some are single and some can repeat
			j := i // we have to iterate and figure out how long the delimiter is since these could repeat
			for j < n && input[j] == r {
				j++ // as long as we find the same rune keep going
			}
			runLen := j - i                            // run length is current pos - start of run pos
			if (r == '*' || r == '_') && runLen >= 3 { // if it's bold or italic and there are 3 or more of them
				m := strings.Repeat(string(r), 3)           // multi-rune is 3 runes
				out = append(out, token{kind: m, value: m}) // create token and append
				runLen -= 3                                 // subtract 3 from run length to reset for next iteration
			}
			if runLen >= 2 && r != '^' { // repeat for 2
				m := strings.Repeat(string(r), 2)
				out = append(out, token{kind: m, value: m}) // create token and append
				runLen -= 2                                 // subtract 2 from run length to reset for next iteration
			}
			if runLen == 1 {
				m := string(r)
				out = append(out, token{kind: m, value: m})
			}
			i = j // reset original position to current
			continue
		}
		j := i      // set current position to first non-delimiter rune
		for j < n { // as long as it's not the end of the string
			c := input[j]                                                           // char is current char
			if c == '`' || c == '!' || strings.ContainsAny(string(c), "[]()*_~^") { // if it's a delimiter, break out of the for loop
				break
			}
			j++ // advance 1 rune if not a delimiter
		}
		if j > i { // if current pos > original pos create plaintext token and append
			out = append(out, token{kind: "text", value: input[i:j]})
		}
		i = j // set original pos to current pos
	}
	return out // return tokens slice
}
