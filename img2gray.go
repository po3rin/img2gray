package img2gray

import (
	"image"
	_ "image/jpeg"
	"image/png"
	"log"
	"os"
)

// ToGray generate gray image.
func ToGray(input, output string) {
	f, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	img, _, err := image.Decode(f)
	if err != nil {
		log.Fatal(err)
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
		log.Fatal(err)
	}
	defer f.Close()

	if err := png.Encode(o, grayImg); err != nil {
		log.Fatal(err)
	}
}
