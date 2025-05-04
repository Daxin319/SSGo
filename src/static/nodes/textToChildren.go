package nodes

func TextToChildren(s string) []HTMLNode {
	var final []HTMLNode

	nodes := TextToTextNodes(s)

	for _, node := range nodes {
		newNode, _ := TextNodeToHTMLNode(&node)
		final = append(final, newNode)
	}

	return final
}
