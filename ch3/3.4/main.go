package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
)

const (
	cells   = 100         // number of grid cells
	xyrange = 30.0        // axis ranges (-xyrange..+xyrange)
	angle   = math.Pi / 6 // angle of x, y axes (=30°)
)

var (
	width, height float64 = 800, 600
	xyscale               = width / 2 / xyrange // pixels per x or y unit
	zscale                = height * 0.4        // pixels per z unit
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(writer http.ResponseWriter, request *http.Request) {

	if err := request.ParseForm(); err != nil {
		log.Printf("parseForm: %v\n", err)
	}

	for k, v := range request.Form {
		var err error
		if k == "width" {
			if width, err = strconv.ParseFloat(v[0], 64); err != nil {
				width = 800
				log.Printf("parseWidth: %v\n", err)
			}
		}
		if k == "height" {
			if height, err = strconv.ParseFloat(v[0], 64); err != nil {
				height = 600
				log.Printf("parseHeight: %v\n", err)
			}
		}
	}

	xyscale = width / 2 / xyrange // pixels per x or y unit
	zscale = height * 0.4         // pixels per z unit

	writer.Header().Set("Content-Type", "image/svg+xml")
	surface(writer)
}

func surface(w io.Writer) {
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", int32(width), int32(height))
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			if math.IsNaN(ax) || math.IsNaN(ay) ||
				math.IsNaN(bx) || math.IsNaN(by) ||
				math.IsNaN(cx) || math.IsNaN(cy) ||
				math.IsNaN(dx) || math.IsNaN(dy) {
				continue
			}
			fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintln(w, "</svg>")
}

func corner(i, j int) (float64, float64) {
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

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}
