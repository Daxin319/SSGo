package nodes

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type Lnk struct {
	Text string
	Url  string
}

var linkReg = regexp.MustCompile(`\[(.*?)\]\((.*?)\)`)

func ExtractLinks(s string) []Lnk {
	var results []Lnk
	matches := linkReg.FindAllStringSubmatchIndex(s, -1)

	for _, m := range matches {
		start := m[0]
		// skip if preceded by '!'
		if start > 0 && s[start-1] == '!' {
			continue
		}

		text := s[m[2]:m[3]]
		url := s[m[4]:m[5]]
		results = append(results, Lnk{Text: text, Url: url})
	}
	return results
}

func SplitNodesLink(oldNodes []TextNode) ([]TextNode, error) {
	var newNodes []TextNode
	for _, node := range oldNodes {
		if node.TextType != Text {
			newNodes = append(newNodes, node)
			continue
		}

		links := ExtractLinks(node.Text)
		if len(links) == 0 {
			newNodes = append(newNodes, node)
			continue
		}

		for _, link := range links {
			if len(link.Text) == 0 || len(link.Url) == 0 {
				return []TextNode{}, errors.New("Invalid markdown, empty text or url field.")
			}
			formatted := fmt.Sprintf("[%s](%s)", link.Text, link.Url)
			splitText := strings.Split(node.Text, formatted)
			if len(splitText[0]) != 0 {
				newNode := TextNode{
					Text:     splitText[0],
					TextType: Text,
				}
				newNodes = append(newNodes, newNode)
			}
			newNode := TextNode{
				Text:     link.Text,
				TextType: Link,
				Url:      link.Url,
			}
			newNodes = append(newNodes, newNode)
			node.Text = splitText[1]
		}
		if len(node.Text) != 0 {
			newNodes = append(newNodes, node)
		}
	}
	return newNodes, nil
}
