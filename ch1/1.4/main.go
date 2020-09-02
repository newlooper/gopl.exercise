package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for filename, lines := range counts {
		for line, n := range lines {
			if n > 1 {
				fmt.Printf("%d\t%s found in %s\n", n, line, filename)
			}
		}
	}
}

func countLines(f *os.File, counts map[string]map[string]int) {
	input := bufio.NewScanner(f)
	filename := f.Name()
	for input.Scan() {
		if counts[filename] == nil {
			counts[filename] = make(map[string]int)
		}
		counts[filename][input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}
