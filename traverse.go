package img2gray

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

// TraverseGray traverse files to generate gray images.
func TraverseGray(path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if isHiddenDirectoryOrFile(file) {
			continue
		}
		filePath := filepath.Join(path, file.Name())
		output := filepath.Join(path, "/gray_"+file.Name())
		if file.IsDir() {
			TraverseGray(filePath)
			continue
		}
		if isImage(filePath) {
			fmt.Printf("detect image file %v\n", filePath)
			ToGray(filePath, output)
		}
	}
}
