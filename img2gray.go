package img2gray

import (
	"image"
	_ "image/jpeg"
	"image/png"
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

// ToGray generate gray image.
func ToGray(input, output string) error {
	f, err := os.Open(input)
	if err != nil {
		return err
	}
	defer f.Close()
	img, _, err := image.Decode(f)
	if err != nil {
		return err
	}

	// Convert image to grayscale
	grayImg := image.NewGray(img.Bounds())
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			grayImg.Set(x, y, img.At(x, y))
		}
	}
	o, err := os.Create(output)
	if err != nil {
		return err
	}
	defer f.Close()

	if err := png.Encode(o, grayImg); err != nil {
		return err
	}
	return nil
}
