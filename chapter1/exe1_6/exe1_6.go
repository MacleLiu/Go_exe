// 通过在画板中添加更多颜色，然后通过有趣的方式改变SetColorIndex的第三个参数，修改莉萨如程序来产生多种色彩的图片。
// 本例中使用了随机数来改变图形颜色
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var palette = []color.Color{
	color.White,
	color.RGBA{0x00, 0xff, 0x00, 0xff},
	color.RGBA{0x00, 0xff, 0xff, 0xff},
	color.RGBA{0xff, 0xff, 0x00, 0xff},
	color.RGBA{0xff, 0x00, 0xff, 0xff},
	color.RGBA{0xcc, 0x00, 0x00, 0xff},
}

/*
	 const (
		blackIndex = 0
		colorIndex = 1

)
*/

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	if len(os.Args) > 1 && os.Args[1] == "web" {
		handler := func(w http.ResponseWriter, r *http.Request) {
			lissajous(w)
		}
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	} else {
		//lissajous(os.Stdout)  //Linux环境下

		//Windows环境下的两种方法
		//来自http://centphp.com/view/335
		f, err := os.Create("a.gif")
		if err != nil {
			panic(err)
		}
		defer f.Close()
		lissajous(f)

		/* buf := &bytes.Buffer{}
		lissajous(buf)
		if err := ioutil.WriteFile("a.gif", buf.Bytes(), 0666); err != nil {
			panic(err)
		} */
	}
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	colorIndex := uint8(rand.Intn(5) + 1) //使用随机数来改变图形颜色
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), colorIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
