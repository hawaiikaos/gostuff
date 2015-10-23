package main

import (
	"golang.org/x/tour/pic"
	"image"
	//"fmt"
	"image/color"
	)

type Image struct{
	x, y int
	c uint8
}

func (i *Image) ColorModel() color.Model {
	return color.RGBAModel
}


func (i *Image) At(x, y int) color.Color {
	return color.RGBA{i.c+uint8(x), i.c+uint8(y), 255, 255}
}

func (i *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.x, i.y)
}

func main() {
	m := Image{100,200,200}
	pic.ShowImage(&m)
}