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
			expected: "<div><pre><code>\nThis is text that _should_ remain\nthe **same** even with inline stuff\n</code></pre></div>",
		},
		{
			name:     "Empty string returns empty div",
			markdown: "",
			expected: "<div></div>",
		},
		{
			name:     "Only code block",
			markdown: "```\nfmt.Println(\"Hello\")\n```",
			expected: "<div><pre><code>\nfmt.Println(\"Hello\")\n</code></pre></div>",
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
			expected: "<div><ul><li><div><p>Item one</p></div></li><li><div><p>Item two</p></div></li><li><div><p>Item three</p></div></li></ul></div>",
		},
		{
			name:     "Ordered list",
			markdown: "1. First item\n2. Second item\n3. Third item\n\n",
			expected: "<div><ol><li><div><p>First item</p></div></li><li><div><p>Second item</p></div></li><li><div><p>Third item</p></div></li></ol></div>",
		},
		{
			name:     "H1 header",
			markdown: "# This is an H1\n\n",
			expected: "<div><h1>This is an H1</h1></div>",
		},
		{
			name:     "H2 header",
			markdown: "## Subheading H2\n\n",
			expected: "<div><h2>Subheading H2</h2></div>",
		},
		{
			name:     "H3 header",
			markdown: "### Heading Three\n\n",
			expected: "<div><h3>Heading Three</h3></div>",
		},
		{
			name:     "H4 header",
			markdown: "#### Fourth-level heading\n\n",
			expected: "<div><h4>Fourth-level heading</h4></div>",
		},
		{
			name:     "H5 header",
			markdown: "##### H5 test\n\n",
			expected: "<div><h5>H5 test</h5></div>",
		},
		{
			name:     "H6 header",
			markdown: "###### Tiny heading\n\n",
			expected: "<div><h6>Tiny heading</h6></div>",
		},
		{
			name:     "Header with inline formatting",
			markdown: "### Bold and _italic_ in header\n\n",
			expected: "<div><h3>Bold and <i>italic</i> in header</h3></div>",
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
