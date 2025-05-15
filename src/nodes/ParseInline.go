package nodes

import "strings"

func ParseInline(s string) []TextNode {
	type delim struct {
		kind  string
		start int
	}

	var result []TextNode
	var stack []delim
	var textBuf strings.Builder

	flushText := func() {
		if textBuf.Len() > 0 {
			result = append(result, TextNode{Text: textBuf.String(), TextType: Text})
			textBuf.Reset()
		}
	}

	for pos := 0; pos < len(s); {
		if len(stack) > 0 {
			top := stack[len(stack)-1]
			if strings.HasPrefix(s[pos:], top.kind) {
				flushText()
				children := result[top.start:]
				result = result[:top.start]
				var wrapped TextNode
				switch top.kind {
				case "***":
					wrapped = TextNode{
						TextType: Bold,
						Children: []TextNode{
							{TextType: Italic, Children: children},
						},
					}
				case "**":
					wrapped = TextNode{TextType: Bold, Children: children}
				case "*", "_":
					wrapped = TextNode{TextType: Italic, Children: children}
				}
				result = append(result, wrapped)
				stack = stack[:len(stack)-1]
				pos += len(top.kind)
				continue
			}
		}

		if strings.HasPrefix(s[pos:], "`") {
			flushText()
			end := strings.Index(s[pos+1:], "`")
			if end < 0 {
				textBuf.WriteByte('`')
				pos++
			} else {
				content := s[pos+1 : pos+1+end]
				result = append(result, TextNode{Text: content, TextType: Code})
				pos += end + 2
			}
			continue
		}

		switch {
		case strings.HasPrefix(s[pos:], "***"):
			flushText()
			stack = append(stack, delim{kind: "***", start: len(result)})
			pos += 3
		case strings.HasPrefix(s[pos:], "**"):
			flushText()
			stack = append(stack, delim{kind: "**", start: len(result)})
			pos += 2
		case s[pos] == '*' || s[pos] == '_':
			flushText()
			stack = append(stack, delim{kind: string(s[pos]), start: len(result)})
			pos++
		default:
			textBuf.WriteByte(s[pos])
			pos++
		}
	}

	flushText()

	for _, d := range stack {
		result = append([]TextNode{{Text: d.kind, TextType: Text}}, result...)
	}

	return result
}
