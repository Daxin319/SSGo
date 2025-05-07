package main

import (
	"errors"
	"fmt"
	"os"
)

func CopyStaticToDocs(path string) error {
	pDir := path + "/docs"
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
