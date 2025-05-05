package main

import (
	"fmt"
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
}
