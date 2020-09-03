package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int) // counts of Unicode characters
	input := bufio.NewScanner(os.Stdin)

	input.Split(bufio.ScanWords)
	for input.Scan() {
		counts[input.Text()]++
	}

	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v", err)
		os.Exit(1)
	}

	for k, v := range counts {
		fmt.Printf("%-30s%d\n", k, v)
	}
}
