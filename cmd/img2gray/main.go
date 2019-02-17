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
		err := img2gray.ToGray(*imgPath, *output)
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	err := img2gray.TraverseGray(*imgPath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Done!")
}
