package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	var basePath string
	if len(os.Args) > 1 && os.Args[1] != "serve" {
		basePath = os.Args[1]
	} else {
		basePath = "/"
	}

	path, err := os.Getwd()
	if err != nil {
		fmt.Println("error getting working directory")
	}

	err = CopyStaticToDocs(path)
	if err != nil {
		msg := fmt.Errorf("Error copying files: %v", err)
		fmt.Println(msg)
	}

	GeneratePagesRecursive(path+"/content", path+"/docs", path+"/template.html", basePath+"/docs")

	if len(os.Args) > 1 && os.Args[1] == "serve" {
		http.Handle("/", http.FileServer(http.Dir("./docs")))
		http.ListenAndServe(":3000", nil)
	}

}
