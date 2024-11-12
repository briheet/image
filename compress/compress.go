package compress

import (
	"image"
	"image/png"
	"os"

	pngquant "github.com/yusukebe/go-pngquant"
)

func CompressToFile(got string, name string) (*os.File, error) {
	given, err := os.Open(got)
	if err != nil {
		return nil, err
	}
	defer given.Close()

	img, _, err := image.Decode(given)
	if err != nil {
		return nil, err
	}

	compressedImg, err := pngquant.Compress(img, "1")
	if err != nil {
		return nil, err
	}

	compressedFile, err := os.Create(name)
	if err != nil {
		return nil, err
	}
	defer compressedFile.Close()

	err = png.Encode(compressedFile, compressedImg)
	if err != nil {
		return nil, err
	}

	return compressedFile, nil
}
