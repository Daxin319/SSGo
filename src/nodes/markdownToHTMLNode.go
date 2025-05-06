package nodes

import (
	"main/src/blocks"
	"strconv"
	"strings"
)

func MarkdownToHTMLNode(s string) HTMLNode {
	var node HTMLNode

	blcks := blocks.MarkdownToBlocks(s)
	bNodes := []HTMLNode{}

	for _, blck := range blcks {

		bType := blocks.BlockToBlockType(blck)

		switch bType {
		case blocks.Heading:
			trimmed := strings.TrimLeft(blck, "# ")
			node = HTMLNode{
				Tag:      "h" + strconv.Itoa(blocks.HeaderNum(blck)),
				Children: TextToChildren(trimmed),
			}
			bNodes = append(bNodes, node)

		case blocks.Paragraph:
			cleaned := CleanNewlines(blck)

			node = HTMLNode{
				Tag:      "p",
				Children: TextToChildren(cleaned),
			}
			bNodes = append(bNodes, node)

		case blocks.Code:
			if string(blck[len(blck)-1]) != "\n" {
				blck += "\n"
			}
			child := HTMLNode{
				Tag:   "code",
				Value: strings.TrimLeft(strings.TrimSuffix(blck, "```\n"), "`"),
			}
			node = HTMLNode{
				Tag:      "pre",
				Children: []HTMLNode{child},
			}
			bNodes = append(bNodes, node)

		case blocks.Quote:
			joined := CleanQuotes(blck)

			node = HTMLNode{
				Tag:      "blockquote",
				Children: TextToChildren(joined),
			}
			bNodes = append(bNodes, node)

		case blocks.UnorderedList:
			children := CleanLists(blck)

			node = HTMLNode{
				Tag:      "ul",
				Children: children,
			}
			bNodes = append(bNodes, node)

		case blocks.OrderedList:
			children := CleanLists(blck)

			node = HTMLNode{
				Tag:      "ol",
				Children: children,
			}
			bNodes = append(bNodes, node)

		default:
			continue
		}
	}
	return HTMLNode{Tag: "div", Children: bNodes}
}
