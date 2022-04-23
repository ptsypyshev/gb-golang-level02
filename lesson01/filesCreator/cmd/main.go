package main

import (
	"flag"
	"fmt"
	"github.com/ptsypyshev/gb-golang-level02/lesson01/filesCreator"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	var dirPath string
	var numFiles int
	flag.StringVar(&dirPath, "dir", "", "Specify dir to create files.")
	flag.IntVar(&numFiles, "number", 100, "Specify number of new files.")
	flag.Parse()

	if dirPath == "" {
		curPath, err := os.Getwd()
		if err != nil {
			fmt.Println("Cannot get current work directory")
			os.Exit(1)
		}
		dirPath = filepath.Join(curPath, "files")
		if _, err := os.Stat(dirPath); os.IsNotExist(err) {
			if mkdirError := os.Mkdir(dirPath, 0755); mkdirError != nil {
				fmt.Println("Cannot make directory for new files")
				os.Exit(1)
			}
		}
	}

	fmt.Printf("Try to create %d files in %s\n", numFiles, dirPath)

	for i := 0; i < numFiles; i++ {
		strStep := strconv.Itoa(i)
		if err := filesCreator.CreateFile(dirPath, "test"+strStep+".txt", strStep); err != nil {
			fmt.Println(err.Error())
		}
	}

}
