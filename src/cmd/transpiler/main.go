package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Daxin319/SSGo/src/fileio"
)

func main() {
	var basePath string
	if len(os.Args) > 1 && os.Args[1] != "serve" {
		basePath = os.Args[1]
	} else {
		basePath = ""
	}

	path, err := os.Getwd()
	if err != nil {
		fmt.Println("error getting working directory")
	}

	err = fileio.CopyStaticToDocs(path)
	if err != nil {
		msg := fmt.Errorf("error copying files: %v", err)
		fmt.Println(msg)
	}

	fileio.GeneratePagesRecursive(path+"/content", path+"/docs", path+"/template.html", basePath)

	if len(os.Args) > 1 && os.Args[1] == "serve" {
		http.Handle("/", http.FileServer(http.Dir("./docs")))
		err := http.ListenAndServe(":3000", nil)
		if err != nil {
			return
		}
	}

}
