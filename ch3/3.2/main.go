package main

import (
	"fmt"
	"io"
	"math"
	"os"
)

const (
	width, height = 800, 600            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

type zPos func(x, y float64) float64

func svg(w io.Writer, f zPos) {
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j, f)
			bx, by := corner(i, j, f)
			cx, cy := corner(i, j+1, f)
			dx, dy := corner(i+1, j+1, f)
			if math.IsNaN(ax) || math.IsNaN(ay) ||
				math.IsNaN(bx) || math.IsNaN(by) ||
				math.IsNaN(cx) || math.IsNaN(cy) ||
				math.IsNaN(dx) || math.IsNaN(dy) {
				continue
			}
			fmt.Fprintf(w, "<polygon style='stroke: %s; fill: #FFFFFF' points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				"#888888", ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintln(w, "</svg>")
}

func corner(i, j int, f zPos) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func moguls(x, y float64) float64 {
	r := (x * x + y * y + 15 * (math.Sin(x) * math.Sin(x) + math.Sin(y) * math.Sin(y))) * 0.001
	return math.Max(r, 0.0)
}

func eggbox(x, y float64) float64 {
	return math.Pow(2, math.Sin(x)) * math.Pow(2, math.Sin(y)) / 12
}

func saddle(x, y float64) float64 {
	a := 25.0
	b := 20.0
	a2 := a * a
	b2 := b * b
	return (y*y/a2 - x*x/b2)
}

func main() {
	usage := "Usage: eggbox|moguls|saddle"
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(1)
	}
	var f zPos
	switch os.Args[1] {
	case "eggbox":
		f = eggbox
	case "moguls":
		f = moguls
	case "saddle":
		f = saddle
	default:
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(1)
	}
	svg(os.Stdout, f)
}
