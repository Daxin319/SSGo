package nodes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarkdownToHTMLNode(t *testing.T) {
	tests := []struct {
		name     string
		markdown string
		expected string
	}{
		{
			name:     "Paragraphs with formatting",
			markdown: "This is **bolded** paragraph\ntext in a p\ntag here\n\nThis is another paragraph with _italic_ text and `code` here\n\n",
			expected: "<div><p>This is <b>bolded</b> paragraph text in a p tag here</p><p>This is another paragraph with <i>italic</i> text and <code>code</code> here</p></div>",
		},
		{
			name:     "Code block preserves raw text",
			markdown: "```\nThis is text that _should_ remain\nthe **same** even with inline stuff\n```\n",
			expected: "<div><pre><code>This is text that _should_ remain\nthe **same** even with inline stuff\n</code></pre></div>",
		},
		{
			name:     "Empty string returns empty div",
			markdown: "",
			expected: "<div></div>",
		},
		{
			name:     "Only code block",
			markdown: "```\nfmt.Println(\"Hello\")\n```",
			expected: "<div><pre><code>fmt.Println(\"Hello\")\n</code></pre></div>",
		},
		{
			name:     "Single paragraph no formatting",
			markdown: "Just a single paragraph.",
			expected: "<div><p>Just a single paragraph.</p></div>",
		},
		{
			name:     "Quote block",
			markdown: "> This is a quote.\n> It spans multiple lines.\n\n",
			expected: "<div><blockquote>This is a quote.\nIt spans multiple lines.</blockquote></div>",
		},
		{
			name:     "Unordered list",
			markdown: "- Item one\n- Item two\n- Item three\n\n",
			expected: "<div><ul><li>Item one</li><li>Item two</li><li>Item three</li></ul></div>",
		},
		{
			name:     "Ordered list",
			markdown: "1. First item\n2. Second item\n3. Third item\n\n",
			expected: "<div><ol><li>First item</li><li>Second item</li><li>Third item</li></ol></div>",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := MarkdownToHTMLNode(tt.markdown)
			html := h.ToHTML()
			assert.Equal(t, tt.expected, html)
		})
	}
}
