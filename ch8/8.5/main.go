package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"os"
	"sync"
	"time"
)

const (
	workers                = 128 // 视情况调整
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 1024, 1024
)

var rows = make(chan int, height) // 多个 goroutines 从此队列中取

func init() {
	for row := 0; row < height; row++ {
		rows <- row
	}
	close(rows) // 只能从队列中取，不能再入
}

func main() {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	wg := sync.WaitGroup{}

	start := time.Now()

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			for py := range rows { // 无序的争抢
				y := float64(py)/height*(ymax-ymin) + ymin
				for px := 0; px < width; px++ {
					x := float64(px)/width*(xmax-xmin) + xmin
					// Image point (px, py) represents complex value z.
					z := complex(x, y)
					img.Set(px, py, mandelbrot(z))
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()
	png.Encode(os.Stdout, img) // NOTE: ignoring errors

	log.Fatalf("%f", time.Since(start).Seconds()) // 日志发到 stderr，避免掺入到图片中
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

/*
并发程序用时大约： 0.077999s

非并发的 gopl.io/ch3/mandelbrot/main.go 用时大约： 0.204000s
 */
