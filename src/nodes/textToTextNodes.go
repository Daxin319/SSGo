package nodes

// TextToTextNodes converts a markdown string into a slice of TextNode representing inline structure.
func TextToTextNodes(s string) []TextNode {
	tokens := tokenizeInline(s)
	nodes, _ := parseTokens(tokens, 0, "")
	return nodes
}
