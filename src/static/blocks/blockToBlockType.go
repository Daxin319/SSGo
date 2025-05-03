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
	if block[:1] == "#" {
		return Heading
	}
	if block[:3] == "```" && block[len(block)-3:] == "```" {
		return Code
	}
	if block[:2] == "> " {
		return Quote
	}
	if block[:2] == "- " || block[:2] == "* " {
		return UnorderedList
	}
	if isOrderedList(block) {
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
