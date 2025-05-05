package nodes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTextNodeToHTMLNode(t *testing.T) {
	tests := []struct {
		name     string
		input    *TextNode
		expected string
	}{
		{
			name: "plain text",
			input: &TextNode{
				Text:     "This is a text node.",
				TextType: Text,
			},
			expected: "HTMLNode(, This is a text node., [], map[])",
		},
		{
			name: "link",
			input: &TextNode{
				Text:     "Example",
				TextType: Link,
				Url:      "https://example.com",
			},
			expected: "HTMLNode(a, Example, [], map[href:https://example.com])",
		},
		{
			name: "image",
			input: &TextNode{
				Text:     "An image",
				TextType: Image,
				Url:      "img.png",
			},
			expected: "HTMLNode(img, , [], map[alt:An image src:img.png])",
		},
		{
			name: "bold",
			input: &TextNode{
				Text:     "Bold!",
				TextType: Bold,
			},
			expected: "HTMLNode(b, Bold!, [], map[])",
		},
		{
			name: "italic",
			input: &TextNode{
				Text:     "Italic!",
				TextType: Italic,
			},
			expected: "HTMLNode(i, Italic!, [], map[])",
		},
		{
			name: "code",
			input: &TextNode{
				Text:     "fmt.Println()",
				TextType: Code,
			},
			expected: "HTMLNode(code, fmt.Println(), [], map[])",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			n, _ := TextNodeToHTMLNode(tc.input)
			assert.Equal(t, tc.expected, n.Repr())
		})
	}
}
