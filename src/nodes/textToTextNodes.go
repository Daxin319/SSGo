package nodes

import (
	"fmt"
)

func TextToTextNodes(s string) []TextNode {
	toks := TokenizeInline(s)
	fmt.Println("TOKENS:")
	for _, t := range toks {
		fmt.Printf("kind=%q value=%q\n", t.kind, t.value)
	}
	return ParseInlineStack(toks)
}
