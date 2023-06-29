package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct {
	width, height int
	color         uint8
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.width, i.height)
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) At(x, y int) color.Color {
	//return color.RGBA{i.color + uint8(x), i.color + uint8(y), 255, 255}
	//return color.RGBA{(uint8(x) + uint8(y)) / 2, (uint8(x) + uint8(y)) / 2, 255, 255}
	v := uint8(x * y)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{256, 256, 100}
	pic.ShowImage(m)
}
