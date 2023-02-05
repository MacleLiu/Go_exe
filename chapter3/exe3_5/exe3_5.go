// mandelbrot函数生成一个PNG格式的Mandelbrot分形图
// 用image.NewRGBA函数和color.RGBA类型或color.YCbCr类型实现一个Mandelbrot集的全彩图
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}
	//png.Encode(os.Stdout, img)
	f, err := os.Create("a.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	png.Encode(f, img)
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			//return color.Gray{255 - contrast*n}
			//return color.RGBA{255, 255 - contrast*n, 255, 255}
			return color.YCbCr{255, 255, 255 - contrast*n}
		}
	}
	return color.Black
}
