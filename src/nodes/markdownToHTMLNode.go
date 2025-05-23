package nodes

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var reg = regexp.MustCompile(`\\[ \t]*\n`)

func MarkdownToHTMLNode(input string) TextNode {
	lBreaks := reg.ReplaceAllString(input, "<br />\n")
	s := SanitizeNulls(lBreaks)
	fmt.Println("RAW BYTES:")
	for i, r := range s {
		fmt.Printf("%02d: %q (%[2]U)\n", i, r)
	}

	var node TextNode
	blcks := MarkdownToBlocks(s)
	bNodes := []TextNode{}

	for _, blck := range blcks {
		bType := BlockToBlockType(blck)

		switch bType {
		case Heading:
			trimmed := strings.TrimLeft(blck, "# ")
			n, _ := HeaderNum(blck)
			children := TextToChildren(trimmed)
			node = TextNode{
				Tag:      "h" + strconv.Itoa(n),
				Children: MapToHTMLChildren(children, 0),
			}
			bNodes = append(bNodes, node)

		case Paragraph:
			cleaned := CleanNewlines(blck)
			children := TextToChildren(cleaned)
			node = TextNode{
				Tag:      "p",
				Children: MapToHTMLChildren(children, 0),
			}
			bNodes = append(bNodes, node)

		case CodeBlock:
			lines := strings.Split(blck, "\n")
			body := ""
			if len(lines) > 2 {
				raw := strings.Join(lines[1:len(lines)-1], "\n")
				body = UnescapeString(UnescapeString(raw))
			}
			codeNode := TextNode{
				Tag:      "code",
				Props:    make(map[string]string),
				Children: []TextNode{{Text: body, TextType: Text}},
			}
			node = TextNode{
				Tag:      "pre",
				Children: []TextNode{codeNode},
			}
			bNodes = append(bNodes, node)

		case Quote:
			joined := CleanQuotes(blck)
			children := TextToChildren(joined)
			node = TextNode{
				Tag:      "blockquote",
				Children: MapToHTMLChildren(children, 0),
			}
			bNodes = append(bNodes, node)

		case UnorderedList:
			children := CleanLists(blck)
			node = TextNode{
				Tag:      "ul",
				Children: children,
			}
			bNodes = append(bNodes, node)

		case OrderedList:
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
