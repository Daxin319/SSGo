package blocks

import (
	"regexp"
	"strconv"
	"strings"
)

type BlockType int

const (
	Paragraph BlockType = iota
	Heading
	Code
	Quote
	UnorderedList
	OrderedList
)

func BlockToBlockType(block string) BlockType {
	trimmed := strings.TrimLeft(strings.TrimRight(block, " \n"), " \n")

	if len(trimmed) >= 1 && trimmed[0] == '#' {
		n, i := HeaderNum(trimmed)
		if string(trimmed[n-1]) != "#" || string(trimmed[i]) != " " {
			return Paragraph
		}
		return Heading
	}
	if len(trimmed) >= 6 && trimmed[:3] == "```" && trimmed[len(trimmed)-3:] == "```" {
		return Code
	}
	if len(trimmed) >= 2 && trimmed[:2] == "> " {
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

var re = regexp.MustCompile(`^\d+\. `)

func isOrderedList(block string) bool {
	split := strings.Split(block, "\n")

	expected := 1
	for _, line := range split {
		trimmed := strings.TrimLeft(line, " ")
		match := re.FindStringSubmatch(trimmed)
		if len(match) < 1 {
			return false
		}
		trim := strings.TrimRight(match[0], ". ")

		n, err := strconv.Atoi(trim)
		if err != nil || n != expected {
			return false
		}
		expected++
	}
	return true
}
