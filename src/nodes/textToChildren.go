package nodes

const maxNodeDepth = 100

func TextToChildren(s string) []TextNode {
	rawNodes := TextToTextNodes(s)

	var limited []TextNode
	for _, n := range rawNodes {
		if safe, ok := limitDepth(n, 0); ok {
			limited = append(limited, safe)
		}
	}
	return limited
}

// limitDepth adds recursion depth check
func limitDepth(n TextNode, depth int) (TextNode, bool) {
	if depth > maxNodeDepth {
		return TextNode{Text: "[MAX DEPTH EXCEEDED]", TextType: Text}, false
	}
	if len(n.Children) == 0 {
		return n, true
	}
	var newChildren []TextNode
	for _, c := range n.Children {
		limited, ok := limitDepth(c, depth+1)
		if ok {
			newChildren = append(newChildren, limited)
		}
	}
	n.Children = newChildren
	return n, true
}
