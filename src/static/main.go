package main

import (
	"main/src/static/nodes"
	"net/http"
)

func main() {
	node := &nodes.TextNode{
		Text:     "This is some anchor text",
		TextType: nodes.Link,
		Url:      "https://www.boot.dev",
	}
	node.Repr()

	http.Handle("/", http.FileServer(http.Dir("./public")))
	http.ListenAndServe(":3000", nil)
}
