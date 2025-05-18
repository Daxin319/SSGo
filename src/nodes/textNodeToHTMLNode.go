package nodes

func mapToHTMLChildren(children []TextNode, depth int) []TextNode {
	out := make([]TextNode, 0, len(children))
	for _, c := range children {
		out = append(out, textNodeToHTMLNodeInternal(c, depth+1))
	}
	return out
}

func textNodeToHTMLNodeInternal(n TextNode, depth int) TextNode {
	if depth > 1000 {
		return n
	}

	switch n.TextType {
	case Strikethrough:
		n.Tag = "s"
		n.Children = mapToHTMLChildren(n.Children, depth+1)
		n.Text = ""

	case Subscript:
		n.Tag = "sub"
		n.Children = mapToHTMLChildren(n.Children, depth+1)
		n.Text = ""

	case BoldItalic:
		em := TextNode{Tag: "em", Children: mapToHTMLChildren(n.Children, depth+1)}
		n.Tag = "strong"
		n.Children = []TextNode{em}
		n.Text = ""

	case Bold:
		n.Tag = "strong"
		n.Children = mapToHTMLChildren(n.Children, depth+1)
		n.Text = ""

	case Italic:
		n.Tag = "em"
		n.Children = mapToHTMLChildren(n.Children, depth+1)
		n.Text = ""

	case Code:
		n.Tag = "code"
		var content string
		if n.Text != "" {
			content = n.Text
		} else {
			content = n.Value
		}
		n.Children = []TextNode{{Text: content, TextType: Text}}
		n.Text = ""
		n.Value = ""

	case Link:
		n.Tag = "a"
		if n.Props == nil {
			n.Props = make(map[string]string)
		}
		n.Props["href"] = n.Url
		n.Children = mapToHTMLChildren(n.Children, depth+1)
		n.Text = ""

	case Image:
		n.Tag = "img"
		if n.Props == nil {
			n.Props = make(map[string]string)
		}
		n.Props["src"] = n.Url
		alt := ""
		for _, c := range n.Children {
			alt += c.Text
		}
		n.Props["alt"] = alt
		n.Children = nil
		n.Text = ""

	default:
		n.Children = mapToHTMLChildren(n.Children, depth+1)
	}
	return n
}
