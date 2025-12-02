package inline

import (
	"fmt"

	"github.com/Daxin319/SSGo/src/nodes"
	"github.com/Daxin319/SSGo/src/tokenizer"
)

func TextToTextNodes(s string) []nodes.TextNode {
	tokens := tokenizer.TokenizeInline(s)
	fmt.Println("TOKENS:")
	for _, t := range tokens {
		fmt.Printf("kind=%q value=%q\n", t.Kind, t.Value)
	}
	return ParseInlineStack(tokens)
}
