package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	fmt.Printf("%s", squashUnicodeSpace([]byte("乱花 渐欲  迷人眼,　　浅草 才能   没马蹄")))
}

func squashUnicodeSpace(s []byte) []byte {
	result := s[:0:0]
	var prev rune

	for i := 0; i < len(s); {
		r, size := utf8.DecodeRune(s[i:])
		if !unicode.IsSpace(r) {
			result = append(result, s[i:i+size]...)
		} else if unicode.IsSpace(r) && !unicode.IsSpace(prev) {
			result = append(result, ' ')
		}
		prev = r
		i += size
	}
	return result
}
