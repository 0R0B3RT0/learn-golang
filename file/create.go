package file

import (
	"os"
)

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
