package nodes

import "fmt"

type enum int

const (
	Text enum = iota
	Bold
	Italic
	Code
	Link
	Image
	Boldtalic
)

type TextNode struct {
	Text     string
	TextType enum
	Url      string
	Tag      string
	Value    string
	Children []TextNode
	Props    map[string]string
}

type HTMLNode interface {
	PropsToHTML() string
	ToHTML() string
}

func (h *TextNode) PropsToHTML() string {
	var finalString string
	for key := range h.Props {
		s := fmt.Sprintf(` %s="%s" `, key, h.Props[key])
		finalString += s
	}
	return finalString
}

func (h *TextNode) Repr() string {
	if h.Tag != "" {
		return fmt.Sprintf("HTMLNode(%s, %s, %v, %v)", h.Tag, h.Value, h.Children, h.Props)
	} else {
		return fmt.Sprintf("TextNode(%s, %s, %s)", h.Text, String(h.TextType), h.Url)
	}
}

func (h *TextNode) ToHTML() string {
	var cString string

	switch h.Tag {
	case "":
		return fmt.Sprintf("%s", h.Value)

	case "img":
		return fmt.Sprintf("<%s%s/>", h.Tag, h.PropsToHTML())

	default:
		if len(h.Children) == 0 {
			if len(h.Props) == 0 {
				return fmt.Sprintf("<%s>%s</%s>", h.Tag, h.Value, h.Tag)
			}
			return fmt.Sprintf("<%s%s>%s</%s>", h.Tag, h.PropsToHTML(), h.Value, h.Tag)
		}
		for _, child := range h.Children {
			cString += child.ToHTML()
		}
		if len(h.Props) == 0 {
			return fmt.Sprintf("<%s>%s</%s>", h.Tag, cString, h.Tag)
		}
		return fmt.Sprintf("<%s%s>%s</%s>", h.Tag, h.PropsToHTML(), cString, h.Tag)
	}
}

func String(t enum) string {
	switch t {
	case 0:
		return "text"
	case 1:
		return "bold"
	case 2:
		return "italic"
	case 3:
		return "code"
	case 4:
		return "link"
	case 5:
		return "image"
	case 6:
		return "boldtalic"
	default:
		return "unknown text type"
	}
}
