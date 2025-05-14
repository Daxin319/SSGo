package main

import (
	"errors"
	"main/src/blocks"
	"strings"
)

func ExtractTitle(s string) (string, string, error) {
	blcks := blocks.MarkdownToBlocks(s)
	var fixed []string
	var title string
	count := 0

	for _, blck := range blcks {
		bType := blocks.BlockToBlockType(blck)
		if bType != blocks.Heading {
			fixed = append(fixed, blck)
		} else {
			n, _ := blocks.HeaderNum(blck)
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
