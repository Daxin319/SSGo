package nodes

func TextToChildren(s string) []TextNode {
	var final []TextNode

	nodes := TextToTextNodes(s)

	for _, node := range nodes {
		newNode, _ := TextNodeToHTMLNode(&node)
		final = append(final, newNode)
	}

	return final
}
