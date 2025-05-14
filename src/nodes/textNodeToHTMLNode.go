package nodes

import "errors"

func TextNodeToHTMLNode(t *TextNode) (TextNode, error) {
	switch String(t.TextType) {
	case "text":
		return TextNode{
			Value: t.Text,
		}, nil

	case "bold":
		return TextNode{
			Tag:   "b",
			Value: t.Text,
		}, nil

	case "italic":
		return TextNode{
			Tag:   "i",
			Value: t.Text,
		}, nil

	case "boldtalic":
		return TextNode{
			Tag: "b",
			Children: []TextNode{
				{
					Tag:   "i",
					Value: t.Text,
				},
			},
		}, nil

	case "code":
		return TextNode{
			Tag:   "code",
			Value: t.Text,
		}, nil

	case "link":
		return TextNode{
			Tag:   "a",
			Value: t.Text,
			Props: map[string]string{"href": t.Url},
		}, nil

	case "image":
		return TextNode{
			Tag: "img",
			Props: map[string]string{
				"src": t.Url,
				"alt": t.Text,
			},
		}, nil

	default:
		return TextNode{}, errors.New("invalid text type")
	}

}
