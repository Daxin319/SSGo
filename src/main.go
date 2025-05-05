package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	path, err := os.Getwd()
	if err != nil {
		msg := fmt.Errorf("Error getting working directory: %v", err)
		fmt.Println(msg)
	}
	err = CopyStaticToPublic(path)
	if err != nil {
		msg := fmt.Errorf("Error copying files: %v", err)
		fmt.Println(msg)
	}
}

func CopyStaticToPublic(path string) error {
	pDir := path + "/public"
	_, err := os.Stat(pDir)
	if err == nil {
		err = os.RemoveAll(pDir)
	}
	sDir := path + "/static"
	dir := os.DirFS(sDir)

	err = os.CopyFS(pDir, dir)
	if err != nil {
		fmt.Println(err)
		return errors.New("Error copying files")
	}
	return nil
}
