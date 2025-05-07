package nodes

import "fmt"

type HTMLNode struct {
	Tag      string
	Value    string
	Children []HTMLNode
	Props    map[string]string
}

func (h *HTMLNode) PropsToHTML() string {
	var finalString string
	for key := range h.Props {
		s := fmt.Sprintf(` %s="%s" `, key, h.Props[key])
		finalString += s
	}
	return finalString
}

func (h *HTMLNode) Repr() string {
	return fmt.Sprintf("HTMLNode(%s, %s, %v, %v)", h.Tag, h.Value, h.Children, h.Props)
}

func (h *HTMLNode) ToHTML() string {
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
