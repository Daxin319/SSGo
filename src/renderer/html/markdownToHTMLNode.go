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

var reg = regexp.MustCompile(`(?:\\[ \t]*\n|[ ]{2,}\n)`)

func MarkdownToHTMLNode(input string) nodes.TextNode {
	lBreaks := reg.ReplaceAllString(input, "<br />\n")
	s := blocks.SanitizeNulls(lBreaks)
	fmt.Println("RAW BYTES:")
	for i, r := range s {
		fmt.Printf("%02d: %q (%[2]U)\n", i, r)
	}

	var node nodes.TextNode
	blcks := blocks.MarkdownToBlocks(s)
	bNodes := []nodes.TextNode{}

	// Import the HTML tag/comment regex from blocks
	var htmlTagOrCommentRe = regexp.MustCompile(`^<(?:!--[\s\S]*?--|/?[a-zA-Z][a-zA-Z0-9-]*(?:\s+[^<>]*)?)>$`)

	for _, blck := range blcks {
		if htmlTagOrCommentRe.MatchString(strings.TrimSpace(blck)) {
			// Render as raw HTML node
			node = nodes.TextNode{
				Text:     blck,
				TextType: nodes.RawHTML,
			}
			bNodes = append(bNodes, node)
			continue
		}
		bType := blocks.BlockToBlockType(blck)

		switch bType {
		case blocks.ThematicBreak:
			node = nodes.TextNode{
				Tag:   "hr",
				Props: make(map[string]string),
			}
			bNodes = append(bNodes, node)

		case blocks.Heading:
			trimmed := strings.TrimLeft(blck, "# ")
			n, _ := blocks.HeaderNum(blck)
			children := inline.TextToChildren(trimmed)
			node = nodes.TextNode{
				Tag:      "h" + strconv.Itoa(n),
				Children: nodes.MapToHTMLChildren(children, 0),
			}
			bNodes = append(bNodes, node)

		case blocks.Paragraph:
			cleaned := blocks.CleanNewlines(blck)
			children := inline.TextToChildren(cleaned)
			node = nodes.TextNode{
				Tag:      "p",
				Children: nodes.MapToHTMLChildren(children, 0),
			}
			bNodes = append(bNodes, node)

		case blocks.CodeBlock:
			lines := strings.Split(blck, "\n")
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
			bNodes = append(bNodes, node)

		case blocks.Quote:
			joined := blocks.CleanQuotes(blck)
			children := inline.TextToChildren(joined)
			node = nodes.TextNode{
				Tag:      "blockquote",
				Children: nodes.MapToHTMLChildren(children, 0),
			}
			bNodes = append(bNodes, node)

		case blocks.UnorderedList:
			children := blocks.CleanLists(blck)
			node = nodes.TextNode{
				Tag:      "ul",
				Children: children,
			}
			bNodes = append(bNodes, node)

		case blocks.OrderedList:
			children := blocks.CleanLists(blck)
			node = nodes.TextNode{
				Tag:      "ol",
				Children: children,
			}
			bNodes = append(bNodes, node)

		default:
			continue
		}
	}

	root := nodes.TextNode{Tag: "div", Children: bNodes}
	return root
}
