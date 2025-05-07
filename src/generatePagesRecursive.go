package main

import (
	"fmt"
	"os"
	"strings"
)

func GeneratePagesRecursive(fromDirPath, destDirPath, templatePath, basePath string) {
	dir, err := os.ReadDir(fromDirPath)
	if err != nil {
		fmt.Println("error reading content directory")
	}

	for _, entry := range dir {
		if entry.IsDir() {
			err = os.Mkdir(destDirPath+"/"+entry.Name(), 0755)
			if err != nil {
				fmt.Println("error creating public directory")
			}
			GeneratePagesRecursive(fromDirPath+"/"+entry.Name(), destDirPath+"/"+entry.Name(), templatePath, basePath)
		} else {
			GeneratePage(fromDirPath+"/"+entry.Name(), destDirPath+"/"+strings.TrimRight(entry.Name(), ".md")+".html", templatePath, basePath)
		}
	}
}
