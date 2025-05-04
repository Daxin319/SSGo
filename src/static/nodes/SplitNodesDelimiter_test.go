package nodes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitNodesDelimiter(t *testing.T) {
	tests := []struct {
		name      string
		input     []TextNode
		delimiter string
		textType  enum
		expected  []TextNode
	}{
		{
			name:      "inline code",
			delimiter: "`",
			textType:  Code,
			input: []TextNode{
				{Text: "This is text with a `code` element.", TextType: Text},
			},
			expected: []TextNode{
				{Text: "This is text with a ", TextType: Text},
				{Text: "code", TextType: Code},
				{Text: " element.", TextType: Text},
			},
		},
		{
			name:      "italic",
			delimiter: "_",
			textType:  Italic,
			input: []TextNode{
				{Text: "This is _italic_ text.", TextType: Text},
			},
			expected: []TextNode{
				{Text: "This is ", TextType: Text},
				{Text: "italic", TextType: Italic},
				{Text: " text.", TextType: Text},
			},
		},
		{
			name:      "bold",
			delimiter: "**",
			textType:  Bold,
			input: []TextNode{
				{Text: "This is **bold** text.", TextType: Text},
			},
			expected: []TextNode{
				{Text: "This is ", TextType: Text},
				{Text: "bold", TextType: Bold},
				{Text: " text.", TextType: Text},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			out, err := SplitNodesDelimiter(tc.input, tc.delimiter, tc.textType)
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, out)
		})
	}
}
