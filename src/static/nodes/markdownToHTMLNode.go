package nodes

import (
	"main/src/static/blocks"
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
			node = HTMLNode{
				Tag:      "h" + strconv.Itoa(headerNum(blck)),
				Children: TextToChildren(blck),
			}
			bNodes = append(bNodes, node)

		case blocks.Paragraph:
			node = HTMLNode{
				Tag:      "p",
				Children: TextToChildren(blck),
			}
			bNodes = append(bNodes, node)

		case blocks.Code:
			child := HTMLNode{
				Tag:   "code",
				Value: strings.TrimLeft(strings.TrimRight(blck, "`"), "`\n"),
			}
			node = HTMLNode{
				Tag:      "pre",
				Children: []HTMLNode{child},
			}
			bNodes = append(bNodes, node)

		case blocks.Quote:
			split := strings.Split(blck, "\n")

			for _, item := range split {
				item = strings.TrimLeft(item, "> ")
			}
			joined := strings.Join(split, "\n")

			node = HTMLNode{
				Tag:      "blockquote",
				Children: TextToChildren(joined),
			}
			bNodes = append(bNodes, node)

		case blocks.UnorderedList:
			split := strings.Split(blck, "\n")

			for _, item := range split {
				item = strings.TrimLeft(item, ">* ")
			}
			joined := strings.Join(split, "\n")
			children := TextToChildren(joined)
			fixed := []HTMLNode{}

			for _, child := range children {
				fixed = append(fixed, child.fixLists())
			}

			node = HTMLNode{
				Tag:      "ul",
				Children: fixed,
			}
			bNodes = append(bNodes, node)

		case blocks.OrderedList:
			split := strings.Split(blck, "\n")

			for _, item := range split {
				item = strings.TrimLeft(item, " ")
			}
			joined := strings.Join(split, "\n")
			children := TextToChildren(joined)
			fixed := []HTMLNode{}

			for _, child := range children {
				fixed = append(fixed, child.fixLists())
			}

			node = HTMLNode{
				Tag:      "ol",
				Children: fixed,
			}
			bNodes = append(bNodes, node)

		default:
			continue
		}
	}
	return HTMLNode{Tag: "div", Children: bNodes}
}

/////////////////////////////
// Helper functions below //
///////////////////////////

func headerNum(block string) int {
	var n int

	for i, char := range block {
		if string(char) != "#" {
			n = i - 1
		}
	}

	return n
}

func (h *HTMLNode) fixLists() HTMLNode {
	return HTMLNode{
		Tag:   "li",
		Value: h.Value,
	}
}
