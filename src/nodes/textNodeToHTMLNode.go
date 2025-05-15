package nodes

func textNodeToHTMLNodeInternal(n TextNode, depth int) TextNode {
	if depth > 1000 {
		return TextNode{Tag: "", Text: "[DEPTH LIMIT HIT]"}
	}

	switch n.TextType {
	case Text:
		return TextNode{Tag: "", Text: n.Text}
	case Bold:
		return TextNode{Tag: "b", Children: mapToHTMLChildren(n.Children, depth+1)}
	case Italic:
		return TextNode{Tag: "i", Children: mapToHTMLChildren(n.Children, depth+1)}
	case Code:
		return TextNode{Tag: "code", Children: mapToHTMLChildren(n.Children, depth+1)}
	case Boldtalic:
		return TextNode{
			Tag: "b",
			Children: []TextNode{
				{Tag: "i", Children: mapToHTMLChildren(n.Children, depth+1)},
			},
		}
	default:
		return TextNode{
			Tag:      n.Tag,
			Text:     n.Text,
			Value:    n.Value,
			Children: mapToHTMLChildren(n.Children, depth+1),
			Props:    n.Props,
		}
	}
}

func mapToHTMLChildren(children []TextNode, depth int) []TextNode {
	var out []TextNode
	for _, c := range children {
		out = append(out, textNodeToHTMLNodeInternal(c, depth))
	}
	return out
}
