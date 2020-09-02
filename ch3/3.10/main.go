package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("%s\n", comma(os.Args[i], 3))
	}
}

func comma(s string, segLen int) string {
	var b bytes.Buffer
	first := len(s) % segLen
	if first == 0 {
		first = segLen
	}
	b.WriteString(s[:first])
	for i := first; i < len(s); i += segLen {
		b.WriteByte(',')
		b.WriteString(s[i : i+segLen])
	}
	return b.String()
}
