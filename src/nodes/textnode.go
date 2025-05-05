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
)

type TextNode struct {
	Text     string
	TextType enum
	Url      string
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
	default:
		return "unknown text type"
	}
}

func (t *TextNode) EqN(a *TextNode) bool {
	if t.Text == a.Text && t.TextType == a.TextType && t.Url == a.Url {
		return true
	}
	return false
}

func (t *TextNode) Repr() string {
	return fmt.Sprintf("TextNode(%s, %s, %s)", t.Text, String(t.TextType), t.Url)
}
