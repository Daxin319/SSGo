package tokenizer

import (
	"html"
	"strconv"
)

// parseEntity attempts to parse an HTML entity or character reference starting at position i
// Returns the parsed entity and the number of runes consumed, or empty string and 0 if not an entity
func parseEntity(runes []rune, i int) (string, int) {
	if i >= len(runes) || runes[i] != '&' {
		return "", 0
	}

	// Find the end of the entity (semicolon or end of input)
	end := i + 1
	for end < len(runes) && runes[end] != ';' && runes[end] != '&' {
		end++
	}
	if end >= len(runes) || runes[end] != ';' {
		return "", 0
	}

	// Try to unescape the entity
	entity := string(runes[i : end+1])
	unescaped := html.UnescapeString(entity)
	if unescaped != entity {
		return unescaped, end - i + 1
	}

	// If html.UnescapeString didn't handle it, try numeric character references
	if i+2 < len(runes) && runes[i+1] == '#' {
		base := 10
		start := i + 2
		if start < len(runes) && runes[start] == 'x' {
			base = 16
			start++
		}

		// Find the end of the number
		numEnd := start
		for numEnd < len(runes) && ((base == 10 && runes[numEnd] >= '0' && runes[numEnd] <= '9') ||
			(base == 16 && ((runes[numEnd] >= '0' && runes[numEnd] <= '9') ||
				(runes[numEnd] >= 'a' && runes[numEnd] <= 'f') ||
				(runes[numEnd] >= 'A' && runes[numEnd] <= 'F')))) {
			numEnd++
		}

		if numEnd > start && numEnd < len(runes) && runes[numEnd] == ';' {
			numStr := string(runes[start:numEnd])
			var num int64
			if base == 10 {
				num, _ = strconv.ParseInt(numStr, 10, 32)
			} else {
				num, _ = strconv.ParseInt(numStr, 16, 32)
			}
			if num > 0 && num <= 0x10FFFF {
				return string(rune(num)), numEnd - i + 1
			}
		}
	}

	return "", 0
}
