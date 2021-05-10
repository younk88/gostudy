package main

import (
	"fmt"
	"image"
	"image/color"
	"os"
	"path/filepath"

	"github.com/disintegration/imageorient"
	"golang.org/x/tour/pic"
)


type Image struct {
	r int
	g int
	b int
}

func (m Image)ColorModel() color.Model {
	return color.RGBAModel
}

func (m Image)Bounds() image.Rectangle {
	return image.Rect(0, 0, 64, 64)
}

func (m Image)At(x, y int) color.Color {
	//del := x + y / 2
	//return color.RGBA{(uint8)(m.r - del), (uint8)(m.g - del), (uint8)(m.b - del), 1}
	return color.RGBA{uint8(m.r), uint8(m.g), uint8(m.b), 1}
}

func main() {
	m := Image{88, 44, 22}
	fmt.Println(m)
	pic.ShowImage(m)

	fName := os.Args[1]
	f, _ := os.Open(fName)
	info, s, err := imageorient.DecodeConfig(f)
	fmt.Printf("info: %v, format: %v, err:%v\n", info, s, err)
	fmt.Println(filepath.Ext(fName))
}