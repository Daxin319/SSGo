package blocks

import (
	"regexp"
	"strings"
)

type BlockType int

const (
	Paragraph BlockType = iota
	Heading
	CodeBlock
	Quote
	UnorderedList
	OrderedList
	ThematicBreak
)

func BlockToBlockType(block string) BlockType {
	trimmed := strings.TrimLeft(strings.TrimRight(block, " \r\n"), " \r\n")

	if isThematicBreak(trimmed) {
		return ThematicBreak
	}

	if len(trimmed) >= 1 && trimmed[0] == '#' {
		n, i := HeaderNum(trimmed)
		if string(trimmed[n-1]) != "#" || string(trimmed[i]) != " " {
			return Paragraph
		}
		return Heading
	}
	if len(trimmed) >= 6 {
		if trimmed[:3] == "```" && trimmed[len(trimmed)-3:] == "```" {
			return CodeBlock
		}
		if trimmed[:3] == "~~~" && trimmed[len(trimmed)-3:] == "~~~" {
			return CodeBlock
		}
	}
	if isBlockquote(block) {
		return Quote
	}
	if len(trimmed) >= 2 && (trimmed[:2] == "- " || trimmed[:2] == "* ") {
		return UnorderedList
	}
	if isOrderedList(trimmed) {
		return OrderedList
	}
	return Paragraph
}

///////////////////////////////////////
// Helper funcs for BlockToBlockType //
///////////////////////////////////////

func isThematicBreak(line string) bool {
	trimmed := strings.TrimSpace(line)
	if len(trimmed) < 3 {
		return false
	}
	char := rune(trimmed[0])
	if char != '*' && char != '-' && char != '_' {
		return false
	}
	count := 0
	for _, r := range trimmed {
		if r == char {
			count++
		} else if r != ' ' {
			return false
		}
	}
	return count >= 3
}

var re = regexp.MustCompile(`^\d+\. `)

func isOrderedList(block string) bool {
	split := strings.Split(block, "\n")
	firstLine := strings.TrimLeft(split[0], " ")
	return isListItemStart(firstLine) && (len(firstLine) > 0 && firstLine[0] >= '0' && firstLine[0] <= '9')
}

func isBlockquote(block string) bool {
	return strings.HasPrefix(block, ">")
}
