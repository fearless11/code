// Lissajous 利萨如 generates GIF animations of random Lissajous figures
/*
GIF
(Graphics Interchange Format)的原义是“图像互换格式”，
是CompuServe公司在1987年开发的图像文件格式.
GIF文件的数据，是一种基于LZW算法的连续色调的无损压缩格式。
其压缩率一般在50%左右，它不属于任何应用程序。
GIF格式的另一个特点是其在一个GIF文件中可以存多幅彩色图像，
如果把存于一个文件中的多幅图像数据逐幅读出并显示到屏幕上，
就可构成一种最简单的动画。
*/
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

// 调色板
var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

func Lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	Lissajous(os.Stdout)
}
