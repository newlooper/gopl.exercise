package charcount

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func charCount(r io.Reader) (counts map[rune]int, cats map[string]int, utflen map[int]int, invalid int) {
	counts = make(map[rune]int) // counts of Unicode characters
	cats = make(map[string]int)
	utflen = make(map[int]int) // count of lengths of UTF-8 encodings

	in := bufio.NewReader(r)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		for catName, rangeTable := range unicode.Properties {
			if unicode.In(r, rangeTable) {
				cats[catName]++
			}
		}
		counts[r]++
		utflen[n]++
	}
	return
}
