package blocks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarkdownToBlocks(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:  "Single block",
			input: "  This is a single paragraph.  ",
			expected: []string{
				"This is a single paragraph.",
			},
		},
		{
			name:  "Two paragraphs with trimming and newline preservation",
			input: " First paragraph. \n\n  Second paragraph.\n ",
			expected: []string{
				"First paragraph.",
				"Second paragraph.\n",
			},
		},
		{
			name:  "Empty and whitespace-only blocks",
			input: "Valid block.\n\n   \n\n\n\nAnother one.\n\n\n",
			expected: []string{
				"Valid block.",
				"Another one.",
			},
		},
		{
			name:  "Block with just newlines is removed",
			input: "\n\n\n\nMiddle block.\n\n\n\n",
			expected: []string{
				"Middle block.",
			},
		},
		{
			name:     "All whitespace and newlines",
			input:    " \n \n \n\n \n",
			expected: []string{},
		},
		{
			name:  "Mixed formatting with lists and preserved newlines",
			input: "This is **bolded** paragraph\n\nThis is another paragraph with _italic_ text and `code` here\nThis is the same paragraph on a new line\n\n- This is a list\n- with items",
			expected: []string{
				"This is **bolded** paragraph",
				"This is another paragraph with _italic_ text and `code` here\nThis is the same paragraph on a new line",
				"- This is a list\n- with items",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MarkdownToBlocks(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
