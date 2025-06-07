package inline

import "github.com/Daxin319/SSGo/src/nodes"

// wrap creates a new TextNode with the appropriate type based on the marker
// and wraps the children nodes inside it.
func wrap(marker string, children []nodes.TextNode) nodes.TextNode {
	switch marker {
	case "*", "_":
		return nodes.TextNode{TextType: nodes.Italic, Children: children}
	case "**", "__":
		return nodes.TextNode{TextType: nodes.Bold, Children: children}
	case "***", "___":
		return nodes.TextNode{TextType: nodes.BoldItalic, Children: children}
	case "~~":
		return nodes.TextNode{TextType: nodes.Strikethrough, Children: children}
	case "~":
		return nodes.TextNode{TextType: nodes.Subscript, Children: children}
	case "^":
		return nodes.TextNode{TextType: nodes.Superscript, Children: children}
	case "==":
		return nodes.TextNode{TextType: nodes.Highlight, Children: children}
	}
	return nodes.TextNode{TextType: nodes.Text, Text: marker}
}
