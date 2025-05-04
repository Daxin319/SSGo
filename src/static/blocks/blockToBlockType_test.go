package blocks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBlockToBlockType(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected BlockType
	}{
		{
			name:     "Paragraph block",
			input:    "Just some regular text with no formatting.",
			expected: Paragraph,
		},
		{
			name:     "Heading block",
			input:    "# This is a heading",
			expected: Heading,
		},
		{
			name:     "Code block",
			input:    "```\ncode goes here\n```",
			expected: Code,
		},
		{
			name:     "Quote block",
			input:    "> This is a quote.",
			expected: Quote,
		},
		{
			name:     "Unordered list with dash",
			input:    "- First item\n- Second item",
			expected: UnorderedList,
		},
		{
			name:     "Unordered list with asterisk",
			input:    "* Item A\n* Item B",
			expected: UnorderedList,
		},
		{
			name:     "Ordered list with number",
			input:    "1. First\n2. Second",
			expected: OrderedList,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := BlockToBlockType(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
