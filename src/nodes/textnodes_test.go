package nodes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTextNodes(t *testing.T) { // 6 tests
	nodes := []TextNode{
		{
			Text:     "This is a normal node",
			TextType: Text,
		},
		{
			Text:     "This is a bold node",
			TextType: Bold,
		},
		{
			Text:     "This is an italic node",
			TextType: Italic,
		},
		{
			Text:     "This is a code node",
			TextType: Code,
		},
		{
			Text:     "This is a link node",
			TextType: Link,
			Url:      "https://www.google.com",
		},
		{
			Text:     "This is an image node",
			TextType: Image,
			Url:      "https://http.cat/images/200.jpg",
		},
	}

	for i := range nodes {
		switch i {
		case 0:
			assert.Equal(t, "TextNode(This is a normal node, text, )", nodes[i].Repr())
		case 1:
			assert.Equal(t, "TextNode(This is a bold node, bold, )", nodes[i].Repr())
		case 2:
			assert.Equal(t, "TextNode(This is an italic node, italic, )", nodes[i].Repr())
		case 3:
			assert.Equal(t, "TextNode(This is a code node, code, )", nodes[i].Repr())
		case 4:
			assert.Equal(t, "TextNode(This is a link node, link, https://www.google.com)", nodes[i].Repr())
		case 5:
			assert.Equal(t, "TextNode(This is an image node, image, https://http.cat/images/200.jpg)", nodes[i].Repr())
		}
	}
}
