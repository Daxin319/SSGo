package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	path, err := os.Getwd()
	if err != nil {
		msg := fmt.Errorf("Error getting working directory: %v", err)
		fmt.Println(msg)
	}
	err = copyStaticToPublic(path)
	if err != nil {
		msg := fmt.Errorf("Error copying files: %v", err)
		fmt.Println(msg)
	}

	generatePage(path+"/content/index.md", path+"/public/index.html", path+"/template.html")

	http.Handle("/", http.FileServer(http.Dir("./public")))
	http.ListenAndServe(":3000", nil)

}
