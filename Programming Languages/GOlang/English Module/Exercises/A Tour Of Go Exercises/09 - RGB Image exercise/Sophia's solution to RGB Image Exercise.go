package main

/*

	This code does not work in an IDE

import "golang.org/x/tour/pic"
import "image"
import "image/color"

type Image struct {
	Width, Height int
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.Width, i.Height)
}

func (i Image) At(x int, y int) color.Color {
	return color.RGBA{uint8(((x * x) * (y * y)) / 4), 0, 0, 255}
}

func main() {
	m := Image{500, 500}
	pic.ShowImage(m)
}
*/
