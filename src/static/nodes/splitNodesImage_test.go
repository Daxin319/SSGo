package nodes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractImages(t *testing.T) {
	input := "![rick roll](https://i.imgur.com/aKaOqIh.gif) and ![obi wan](https://i.imgur.com/fJRm4Vk.jpeg)"
	expected := []Img{
		{Alt: "rick roll", Url: "https://i.imgur.com/aKaOqIh.gif"},
		{Alt: "obi wan", Url: "https://i.imgur.com/fJRm4Vk.jpeg"},
	}
	assert.Equal(t, expected, ExtractImages(input))
}

func TestSplitNodesImage(t *testing.T) {
	tests := []struct {
		name     string
		input    []TextNode
		expected []TextNode
	}{
		{
			name: "single image",
			input: []TextNode{
				{Text: "Look at this ![cat](cat.jpg)", TextType: Text},
			},
			expected: []TextNode{
				{Text: "Look at this ", TextType: Text},
				{Text: "cat", TextType: Image, Url: "cat.jpg"},
			},
		},
		{
			name: "multiple images",
			input: []TextNode{
				{Text: "A ![one](1.png) and a ![two](2.png)", TextType: Text},
			},
			expected: []TextNode{
				{Text: "A ", TextType: Text},
				{Text: "one", TextType: Image, Url: "1.png"},
				{Text: " and a ", TextType: Text},
				{Text: "two", TextType: Image, Url: "2.png"},
			},
		},
		{
			name: "no images",
			input: []TextNode{
				{Text: "Just a paragraph.", TextType: Text},
			},
			expected: []TextNode{
				{Text: "Just a paragraph.", TextType: Text},
			},
		},
		{
			name: "image at start",
			input: []TextNode{
				{Text: "![img](img.png) is first.", TextType: Text},
			},
			expected: []TextNode{
				{Text: "img", TextType: Image, Url: "img.png"},
				{Text: " is first.", TextType: Text},
			},
		},
		{
			name: "image at end",
			input: []TextNode{
				{Text: "Ending with ![img](img.png)", TextType: Text},
			},
			expected: []TextNode{
				{Text: "Ending with ", TextType: Text},
				{Text: "img", TextType: Image, Url: "img.png"},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			out, err := SplitNodesImage(tc.input)
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, out)
		})
	}
}
