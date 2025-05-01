package nodes

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type Img struct {
	Alt string
	Url string
}

var imgReg = regexp.MustCompile(`!\[([^\[\]]*)\]\(([^\(\)]*)\)`)

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
