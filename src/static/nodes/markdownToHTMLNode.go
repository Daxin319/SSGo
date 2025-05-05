package nodes

import (
	"bytes"
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
			trimmed := strings.TrimLeft(blck, "# ")
			node = HTMLNode{
				Tag:      "h" + strconv.Itoa(headerNum(blck)),
				Children: TextToChildren(trimmed),
			}
			bNodes = append(bNodes, node)

		case blocks.Paragraph:
			split := strings.Split(blck, "\n")
			var replaced []string

			for _, line := range split {
				if len(line) == 0 || line == " " || line == "\n" {
					continue
				}
				replaced = append(replaced, line)
			}
			joined := strings.Join(replaced, " ")

			node = HTMLNode{
				Tag:      "p",
				Children: TextToChildren(joined),
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
			split := strings.Split(blck, "\n")
			var fixed []string

			for _, item := range split {
				trimmed := strings.TrimLeft(item, "> ")
				fixed = append(fixed, trimmed)
			}
			joined := strings.Join(fixed, "\n")

			node = HTMLNode{
				Tag:      "blockquote",
				Children: TextToChildren(joined),
			}
			bNodes = append(bNodes, node)

		case blocks.UnorderedList:
			split := strings.Split(blck, "\n")
			var children []HTMLNode

			for _, item := range split {
				trimmed := strings.TrimLeft(item, "-* ")
				child := HTMLNode{
					Tag:   "li",
					Value: trimmed,
				}
				children = append(children, child)
			}

			node = HTMLNode{
				Tag:      "ul",
				Children: children,
			}
			bNodes = append(bNodes, node)

		case blocks.OrderedList:
			split := strings.Split(blck, "\n")
			var children []HTMLNode

			for _, item := range split {
				trim1 := strings.TrimLeft(item, " ")
				wsIdx := bytes.Index([]byte(trim1), []byte(" "))
				trimmed := trim1[wsIdx+1:]

				child := HTMLNode{
					Tag:   "li",
					Value: trimmed,
				}
				children = append(children, child)
			}

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
