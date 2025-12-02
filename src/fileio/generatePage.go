package fileio

import (
	"fmt"
	"os"
	"strings"

	"github.com/Daxin319/SSGo/src/blocks"
	"github.com/Daxin319/SSGo/src/renderer/html"
)

func generatePage(fromPath, destPath, templatePath, basePath string) {
	fmt.Printf("Generating page from %s to %s using template at %s\n", fromPath, destPath, templatePath)

	src, err := os.Open(fromPath)
	if err != nil {
		fmt.Println("error opening fromPath file")
	}
	readSrc, err := os.ReadFile(src.Name())
	if err != nil {
		fmt.Println("error reading data from fromPath")
	}
	fmt.Printf("%q\n", readSrc)
	temp, err := os.Open(templatePath)
	if err != nil {
		fmt.Println("error opening templatePath file")
	}
	readTemp, err := os.ReadFile(temp.Name())
	if err != nil {
		fmt.Println("error reading data from templatePath")
	}

	title, content, err := blocks.ExtractTitle(string(readSrc))

	var cString string
	node := html.MarkdownToHTMLNode(content)

	cString += node.ToHTML()

	titleTemp := strings.Replace(string(readTemp), "{{ Title }}", title, 1)
	contentTemp := strings.Replace(titleTemp, "{{ Content }}", cString, 1)
	hrefTemp := strings.ReplaceAll(contentTemp, `href="/`, `href="`+basePath+"/")
	srcTemp := strings.ReplaceAll(hrefTemp, `src="/`, `src="`+basePath+"/")
	finalTemp := strings.Replace(srcTemp, `docs/index.css`, `index.css`, 1)

	err = os.WriteFile(destPath, []byte(finalTemp), 0755)
	if err != nil {
		return
	}

	return
}
