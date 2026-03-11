package nodes

import "strings"

type Ref struct {
	URL   string
	Title string
}

var refDefs = map[string]Ref{}

func SetRef(label, url, title string) {
	refDefs[strings.ToLower(label)] = Ref{URL: url, Title: title}
}

func GetRef(label string) (string, string, bool) {
	ref, ok := refDefs[strings.ToLower(label)]
	return ref.URL, ref.Title, ok
}
