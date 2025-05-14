package nodes

import (
	"main/src/blocks"
	"strconv"
	"strings"
)

func MarkdownToHTMLNode(s string) TextNode {
	var node TextNode

	blcks := blocks.MarkdownToBlocks(s)
	bNodes := []TextNode{}

	for _, blck := range blcks {

		bType := blocks.BlockToBlockType(blck)

		switch bType {
		case blocks.Heading:
			trimmed := strings.TrimLeft(blck, "# ")
			node = TextNode{
				Tag:      "h" + strconv.Itoa(blocks.HeaderNum(blck)),
				Children: TextToChildren(trimmed),
			}
			bNodes = append(bNodes, node)

		case blocks.Paragraph:
			cleaned := CleanNewlines(blck)

			node = TextNode{
				Tag:      "p",
				Children: TextToChildren(cleaned),
			}
			bNodes = append(bNodes, node)

		case blocks.Code:
			if string(blck[len(blck)-1]) != "\n" {
				blck += "\n"
			}
			child := TextNode{
				Tag:   "code",
				Value: strings.TrimLeft(strings.TrimSuffix(blck, "```\n"), "`\n"),
			}
			node = TextNode{
				Tag:      "pre",
				Children: []TextNode{child},
			}
			bNodes = append(bNodes, node)

		case blocks.Quote:
			joined := CleanQuotes(blck)

			node = TextNode{
				Tag:      "blockquote",
				Children: TextToChildren(joined),
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
	return TextNode{Tag: "div", Children: bNodes}
}
