package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("%s\n", comma(os.Args[i], 3))
	}
}

func comma(s string, segLen int) string {

	var b bytes.Buffer

	integerStart := 0

	if s[0] == '+' || s[0] == '-' {
		b.WriteByte(s[0])
		integerStart = 1
	}

	number := s[integerStart:]

	arrNumber := strings.Split(number, ".")
	if len(arrNumber) == 2 {
		number = arrNumber[0]
	}

	first := len(number) % segLen
	if first == 0 {
		first = segLen
	}
	b.WriteString(number[:first])
	for i := first; i < len(number); i += segLen {
		b.WriteByte(',')
		b.WriteString(number[i : i+segLen])
	}

	if len(arrNumber) == 2 {
		b.WriteByte('.')
		b.WriteString(arrNumber[1])
	}
	return b.String()
}
