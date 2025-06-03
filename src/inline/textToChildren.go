package inline

import "github.com/Daxin319/SSGo/src/nodes"

func TextToChildren(s string) []nodes.TextNode {
	return TextToTextNodes(s)
}
