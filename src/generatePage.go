package main

import (
	"fmt"
	"main/src/nodes"
	"os"
	"strings"
)

func GeneratePage(fromPath, destPath, templatePath, basePath string) {
	fmt.Printf("Generating page from %s to %s using template at %s\n", fromPath, destPath, templatePath)

	src, err := os.Open(fromPath)
	if err != nil {
		fmt.Println("error opening fromPath file")
	}
	readSrc, err := os.ReadFile(src.Name())
	if err != nil {
		fmt.Println("error reading data from fromPath")
	}
	temp, err := os.Open(templatePath)
	if err != nil {
		fmt.Println("error opening templatePath file")
	}
	readTemp, err := os.ReadFile(temp.Name())
	if err != nil {
		fmt.Println("error reading data from templatePath")
	}

	title, content, err := ExtractTitle(string(readSrc))

	var cString string
	node := nodes.MarkdownToHTMLNode(content)

	cString += node.ToHTML()

	titleTemp := strings.Replace(string(readTemp), "{{ Title }}", title, 1)
	contentTemp := strings.Replace(string(titleTemp), "{{ Content }}", cString, 1)
	hrefTemp := strings.ReplaceAll(string(contentTemp), `href="/`, `href="`+basePath+"/")
	finalTemp := strings.ReplaceAll(string(hrefTemp), `src=/`, `src=`+basePath+"/")

	os.WriteFile(destPath, []byte(finalTemp), 0755)
}
