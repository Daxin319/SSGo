package nodes

import (
	"bytes"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

func SplitNodesDelimiter(oldNodes []TextNode, delimiter string, textType enum) ([]TextNode, error) {
	var newNodes []TextNode
	for _, node := range oldNodes {
		if node.TextType != Text {
			newNodes = append(newNodes, node)
			continue
		}

		idx := bytes.Index([]byte(node.Text), []byte(delimiter))
		if idx == -1 {
			newNodes = append(newNodes, node)
			continue
		}
		newNode := TextNode{
			Text:     node.Text[:idx],
			TextType: Text,
			Url:      node.Url,
		}
		newNodes = append(newNodes, newNode)

		if len(delimiter) > 1 {
			idx += len(delimiter) - 1
		}

		idx2 := bytes.Index([]byte(node.Text[idx+1:]), []byte(delimiter))
		if idx2 == -1 {
			return []TextNode{}, errors.New("no closing delimiter found, invalid markdown syntax")
		}
		newNode = TextNode{
			Text:     node.Text[idx+1 : idx+1+idx2],
			TextType: textType,
			Url:      node.Url,
		}
		newNodes = append(newNodes, newNode)

		if len(delimiter) > 1 {
			idx2 += len(delimiter) - 1
		}

		finalNode := TextNode{
			Text:     node.Text[idx+idx2+2:],
			TextType: Text,
			Url:      node.Url,
		}
		recurse, err := SplitNodesDelimiter([]TextNode{finalNode}, delimiter, textType)
		if err != nil {
			return []TextNode{}, errors.New("invalid markdown syntax, unmatched delimiter found")
		}
		newNodes = append(newNodes, recurse...)
	}

	return newNodes, nil

}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////
//                                      Images and Links Handled Below                                     //
/////////////////////////////////////////////////////////////////////////////////////////////////////////////

type Img struct {
	Alt string
	Url string
}

var imgReg = regexp.MustCompile(`!\[([^\[\]]*)\]\(([^\(\)]*)\)`)

type Lnk struct {
	Text string
	Url  string
}

var linkReg = regexp.MustCompile(`\[(.*?)\]\((.*?)\)`)

func ExtractImages(s string) []Img {
	var results []Img
	matches := imgReg.FindAllStringSubmatch(s, -1)
	for _, m := range matches {
		if len(m) == 3 {
			results = append(results, Img{Alt: m[1], Url: m[2]})
		}
	}
	return results
}

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

func SplitNodesImage(oldNodes []TextNode) ([]TextNode, error) {
	var newNodes []TextNode
	for _, node := range oldNodes {
		if node.TextType != Text {
			newNodes = append(newNodes, node)
			continue
		}

		images := ExtractImages(node.Text)
		if len(images) == 0 {
			newNodes = append(newNodes, node)
			continue
		}

		for _, image := range images {
			if len(image.Alt) == 0 || len(image.Url) == 0 {
				return []TextNode{}, errors.New("Invalid image markdown. Empty alt text or URL")
			}
			formatted := fmt.Sprintf("![%s](%s)", image.Alt, image.Url)
			splitText := strings.Split(node.Text, formatted)
			if len(splitText[0]) != 0 {
				newNode := TextNode{
					Text:     splitText[0],
					TextType: Text,
				}
				newNodes = append(newNodes, newNode)
			}
			newNode := TextNode{
				Text:     image.Alt,
				TextType: Image,
				Url:      image.Url,
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
