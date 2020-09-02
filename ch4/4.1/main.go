package main

import (
	"crypto/sha256"
	"fmt"
)

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func main() {
	s1 := sha256.Sum256([]byte("x"))
	s2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%d", bitsDiff(&s1, &s2))
}

func bitsDiff(b1, b2 *[sha256.Size]byte) int {
	var diffBitsCount int
	for i := 0; i < sha256.Size; i++ {
		diffBitsCount += int(pc[b1[i]^b2[i]])
	}
	return diffBitsCount
}
