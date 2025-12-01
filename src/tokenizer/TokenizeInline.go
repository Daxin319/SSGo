package tokenizer

import (
	"fmt"
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

var EmailRE = regexp.MustCompile(`^(?:[A-Za-z0-9._%+\-]+:)?[A-Za-z0-9._%+\-]+@[A-Za-z0-9.\-]+\.[A-Za-z]{2,}$`)

// ProtocolRE Protocol pattern for URLs (http://, https://, etc.)
var ProtocolRE = regexp.MustCompile(`^(?:https?|ftp|ftps|sftp|ws|wss)://[^\s]+$`)

// GfmDomainRE Improved GFM-style domain regex: matches only valid domains/URLs, no spaces or attributes, TLD 2-6 letters
var GfmDomainRE = regexp.MustCompile(`^(?:[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?\.)+[a-zA-Z]{2,6}(?::[0-9]+)?(?:/[a-zA-Z0-9\-._~:/?#\[\]@!$&'()*+,;=%^]*)?$`)

var rawHTMLTagRE = regexp.MustCompile(`^<(?:!--[\s\S]*?--|/?[a-zA-Z][a-zA-Z0-9-]*(?:\s[^>]*)?)>`)

func TokenizeInline(input string) []Token {
	fmt.Printf("Starting tokenization of input length %d\n", len(input))
	var out []Token
	runes := []rune(input)
	n := len(runes)
	for i := 0; i < n; {
		r := runes[i]
		fmt.Printf("Processing byte %d: %q\n", i, r)

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
			remaining := string(runes[i:])
			if match := rawHTMLTagRE.FindString(remaining); match != "" {
				out = append(out, Token{Kind: "raw_html", Value: match})
				i += len([]rune(match))
				continue
			}
			j := i + 1
			for j < n && runes[j] != '>' {
				j++
			}
			if j < n {
				inner := string(runes[i+1 : j])
				if ProtocolRE.MatchString(inner) {
					out = append(out, Token{Kind: "<", Value: inner})
					i = j + 1
					continue
				}
				if EmailRE.MatchString(inner) {
					out = append(out, Token{Kind: "<", Value: inner})
					i = j + 1
					continue
				}
				if GfmDomainRE.MatchString(inner) {
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
				}
				// If we're here, it's not a valid autolink. Treat as raw text.
				out = append(out, Token{Kind: "raw_text", Value: "&lt;" + inner + "&gt;"})
				i = j + 1
				continue
			}
		}

		if r == '`' {
			fmt.Printf("Found backtick at position %d\n", i)
			j := i // count opening ticks
			for j < n && runes[j] == '`' {
				j++
			}
			delimLen := j - i
			fmt.Printf("Found %d backticks at start\n", delimLen)

			k := j // scan for a matching run of exactly delimLen backticks
			found := false
			for k < n {
				// Find next backtick
				for k < n && runes[k] != '`' {
					k++
				}
				if k >= n {
					break
				}

				// Count matching backticks
				l := k
				for l < n && runes[l] == '`' {
					l++
				}

				fmt.Printf("Found potential closing backticks at %d, length %d\n", k, l-k)

				if l-k == delimLen {
					// Check if this is an escaped backtick sequence
					if k > 0 && runes[k-1] == '\\' {
						fmt.Printf("Found escaped backticks at %d, skipping\n", k)
						// Skip this sequence and continue looking
						k = l
						continue
					}

					content := string(runes[j:k])
					if len(content) > 1 && content[0] == ' ' && content[len(content)-1] == ' ' {
						content = content[1 : len(content)-1]
					}
					fmt.Printf("Emitting code token with content: %q\n", content)
					out = append(out, Token{Kind: "code", Value: content})
					i = l
					found = true
					break
				}
				k = l
			}

			if !found {
				fmt.Printf("No matching backticks found, emitting literal backticks\n")
				// No matching backticks found, emit literal backticks
				for range delimLen {
					out = append(out, Token{Kind: "text", Value: "`"})
				}
				i = j
			}
			continue
		}

		if r == '\\' {
			if i+1 < n && strings.ContainsRune(punct, runes[i+1]) {
				out = append(out, Token{Kind: "text", Value: string(runes[i+1])})
				i += 2
				continue
			}
			// If not a valid escape, treat as literal backslash
			out = append(out, Token{Kind: "text", Value: "\\"})
			i++
			continue
		}

		if i+1 < n && runes[i] == '=' && runes[i+1] == '=' { //highlight
			out = append(out, Token{Kind: "==", Value: "=="})
			i += 2
			continue
		}
		if r == '!' && i+1 < n && runes[i+1] == '[' { // image open
			fmt.Printf("Found image marker at %d, next char is %q\n", i, runes[i+1])
			out = append(out, Token{Kind: "![", Value: "!["}) // create image opening Token and append
			i += 2                                            // advance two runes
			fmt.Printf("Advanced to position %d, next char is %q\n", i, runes[i])
			continue
		}

		if r == '[' || r == ']' || r == '(' || r == ')' { // link/list delimiter
			fmt.Printf("Found delimiter %q at %d\n", r, i)
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
			if c == '\\' || c == '`' || c == '!' || c == '<' || strings.ContainsAny(string(c), "[]()*_~^") || (c == '=' && j+1 < n && runes[j+1] == '=') {
				break // delimiter
			}
			j++
		}
		if j > i {
			fmt.Printf("Emitting text token from %d to %d: %q\n", i, j, string(runes[i:j]))
			out = append(out, Token{Kind: "text", Value: string(runes[i:j])}) // create text Token and append
			i = j                                                             // advance original position
		} else {
			// If we didn't find any text to emit, we need to handle the current character
			fmt.Printf("No text to emit at %d, handling current char %q\n", i, r)
			out = append(out, Token{Kind: "text", Value: string(r)})
			i++
		}
	}
	fmt.Printf("Finished tokenization, emitted %d tokens\n", len(out))
	return out // return Tokens slice
}
