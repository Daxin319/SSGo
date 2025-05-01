package nodes

import (
	"bytes"
	"errors"
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
		if idx > 0 {
			newNode := TextNode{
				Text:     node.Text[:idx],
				TextType: Text,
				Url:      node.Url,
			}
			newNodes = append(newNodes, newNode)
		}

		if len(delimiter) > 1 {
			idx += len(delimiter) - 1
		}

		idx2 := bytes.Index([]byte(node.Text[idx+1:]), []byte(delimiter))
		if idx2 == -1 {
			return []TextNode{}, errors.New("no closing delimiter found, invalid markdown syntax")
		}
		newNode := TextNode{
			Text:     node.Text[idx+1 : idx+1+idx2],
			TextType: textType,
			Url:      node.Url,
		}
		newNodes = append(newNodes, newNode)

		if len(delimiter) > 1 {
			idx2 += len(delimiter) - 1
		}

		if len(node.Text[idx+idx2+2:]) < 1 {
			continue
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
