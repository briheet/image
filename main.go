package main

import (
	"fmt"
	"os"

	"github.com/briheet/gopng/compress"
	"github.com/briheet/gopng/grayscale"
)

func compressTheGivenFile(given string, got string) os.File {
	fmt.Println("wait babygurl, its compressing :)")
	compressedFile, err := compress.CompressToFile(given, got)
	if err != nil {
		panic(err)
	}
	defer compressedFile.Close()

	return *compressedFile
}

func convertTograyScale(given string, got string) os.File {
	grayscaleFile, err := grayscale.GrascaleTheFile(given, got)
	if err != nil {
		panic(err)
	}
	defer grayscaleFile.Close()

	return *grayscaleFile
}

func main() {
	fmt.Println("starting...")
	fmt.Println("Pass in the arguments, enter -h for help")

	arguments := os.Args

	if len(arguments) < 2 {
		fmt.Println("Error: Insufficient arguments.")
		fmt.Println("Usage:")
		fmt.Println("  -h for help")
		os.Exit(-1)
	}

	switch arguments[1] {
	case "-h":
		fmt.Println("Help:")
		fmt.Println("Usage:")
		fmt.Println("  -compress  GivenFileName OutputFileName")
		fmt.Println("  -greyscale GivenFileName OutputFileName")
		os.Exit(0)

	case "-compress":
		if len(arguments) != 4 {
			fmt.Println("Error: Incorrect number of arguments for compression.")
			fmt.Println("Usage: -compress GivenFileName OutputFileName")
			os.Exit(-1)
		}
		compressedFile := compressTheGivenFile(arguments[2], arguments[3])
		defer compressedFile.Close()

	case "-greyscale":
		if len(arguments) != 4 {
			fmt.Println("Error: Incorrect number of arguments for compression.")
			fmt.Println("Usage: -greyscale GivenFileName OutputFileName")
			os.Exit(-1)
		}
		grayscaleFile := convertTograyScale(arguments[2], arguments[3])
		defer grayscaleFile.Close()

	default:
		fmt.Println("Error: Unknown command.")
		fmt.Println("Usage:")
		fmt.Println("  -h for help")
		os.Exit(-1)
	}

	fmt.Println("...ending")
	os.Exit(0)
}
