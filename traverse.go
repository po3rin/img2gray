package img2gray

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

// TraverseGray traverse files to generate gray images.
func TraverseGray(path string) error {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
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
			err := ToGray(filePath, output)
			if err != nil {
				return err
			}
			fmt.Printf("generate grayscale image file in %v\n", output)
		}
	}
	return nil
}
