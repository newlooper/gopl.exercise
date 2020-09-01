package main

import (
	"fmt"
	"math"
)

const (
	width, height = 800, 600            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

type PointType int

const (
	middle PointType = 0 // not peak or trough
	peak   PointType = 1 // peak of wave
	trough PointType = 2 // trough of wave
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, pt := corner(i+1, j)
			bx, by, pt1 := corner(i, j)
			cx, cy, pt3 := corner(i, j+1)
			dx, dy, pt2 := corner(i+1, j+1)
			if math.IsNaN(ax) || math.IsNaN(ay) ||
				math.IsNaN(bx) || math.IsNaN(by) ||
				math.IsNaN(cx) || math.IsNaN(cy) ||
				math.IsNaN(dx) || math.IsNaN(dy) {
				continue
			}

			color := "grey"
			if pt == peak || pt1 == peak || pt2 == peak || pt3 == peak {
				color = "#f00"
			} else if pt == trough || pt1 == trough || pt2 == trough || pt3 == trough {
				color = "#00f"
			}
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' style='stroke: %s'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, color)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64, PointType) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z, pt := zPos(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, pt
}

func zPos(x, y float64) (float64, PointType) {
	d := math.Hypot(x, y) // distance from (0,0)
	pt := middle

	if math.Abs(d-math.Tan(d)) < 3 {
		pt = peak
		if 2*(math.Sin(d)-d*math.Cos(d))-d*d*math.Sin(d) > 0 {
			pt = trough
		}
	}
	return math.Sin(d) / d, pt
}
