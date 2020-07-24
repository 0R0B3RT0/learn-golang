package main

import (
	"fmt"
	"os"
)

func main() {
	input := os.Args[1:]
	fmt.Printf("%v\n", CreateFile(input[0], input[1:]...))
}

func CreateFile(baseDir string, names ...string) []string {
	var files = []string{}
	for _, name := range names {
		var fileName = baseDir + "/" + name
		file, _ := os.Create(fileName)
		defer file.Close()

		files = append(files, file.Name())
	}

	return files
}
