package main

import (
	"fmt"
	"os"

	"github.com/briheet/gopng/compress"
)

func main() {
	fmt.Println("starting...")

	compressedFile, err := compress.CompressToFile("Firefox_wallpaper.png", "compress.png")
	if err != nil {
		panic(err)
	}
	defer compressedFile.Close()

	fmt.Println("...ending")

	os.Exit(0)
}
