package nodes

import (
	"fmt"
	"main/src/blocks"
	"strconv"
	"strings"
)

func MarkdownToHTMLNode(s string) TextNode {
	fmt.Println("RAW BYTES:")
	for i, r := range s {
		fmt.Printf("%02d: %q (%[2]U)\n", i, r)
	}

	var node TextNode
	blcks := blocks.MarkdownToBlocks(s)
	bNodes := []TextNode{}

	for _, blck := range blcks {
		bType := blocks.BlockToBlockType(blck)

		switch bType {
		case blocks.Heading:
			trimmed := strings.TrimLeft(blck, "# ")
			n, _ := blocks.HeaderNum(blck)
			children := TextToChildren(trimmed)
			node = TextNode{
				Tag:      "h" + strconv.Itoa(n),
				Children: mapToHTMLChildren(children, 0),
			}
			bNodes = append(bNodes, node)

		case blocks.Paragraph:
			fmt.Println(">>> Paragraph block")
			cleaned := CleanNewlines(blck)
			children := TextToChildren(cleaned)
			node = TextNode{
				Tag:      "p",
				Children: mapToHTMLChildren(children, 0),
			}
			bNodes = append(bNodes, node)

		case blocks.Code:
			if !strings.HasSuffix(blck, "\n") {
				blck += "\n"
			}
			content := strings.TrimLeft(strings.TrimSuffix(blck, "```\n"), "`\n")
			codeNode := TextNode{
				Tag:      "code",
				Children: []TextNode{{Text: content, TextType: Text}},
			}
			node = TextNode{
				Tag:      "pre",
				Children: []TextNode{codeNode},
			}
			bNodes = append(bNodes, node)

		case blocks.Quote:
			joined := CleanQuotes(blck)
			children := TextToChildren(joined)
			node = TextNode{
				Tag:      "blockquote",
				Children: mapToHTMLChildren(children, 0),
			}
			bNodes = append(bNodes, node)

		case blocks.UnorderedList:
			children := CleanLists(blck)
			node = TextNode{
				Tag:      "ul",
				Children: children,
			}
			bNodes = append(bNodes, node)

		case blocks.OrderedList:
			children := CleanLists(blck)
			node = TextNode{
				Tag:      "ol",
				Children: children,
			}
			bNodes = append(bNodes, node)

		default:
			continue
		}
	}

	root := TextNode{Tag: "div", Children: bNodes}
	return root
}
