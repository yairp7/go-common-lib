package graphics

import (
	"image"
	"image/color"
	"image/png"
	"os"

	ds "github.com/yairp7/go-common-lib/ds/basic"
)

func ReadPNG(path string) (image.Image, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	img, err := png.Decode(f)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func SavePNG(src image.Image, path string) error {
	fOut, err := os.Create(path)
	if err != nil {
		return err
	}
	defer fOut.Close()

	return png.Encode(fOut, src)
}

func SearchPixels(
	src image.Image,
	isValid func(point image.Point, clr color.Color) bool,
) []image.Point {
	offsets := []image.Point{
		{X: -1, Y: -1},
		{X: 0, Y: -1},
		{X: 1, Y: -1},
		{X: 1, Y: 0},
		{X: 1, Y: 1},
		{X: 0, Y: 1},
		{X: -1, Y: 1},
		{X: -1, Y: 0},
	}

	rect := src.Bounds()
	stack := make([]image.Point, 0)
	stack = append(stack, image.Point{X: 0, Y: 0})
	visited := ds.NewSet[image.Point]()

	pixelsFound := make([]image.Point, 0)

	for len(stack) > 0 {
		pixel := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]

		visited.Add(pixel)

		if isValid(pixel, src.At(pixel.X, pixel.Y)) {
			pixelsFound = append(pixelsFound, pixel)
		}

		for _, offset := range offsets {
			nextPixel := pixel.Add(offset)
			if visited.Has(nextPixel) {
				continue
			}
			if nextPixel.X < rect.Min.X || nextPixel.X > rect.Max.X {
				continue
			}
			if nextPixel.Y < rect.Min.Y || nextPixel.Y > rect.Max.Y {
				continue
			}
			stack = append(stack, nextPixel)
		}
	}

	return pixelsFound
}
