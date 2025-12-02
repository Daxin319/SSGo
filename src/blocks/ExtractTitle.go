package blocks

import (
	"errors"
	"strings"
)

func ExtractTitle(s string) (string, string, error) {
	blocks := MarkdownToBlocks(s)
	var fixed []string
	var title string
	count := 0

	for _, blck := range blocks {
		bType := BlockToBlockType(blck)
		if bType != Heading {
			fixed = append(fixed, blck)
		} else {
			n, _ := HeaderNum(blck)
			if n != 1 {
				fixed = append(fixed, blck)
			} else {
				title = strings.TrimLeft(strings.TrimRight(blck, " "), "# ")
				count++
			}
		}
	}
	if count == 0 {
		return "", "", errors.New("no h1 header found")
	}
	joined := strings.Join(fixed, "\n\n")
	return title, joined, nil
}
