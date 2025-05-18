package nodes

import "fmt"

type enum int

const (
	Text enum = iota // WHAT THE FUCK IS AN ENUM *Bald eagle noises*
	Bold
	Italic
	Code
	Link
	Image
	BoldItalic
	Strikethrough
	Subscript
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
		s := fmt.Sprintf(` %s="%s"`, key, h.Props[key]) // convert props to html with appropriate whitespace
		finalString += s
	}
	return finalString
}

func (h *TextNode) Repr() string { // this was honestly only here for testing, fuck it
	if h.Tag != "" {
		return fmt.Sprintf("HTMLNode(%s, %s, %v, %v)", h.Tag, h.Value, h.Children, h.Props)
	}
	return fmt.Sprintf("TextNode(%s, %s, %s)", h.Text, String(h.TextType), h.Url)
}

func (h *TextNode) ToHTML() string { // convert node to html string
	var cString string
	switch h.Tag { // switch on tag
	case "": // no tag means it's likely plaintext or a wrapper node
		if len(h.Children) > 0 {
			for _, c := range h.Children { // if there are child nodes convert them to HTML and return the final string
				cString += c.ToHTML()
			}
			return cString
		}
		return h.Text
	case "img": // images are a special case
		return fmt.Sprintf("<%s%s/>", h.Tag, h.PropsToHTML())
	default:
		if len(h.Children) == 0 { // otherwise, if there are no children check for properties, if there are no properties return as is
			if len(h.Props) == 0 {
				return fmt.Sprintf("<%s>%s</%s>", h.Tag, h.Text, h.Tag)
			}
			return fmt.Sprintf("<%s%s>%s</%s>", h.Tag, h.PropsToHTML(), h.Text, h.Tag)
		}
		for _, c := range h.Children { // if there are children convert to html and append to string
			cString += c.ToHTML()
		}
		if len(h.Props) == 0 {
			return fmt.Sprintf("<%s>%s</%s>", h.Tag, cString, h.Tag) // if there are no properties then append in standard structure
		}
		return fmt.Sprintf("<%s%s>%s</%s>", h.Tag, h.PropsToHTML(), cString, h.Tag) // if there are properties convert them to html and return formatted string
	}
}

func String(t enum) string {
	switch t {
	case Text:
		return "text"
	case Bold:
		return "bold"
	case Italic:
		return "italic"
	case Strikethrough:
		return "strikethrough"
	case Subscript:
		return "subscript"
	case Code:
		return "code"
	case Link:
		return "link"
	case Image:
		return "image"
	case BoldItalic:
		return "bolditalic"
	default:
		return "unknown text type"
	}
}
