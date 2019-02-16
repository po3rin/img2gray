package img2gray

import (
	"image"
	"os"
	"strings"
)

// IsHiddenDirectoryOrFile check hidden directory or file.
func isHiddenDirectoryOrFile(f os.FileInfo) bool {
	return strings.HasPrefix(f.Name(), ".")
}

// IsImage check file is image.
func isImage(path string) bool {
	f, _ := os.Open(path)
	defer f.Close()
	_, _, err := image.DecodeConfig(f)
	if err != nil {
		return false
	}
	return true
}
