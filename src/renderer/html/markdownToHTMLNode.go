package html

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/Daxin319/SSGo/src/blocks"
	"github.com/Daxin319/SSGo/src/inline"
	"github.com/Daxin319/SSGo/src/nodes"
)

var reg = regexp.MustCompile(`(\S)[ \t]*\\\n|(\S) {2,}\n`)

func MarkdownToHTMLNode(input string) nodes.TextNode {
	lBreaks := reg.ReplaceAllString(input, "$1$2<br />\n")
	s := blocks.SanitizeNulls(lBreaks)
	fmt.Println("RAW BYTES:")
	for i, r := range s {
		fmt.Printf("%02d: %q (%[2]U)\n", i, r)
	}

	toBlocks := blocks.MarkdownToBlocks(s)
	// First pass: extract reference definitions from all blocks
	for _, block := range toBlocks {
		blocks.ExtractRefDefs(block)
	}

	blockNodes := markdownToBlockNodes(toBlocks)

	root := nodes.TextNode{Tag: "div", Children: blockNodes}
	return root
}

func markdownToBlockNodes(toBlocks []string) []nodes.TextNode {
	var blockNodes []nodes.TextNode
	var node nodes.TextNode

	// Import the HTML tag/comment regex from blocks
	var htmlTagOrCommentRe = regexp.MustCompile(`^<(?:!--[\s\S]*?--|/?[a-zA-Z][a-zA-Z0-9-]*(?:\s+[^<>]*)?)>$`)

	for _, block := range toBlocks {
		block = blocks.ExtractRefDefs(block)
		if strings.TrimSpace(block) == "" {
			continue
		}
		if htmlTagOrCommentRe.MatchString(strings.TrimSpace(block)) {
			// Render as a raw HTML node without any processing
			node = nodes.TextNode{
				Text:     block,
				TextType: nodes.RawHTML,
				Tag:      "", // Ensure no tag is set for raw HTML
			}
			blockNodes = append(blockNodes, node)
			continue
		}
		blockType := blocks.BlockToBlockType(block)

		switch blockType {
		case blocks.ThematicBreak:
			node = nodes.TextNode{
				Tag:   "hr",
				Props: make(map[string]string),
			}
			blockNodes = append(blockNodes, node)

		case blocks.Heading:
			trimmed := strings.TrimLeft(block, "# ")
			n, _ := blocks.HeaderNum(block)
			children := inline.TextToChildren(trimmed)
			node = nodes.TextNode{
				Tag:      "h" + strconv.Itoa(n),
				Children: nodes.MapToHTMLChildren(children, 0),
			}
			blockNodes = append(blockNodes, node)

		case blocks.Paragraph:
			cleaned := blocks.CleanNewlines(block)
			children := inline.TextToChildren(cleaned)
			node = nodes.TextNode{
				Tag:      "p",
				Children: nodes.MapToHTMLChildren(children, 0),
			}
			blockNodes = append(blockNodes, node)

		case blocks.CodeBlock:
			lines := strings.Split(block, "\n")
			body := ""
			if len(lines) > 2 {
				raw := strings.Join(lines[1:len(lines)-1], "\n")
				body = nodes.UnescapeString(raw)
			}
			codeNode := nodes.TextNode{
				Tag:      "code",
				Props:    make(map[string]string),
				Children: []nodes.TextNode{{Text: body, TextType: nodes.Text}},
			}
			node = nodes.TextNode{
				Tag:      "pre",
				Children: []nodes.TextNode{codeNode},
			}
			blockNodes = append(blockNodes, node)

		case blocks.Quote:
			joined := blocks.CleanQuotes(block)
			subBlocks := blocks.MarkdownToBlocks(joined)
			children := markdownToBlockNodes(subBlocks)
			node = nodes.TextNode{
				Tag:      "blockquote",
				Children: children,
			}
			blockNodes = append(blockNodes, node)

		case blocks.UnorderedList:
			children := blocks.CleanLists(block, func(s string) []nodes.TextNode {
				return markdownToBlockNodes(blocks.MarkdownToBlocks(s))
			})
			node = nodes.TextNode{
				Tag:      "ul",
				Children: children,
			}
			blockNodes = append(blockNodes, node)

		case blocks.OrderedList:
			children := blocks.CleanLists(block, func(s string) []nodes.TextNode {
				return markdownToBlockNodes(blocks.MarkdownToBlocks(s))
			})
			node = nodes.TextNode{
				Tag:      "ol",
				Children: children,
			}
			blockNodes = append(blockNodes, node)

		default:
			continue
		}
	}
	return blockNodes
}
