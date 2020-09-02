package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	start := time.Now()

	for i := 0; i < 1000000; i++ {
		strings.Join(os.Args[1:], " ")
	}

	fmt.Println(time.Now().Sub(start))

	start = time.Now()

	for i := 0; i < 1000000; i++ {
		for s, sep, i := "", "", 0; i < len(os.Args); i++ {
			s += sep + os.Args[i]
			sep = " "
		}
	}

	fmt.Println(time.Now().Sub(start))
}
