package tokenizer

import (
	"regexp"
	"strings"

	"golang.org/x/net/publicsuffix"
)

const punct = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"

type Token struct {
	Kind  string // delimiter type or "text"
	Value string // raw text
	// File  string
	// Line  int
	// Col   int
}

var emailRE = regexp.MustCompile(`^(?:[A-Za-z0-9._%+\-]+:)?[A-Za-z0-9._%+\-]+@[A-Za-z0-9.\-]+\.[A-Za-z]{2,}$`)

// Protocol pattern for URLs (http://, https://, etc.)
var protocolRE = regexp.MustCompile(`^(?:https?|ftp|ftps|sftp|ws|wss)://[^\s]+$`)

// Improved GFM-style domain regex: matches only valid domains/URLs, no spaces or attributes, TLD 2-6 letters
var gfmDomainRE = regexp.MustCompile(`^(?:[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?\.)+[a-zA-Z]{2,6}(?::[0-9]+)?(?:/[a-zA-Z0-9\-._~:/?#\[\]@!$&'()*+,;=%]*)?$`)

func TokenizeInline(input string) []Token {
	var out []Token
	runes := []rune(input)
	n := len(runes)
	for i := 0; i < n; {
		r := runes[i]

		// Check for HTML entities and character references
		if r == '&' {
			if entity, consumed := parseEntity(runes, i); consumed > 0 {
				out = append(out, Token{Kind: "text", Value: entity})
				i += consumed
				continue
			}
		}

		// If the next characters form a valid CommonMark "raw HTML" chunk,
		// consume it all as one token of kind="raw_html".
		if r == '<' {
			j := i + 1
			for j < n && runes[j] != '>' {
				j++
			}
			if j < n {
				inner := string(runes[i+1 : j])
				if protocolRE.MatchString(inner) {
					out = append(out, Token{Kind: "<", Value: inner})
					i = j + 1
					continue
				}
				if emailRE.MatchString(inner) {
					out = append(out, Token{Kind: "<", Value: inner})
					i = j + 1
					continue
				}
				if gfmDomainRE.MatchString(inner) {
					// Extract the domain part (before any /, ?, #, or :)
					domain := inner
					for i, c := range domain {
						if c == '/' || c == '?' || c == '#' || c == ':' {
							domain = domain[:i]
							break
						}
					}
					// Remove user:pass@ if present
					if at := strings.LastIndex(domain, "@"); at != -1 {
						domain = domain[at+1:]
					}
					// Use publicsuffix to check if the TLD is real
					if _, icann := publicsuffix.PublicSuffix(domain); icann {
						out = append(out, Token{Kind: "<", Value: inner})
						i = j + 1
						continue
					}
					// Not a valid autolink, treat as text (escape angle brackets)
					out = append(out, Token{Kind: "text", Value: "&lt;" + inner + "&gt;"})
					i = j + 1
					continue
				}
			}
		}

		if r == '`' {
			j := i // count opening ticks
			for j < n && runes[j] == '`' {
				j++
			}
			delimLen := j - i

			k := j // scan for a matching run of exactly delimLen backticks
			for {
				for k < n && runes[k] != '`' { // find next backtick
					k++
				}
				if k >= n {
					break
				}
				l := k // count ticks at k
				for l < n && runes[l] == '`' {
					l++
				}
				if l-k == delimLen {
					if k > 0 && runes[k-1] == '\\' { // **skip** if this run is escaped (preceded by a backslash)
						k = l
						continue
					}
					content := string(runes[j:k])                                                // found closer—capture literal content
					if len(content) > 1 && content[0] == ' ' && content[len(content)-1] == ' ' { // trim a single leading/trailing space per spec
						content = content[1 : len(content)-1]
					}
					out = append(out, Token{Kind: "code", Value: content})
					i = l
					goto nextToken
				}
				k = l
			}

			for range delimLen { // no closer: emit each backtick literally
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
