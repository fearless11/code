//server1 is a minimal "echo" server.
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

// 调色板
var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

type A struct{}

func main() {

	var a *A
	var aa A
	if a == nil {
		fmt.Println("a is nil")
	}
	fmt.Printf("%T\n%T\n", a, aa)

	str := "江欢vera"
	for i, v := range str {
		n := strconv.FormatInt(int64(v), 2)
		fmt.Println(i, n)
	}
	fmt.Println(len(str), []byte(str))

	fmt.Println(fmt.Sprintf("%q", []string{"aa", "bb"}) == `["aa" "bb"]`)

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the path component of the request URL r.
func handler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "URL.Path= %q\n", r.URL.Path)

	if err := r.ParseForm(); err != nil {
		log.Fatal(err)
	}
	for k, v := range r.Form {
		if k == "cycles" {
			cycles, err := strconv.Atoi(v[0])
			if err != nil {
				log.Fatal(err)
			}
			lissajous(w, cycles)
		}
	}

}

func lissajous(out io.Writer, cycles int) {
	const (
		// cycles  = 5
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
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
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
