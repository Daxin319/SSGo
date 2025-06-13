package nodes

import (
	"fmt"
	"strings"
)

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
	Superscript
	Highlight
	RawHTML
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

func escapeHTML(s string) string {
	s = strings.ReplaceAll(s, "&", "&amp;")
	s = strings.ReplaceAll(s, "<", "&lt;")
	s = strings.ReplaceAll(s, ">", "&gt;")
	return s
}

func escapeCodeHTML(s string) string {
	s = strings.ReplaceAll(s, "&", "&amp;")
	s = strings.ReplaceAll(s, "<", "&lt;")
	s = strings.ReplaceAll(s, ">", "&gt;")
	return s
}

func (h *TextNode) PropsToHTML() string {
	var finalString string
	for key := range h.Props {
		s := fmt.Sprintf(` %s="%s"`, key, h.Props[key])
		finalString += s
	}
	return finalString
}

func (h *TextNode) Repr() string {
	if h.Tag != "" {
		return fmt.Sprintf("HTMLNode(%s, %s, %v, %v)", h.Tag, h.Value, h.Children, h.Props)
	}
	return fmt.Sprintf("TextNode(%s, %s, %s)", h.Text, String(h.TextType), h.Url)
}

func (h *TextNode) ToHTML() string {
	if h.TextType == RawHTML {
		return h.Text
	}
	var cString string
	switch h.Tag {
	case "":
		if len(h.Children) > 0 {
			for _, c := range h.Children {
				cString += c.ToHTML()
			}
			return cString
		}
		return h.Text

	case "img":
		return fmt.Sprintf("<%s%s/>", h.Tag, UnescapeString(h.PropsToHTML()))

	case "hr":
		return "<hr />"

	case "code":
		if len(h.Children) == 0 {
			return fmt.Sprintf("<%s>%s</%s>", h.Tag, escapeCodeHTML(h.Text), h.Tag)
		}
		for _, c := range h.Children {
			if c.Text != "" && len(c.Children) == 0 {
				cString += escapeCodeHTML(c.Text)
			} else {
				cString += c.ToHTML()
			}
		}
		return fmt.Sprintf("<%s>%s</%s>", h.Tag, cString, h.Tag)

	default:
		if len(h.Children) == 0 {
			if len(h.Props) == 0 {
				return fmt.Sprintf("<%s>%s</%s>", h.Tag, escapeHTML(h.Text), h.Tag)
			}
			return fmt.Sprintf("<%s%s>%s</%s>", h.Tag, h.PropsToHTML(), escapeHTML(h.Text), h.Tag)
		}
		for _, c := range h.Children {
			cString += c.ToHTML()
		}
		if len(h.Props) == 0 {
			return fmt.Sprintf("<%s>%s</%s>", h.Tag, cString, h.Tag)
		}
		return fmt.Sprintf("<%s%s>%s</%s>", h.Tag, h.PropsToHTML(), cString, h.Tag)
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
	case Superscript:
		return "superscript"
	case Code:
		return "code"
	case Link:
		return "link"
	case Image:
		return "image"
	case BoldItalic:
		return "bolditalic"
	case Highlight:
		return "highlight"
	case RawHTML:
		return "rawhtml"
	default:
		return "unknown text type"
	}
}
