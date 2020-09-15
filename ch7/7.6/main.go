package main

import (
	"flag"
	"fmt"
	"gopl.exercise/ch7/7.6/tempconv"
)

func main() {
	var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature")
	flag.Parse()
	fmt.Println(*temp)
}

/*
go run . -temp 100F
37.77777777777778°C

go run . -temp 100C
100°C

go run . -temp 100K
-173.14999999999998°C
 */
