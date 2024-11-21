package grayscale

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func GrascaleTheFile(got string, name string) (*os.File, error) {
	given, err := os.Open(got)
	if err != nil {
		return nil, err
	}

	defer given.Close()

	img, _, err := image.Decode(given)
	if err != nil {
		return nil, err
	}

	size := img.Bounds().Size()
	rect := image.Rect(0, 0, size.X, size.Y)
	newImage := image.NewRGBA(rect)

	for i := 0; i < size.X; i++ {
		for j := 0; j < size.Y; j++ {
			r, g, b, a := img.At(i, j).RGBA()
			average := (r + g + b) / 3

			newImage.Set(i, j, color.NRGBA64{R: uint16(average), G: uint16(average), B: uint16(average), A: uint16(a)})
		}
	}

	resultFile, err := os.Create(name)
	if err != nil {
		return nil, err
	}

	err = png.Encode(resultFile, newImage)
	if err != nil {
		return nil, err
	}

	return resultFile, nil
}
