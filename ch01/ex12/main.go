// リサジュー図形のサーバを修正して、URL からパラメータ値を読み取るようにしなさい。
// たとえば、http://localhost:8000/?cycles=20 のようなURLでは、周回の回数をデフォルトの5ではなく20に設定するようにしなさい。
// 文字列パラメータを整数へ変換するためにstrconv.Atoi 関数を使いなさい。
// その変換関数のドキュメントは go doc strconv.Atoi で見ることができます。

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
	"strconv"
	"time"
)

var palette = []color.Color{
	color.RGBA{0x00, 0x00, 0x00, 0xff},
	color.RGBA{0x00, 0xff, 0x00, 0xff},
	color.RGBA{0xff, 0x00, 0x00, 0xff},
	color.RGBA{0x00, 0x00, 0xff, 0xff}}

const (
	backgroundIndex = 0 // パレットの最初の色
	lineColorNum    = 3 // 線の色数
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		p := lissajousParams{}
		p.Cycles, _ = strconv.Atoi(r.URL.Query().Get("cycles"))
		p.Res, _ = strconv.ParseFloat(r.URL.Query().Get("res"), 64)
		p.Size, _ = strconv.Atoi(r.URL.Query().Get("size"))
		p.Nframes, _ = strconv.Atoi(r.URL.Query().Get("nframes"))
		p.Delay, _ = strconv.Atoi(r.URL.Query().Get("delay"))

		lissajous(w, p)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type lissajousParams struct {
	Cycles  int
	Res     float64
	Size    int
	Nframes int
	Delay   int
}

func lissajous(out io.Writer, p lissajousParams) {
	if p.Cycles == 0 {
		p.Cycles = 5 // 発振器 x が完了する
	}

	if p.Res == 0 {
		p.Res = 0.001 // 回転の分解能
	}

	if p.Size == 0 {
		p.Size = 100 // 画像キャンバスは [-size..+size] の範囲を扱う
	}

	if p.Nframes == 0 {
		p.Nframes = 64 // アニメーションフレーム数
	}

	if p.Delay == 0 {
		p.Delay = 8 // 10ms 単位でのフレーム間の遅延
	}

	freq := rand.Float64() * 3.0 // 発振器 y の相対周波数
	anim := gif.GIF{LoopCount: p.Nframes}
	phase := 0.0 // 位相差

	for i := 0; i < p.Nframes; i++ {
		rect := image.Rect(0, 0, 2*p.Size+1, 2*p.Size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(p.Cycles)*2*math.Pi; t += p.Res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(p.Size+int(x*float64(p.Size)+0.5),
				p.Size+int(y*float64(p.Size)+0.5),
				uint8(rand.Intn(2))+1)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, p.Delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // 注意: エンコードエラーを無視
}
