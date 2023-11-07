package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func pic(dx, dy int) [][]color.RGBA {
	pixels := make([][]color.RGBA, dy)
	for y := range pixels {
		pixels[y] = make([]color.RGBA, dx)
		for x := range pixels[y] {
			v := uint8((x + y) / 6)
			pixels[y][x] = color.RGBA{v, 13, 22, 255}
		}
	}
	return pixels
}

func Pic(dx, dy int) image.Image {
	pixels := pic(dx, dy)
	img := image.NewRGBA(image.Rect(0, 0, dx, dy))
	for y := range pixels {
		for x, col := range pixels[y] {
			img.Set(x, y, col)
		}
	}
	return img
}

func saveAsPNG(img image.Image, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	return png.Encode(file, img)
}

func main() {
	img := Pic(256, 256)
	saveAsPNG(img, "output.png")
}
