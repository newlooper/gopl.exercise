package main

import (
	"fmt"
	"gopl.exercise/ch7/7.1/words"
)

func main() {
	var w words.WordCounter
	var l words.LineCounter

	test := `a line
a word`

	_, _ = w.Write([]byte(test))
	fmt.Printf("words: %d\n", w)

	_, _ = l.Write([]byte(test))
	fmt.Printf("lines: %d\n", l)
}
