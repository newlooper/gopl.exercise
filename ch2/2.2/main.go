// +build ignore

package main

import (
	"bufio"
	"fmt"
	convlib "gopl.exercise/ch2/2.2"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) > 1 {
		doConvert(os.Args[1:])
	} else {
		scan := bufio.NewScanner(os.Stdin)
		for scan.Scan() {
			doConvert(parseInput(scan.Text()))
		}
	}
}

func parseInput(s string) []string {
	reg := regexp.MustCompile(`\s+`)
	array := reg.Split(strings.TrimSpace(s), -1)
	return array
}

func doConvert(args []string) {
	for _, arg := range args {
		v, err := strconv.ParseFloat(arg, 63)
		if err != nil {
			fmt.Fprintf(os.Stderr, "convert: %v\n", err)
			os.Exit(0)
		}
		f := convlib.Fahrenheit(v)
		c := convlib.Celsius(v)
		ft := convlib.Feet(v)
		m := convlib.Meter(v)
		fmt.Printf("%s = %s, %s = %s, %s = %s, %s = %s\n",
			f, convlib.FToC(f),
			c, convlib.CToF(c),
			ft, convlib.FToM(ft),
			m, convlib.MToF(m))
	}
}
