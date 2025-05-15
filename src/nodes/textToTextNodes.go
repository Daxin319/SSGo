package nodes

func TextToTextNodes(text string) []TextNode {
	return ParseInline(text)
}
