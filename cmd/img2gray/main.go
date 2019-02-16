package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/po3rin/img2gray"
)

var imgPath = flag.String("i", "", "path to the image")
var output = flag.String("o", "gray.png", "path to the output image")
var traverse = flag.Bool("t", false, "traverse direstory")

func main() {
	flag.Parse()
	if !*traverse {
		if *imgPath == "" {
			log.Fatal("i flag is required")
		}
		img2gray.ToGray(*imgPath, *output)
		return
	}
	img2gray.TraverseGray(*imgPath)
	fmt.Println("Done!")
}
