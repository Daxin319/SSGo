package blocks

import "strings"

// clean all whitespace and newline characters from paragraphs. if string is empty after stripping move past it

func CleanNewlines(s string) string {
	split := strings.Split(s, "\n")
	var replaced []string

	for _, line := range split {
		if len(line) == 0 || line == " " || line == "\n" {
			continue
		}
		replaced = append(replaced, line)
	}
	joined := strings.Join(replaced, " ")
	return joined
}
