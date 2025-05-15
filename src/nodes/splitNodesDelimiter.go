package nodes

import "strings"

func SplitNodesDelimiter(nodes []TextNode, delimiter string, textType enum) ([]TextNode, error) {
	var result []TextNode

	for _, node := range nodes {
		if node.TextType != Text || len(node.Children) > 0 {
			result = append(result, node)
			continue
		}

		text := node.Text
		start := 0
		for start < len(text) {
			open := indexOf(text, delimiter, start)
			if open == -1 {
				result = append(result, TextNode{
					Text:     text[start:],
					TextType: Text,
					Url:      node.Url,
				})
				break
			}

			if open > start {
				result = append(result, TextNode{
					Text:     text[start:open],
					TextType: Text,
					Url:      node.Url,
				})
			}

			innerStart := open + len(delimiter)
			close := indexOf(text, delimiter, innerStart)
			if close == -1 {
				result = append(result, TextNode{
					Text:     text[open:],
					TextType: Text,
					Url:      node.Url,
				})
				break
			}

			innerText := text[innerStart:close]
			children := TextToTextNodesWithExclusion(innerText, textType)

			hasVisibleContent := false
			for _, c := range children {
				if strings.TrimSpace(c.Text) != "" || c.Value != "" || len(c.Children) > 0 {
					hasVisibleContent = true
				}
			}

			if hasVisibleContent {
				if textType == Boldtalic {
					result = append(result, TextNode{
						TextType: Bold,
						Children: []TextNode{
							{TextType: Italic, Children: children},
						},
					})
				} else {
					result = append(result, TextNode{
						TextType: textType,
						Children: children,
					})
				}
			} else {
				result = append(result, TextNode{
					Text:     delimiter + innerText + delimiter,
					TextType: Text,
					Url:      node.Url,
				})
			}

			start = close + len(delimiter)
		}
	}

	return result, nil
}

func indexOf(text, delimiter string, start int) int {
	for i := start; i <= len(text)-len(delimiter); i++ {
		if text[i:i+len(delimiter)] == delimiter {
			return i
		}
	}
	return -1
}
