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
		s := fmt.Sprintf(" %s=%s ", key, h.Props[key])
		finalString += s
	}
	return finalString
}

func (h *HTMLNode) Repr() string {
	return fmt.Sprintf("HTMLNode(%s, %s, %v, %v)", h.Tag, h.Value, h.Children, h.Props)
}
