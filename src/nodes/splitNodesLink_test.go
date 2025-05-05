package nodes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractLinks(t *testing.T) {
	input := "[Google](https://google.com) and [GitHub](https://github.com)"
	expected := []Lnk{
		{Text: "Google", Url: "https://google.com"},
		{Text: "GitHub", Url: "https://github.com"},
	}
	assert.Equal(t, expected, ExtractLinks(input))
}

func TestSplitNodesLink(t *testing.T) {
	tests := []struct {
		name     string
		input    []TextNode
		expected []TextNode
	}{
		{
			name: "single link",
			input: []TextNode{
				{Text: "Check [this](https://site.com)", TextType: Text},
			},
			expected: []TextNode{
				{Text: "Check ", TextType: Text},
				{Text: "this", TextType: Link, Url: "https://site.com"},
			},
		},
		{
			name: "multiple links",
			input: []TextNode{
				{Text: "Links: [a](1.com), [b](2.com)", TextType: Text},
			},
			expected: []TextNode{
				{Text: "Links: ", TextType: Text},
				{Text: "a", TextType: Link, Url: "1.com"},
				{Text: ", ", TextType: Text},
				{Text: "b", TextType: Link, Url: "2.com"},
			},
		},
		{
			name: "no links",
			input: []TextNode{
				{Text: "Nothing here.", TextType: Text},
			},
			expected: []TextNode{
				{Text: "Nothing here.", TextType: Text},
			},
		},
		{
			name: "link at start",
			input: []TextNode{
				{Text: "[top](t.com) start.", TextType: Text},
			},
			expected: []TextNode{
				{Text: "top", TextType: Link, Url: "t.com"},
				{Text: " start.", TextType: Text},
			},
		},
		{
			name: "link at end",
			input: []TextNode{
				{Text: "See [this](final.com)", TextType: Text},
			},
			expected: []TextNode{
				{Text: "See ", TextType: Text},
				{Text: "this", TextType: Link, Url: "final.com"},
			},
		},
		{
			name: "link after image",
			input: []TextNode{
				{Text: "![alt](x.png) then [link](url)", TextType: Text},
			},
			expected: []TextNode{
				{Text: "![alt](x.png) then ", TextType: Text},
				{Text: "link", TextType: Link, Url: "url"},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			out, err := SplitNodesLink(tc.input)
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, out)
		})
	}
}
