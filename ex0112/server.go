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
	"net/url"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	p := params{
		cycles:  5,
		res:     0.001,
		size:    100,
		nframes: 64,
		delay:   8,
	}
	if q, err := url.ParseQuery(r.URL.RawQuery); err == nil {
		if cycles, err := strconv.Atoi(q.Get("cycles")); err == nil {
			p.cycles = cycles
		}
		if res, err := strconv.ParseFloat(q.Get("res"), 64); err == nil {
			p.res = res
		}
		if size, err := strconv.Atoi(q.Get("size")); err == nil {
			p.size = size
		}
		if nframes, err := strconv.Atoi(q.Get("nframes")); err == nil {
			p.nframes = nframes
		}
		if delay, err := strconv.Atoi(q.Get("delay")); err == nil {
			p.delay = delay
		}
	}

	lissajous(w, p)
}

var palette = []color.Color{
	color.Black,
	color.RGBA{0x00, 0xff, 0x00, 0xff}, // green
}

const (
	blackIndex = iota
	greenIndex
)

type params struct {
	cycles  int     // 発振器xが完了する周回の回数
	res     float64 // 回転の分解能
	size    int     // 画像キャンバスは[-size..+size]の範囲を扱う
	nframes int     // アニメーションフレーム数
	delay   int     // 10ms単位でのフレーム間の遅延
}

func lissajous(out io.Writer, p params) {
	freq := rand.Float64() * 3.0 // 発振器yの相対周波数
	anim := gif.GIF{LoopCount: p.nframes}
	phase := 0.0 // 位相差
	for i := 0; i < p.nframes; i++ {
		rect := image.Rect(0, 0, 2*p.size+1, 2*p.size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(p.cycles)*2*math.Pi; t += p.res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(p.size+int(x*float64(p.size)+0.5), p.size+int(y*float64(p.size)+0.5), greenIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, p.delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // 注意：エンコードエラーを無視
}
