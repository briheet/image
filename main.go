package main

import (
	"fmt"
	"os"

	"github.com/briheet/gopng/compress"
)

func main() {
	fmt.Println("starting...")

	compressedFile, err := compress.CompressTofile("Firefox_wallpaper.png", "compress.png")
	if err != nil {
		panic(err)
	}
	defer compressedFile.Close()

	os.Exit(0)
}
