package nodes

func TextToTextNodes(s string) []TextNode {
	toks := tokenizeInline(s)
	return parseInlineStack(toks)
}
