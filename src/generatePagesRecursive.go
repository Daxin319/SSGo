package main

import (
	"fmt"
	"os"
	"strings"
)

func generatePagesRecursive(fromDirPath, destDirPath, templatePath string) {
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
			generatePagesRecursive(fromDirPath+"/"+entry.Name(), destDirPath+"/"+entry.Name(), templatePath)
		} else {
			generatePage(fromDirPath+"/"+entry.Name(), destDirPath+"/"+strings.TrimRight(entry.Name(), ".md")+".html", templatePath)
		}
	}
}
