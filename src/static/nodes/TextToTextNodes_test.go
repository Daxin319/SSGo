package nodes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTextToTextNodes(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []TextNode
	}{
		{
			name:  "All formatting types",
			input: "This is `code`, **bold**, _italic_, ![alt](img.png), and [link](url).",
			expected: []TextNode{
				{Text: "This is ", TextType: Text},
				{Text: "code", TextType: Code},
				{Text: ", ", TextType: Text},
				{Text: "bold", TextType: Bold},
				{Text: ", ", TextType: Text},
				{Text: "italic", TextType: Italic},
				{Text: ", ", TextType: Text},
				{Text: "alt", TextType: Image, Url: "img.png"},
				{Text: ", and ", TextType: Text},
				{Text: "link", TextType: Link, Url: "url"},
				{Text: ".", TextType: Text},
			},
		},
		{
			name:  "Code disables nested parsing",
			input: "`**not bold**`",
			expected: []TextNode{
				{Text: "**not bold**", TextType: Code},
			},
		},
		{
			name:  "Bold wrapping italic delimiters",
			input: "**bold _not italic_**",
			expected: []TextNode{
				{Text: "bold _not italic_", TextType: Bold},
			},
		},
		{
			name:  "Italic wrapping bold delimiters",
			input: "_italic **not bold**_",
			expected: []TextNode{
				{Text: "italic **not bold**", TextType: Italic},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := TextToTextNodes(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
