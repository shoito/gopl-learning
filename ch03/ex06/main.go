// スーパーサンプリング(supersampling)は、個々の画素内の複数の点のカラー値を計算して平均を求めることでピクセル化の影響を薄める技法です。
// 最も単純な方法は、個々の画素を四つの「サブピクセル」へ分割することです。その方法を実装しなさい。

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
	for py := 0; py < height; py +=2 {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px +=2 {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			c := mandelbrot(z)
			// Image point (px, py) represents complex value z.
			//img.Set(px, py, c)
			img.Set(px, py, c)
			img.Set(px+1, py, c)
			img.Set(px, py+1, c)
			img.Set(px+1, py+1, c)

		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.RGBA{255 - contrast*n, 0, 0, 255}
		}
	}
	return color.Black
}