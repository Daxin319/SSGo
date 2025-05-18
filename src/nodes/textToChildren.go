package nodes

func TextToChildren(s string) []TextNode {
	rawNodes := TextToTextNodes(s)
	return rawNodes
}
