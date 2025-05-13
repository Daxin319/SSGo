package nodes

import "errors"

func TextNodeToHTMLNode(t *TextNode) (HTMLNode, error) {
	switch String(t.TextType) {
	case "text":
		return HTMLNode{
			Value: t.Text,
		}, nil

	case "bold":
		return HTMLNode{
			Tag:   "b",
			Value: t.Text,
		}, nil

	case "italic":
		return HTMLNode{
			Tag:   "i",
			Value: t.Text,
		}, nil

	case "boldtalic":
		return HTMLNode{
			Tag: "b",
			Children: []HTMLNode{
				{
					Tag:   "i",
					Value: t.Text,
				},
			},
		}, nil

	case "code":
		return HTMLNode{
			Tag:   "code",
			Value: t.Text,
		}, nil

	case "link":
		return HTMLNode{
			Tag:   "a",
			Value: t.Text,
			Props: map[string]string{"href": t.Url},
		}, nil

	case "image":
		return HTMLNode{
			Tag: "img",
			Props: map[string]string{
				"src": t.Url,
				"alt": t.Text,
			},
		}, nil

	default:
		return HTMLNode{}, errors.New("invalid text type")
	}

}
