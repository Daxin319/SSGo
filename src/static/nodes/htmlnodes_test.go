package nodes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHTMLNodes(t *testing.T) { // 3 tests
	nodes := []HTMLNode{
		{
			Tag:      "p",
			Value:    "Hello, world!",
			Children: nil,
			Props:    map[string]string{},
		},
		// Repr: HTMLNode(p, Hello, world!, [], map[])

		{
			Tag:      "a",
			Value:    "Click here",
			Children: nil,
			Props: map[string]string{
				"href":  "https://www.google.com",
				"class": "link",
			},
		},
		// Repr: HTMLNode(a, Click here, [], map[class:link href:https://www.google.com])

		{
			Tag:   "div",
			Value: "",
			Children: []HTMLNode{
				{
					Tag:      "h1",
					Value:    "Title",
					Children: nil,
					Props:    map[string]string{"id": "main-title"},
				},
				{
					Tag:      "p",
					Value:    "This is a paragraph.",
					Children: nil,
					Props: map[string]string{
						"class": "content",
					},
				},
			},
			Props: map[string]string{"class": "container"},
		},
		// Repr: HTMLNode(div, , [{h1 Title [] map[id:main-title]} {p This is a paragraph. [] map[class:content]}], map[class:container])

	}

	for i := range nodes {
		switch i {
		case 0:
			assert.Equal(t, "HTMLNode(p, Hello, world!, [], map[])", nodes[i].Repr())
		case 1:
			assert.Equal(t, "HTMLNode(a, Click here, [], map[class:link href:https://www.google.com])", nodes[i].Repr())
		case 2:
			assert.Equal(t, "HTMLNode(div, , [{h1 Title [] map[id:main-title]} {p This is a paragraph. [] map[class:content]}], map[class:container])", nodes[i].Repr())
		}
	}
}
