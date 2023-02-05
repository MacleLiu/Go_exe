// 运用牛顿法求z^4-1=0的复数解，进行分形吗
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
			img.Set(px, py, newtonMethod(z))
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

func newtonMethod(z complex128) color.Color {
	var x, f, f1 complex128
	var n uint8

	for cmplx.Abs(x-z) >= 1e-5 {
		x = z
		f = cmplx.Pow(x, 4) - 1
		f1 = 4 * cmplx.Pow(x, 3)
		z = x - f/f1
		//统计迭代次数
		if n < 255 {
			n++
		}
	}
	//根据迭代次数设置灰度
	//return color.Gray{255 - n}

	//上色
	if cmplx.Abs(z-1) <= 1e-5 {
		//逼近1，红色
		return color.RGBA{255, 0, 0, 255}
	} else if cmplx.Abs(z+1) <= 1e-5 {
		//逼近-1，绿色
		return color.RGBA{0, 255, 0, 255}
	} else if cmplx.Abs(z-1i) <= 1e-5 {
		//逼近0+1i，蓝色
		return color.RGBA{0, 0, 255, 255}
	} else {
		//逼近0-1i，黄色
		return color.RGBA{255, 255, 0, 255}
	}
}
