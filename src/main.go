package main

import (
	"fmt"
	"main/src/nodes"
	"net/http"
)

func main() {
	node := &nodes.TextNode{
		Text:     "This is some anchor text",
		TextType: nodes.Link,
		Url:      "https://www.boot.dev",
	}
	fmt.Println(node.Repr())

	http.Handle("/", http.FileServer(http.Dir("./public")))
	http.ListenAndServe(":3000", nil)
}
