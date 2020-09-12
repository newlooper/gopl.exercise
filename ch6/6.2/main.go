package main

import (
	"fmt"
	"gopl.exercise/ch6/6.2/intset"
)

func main() {
	s := new(intset.IntSet)
	s.AddAll(1, 2, 3)
	fmt.Printf("%d\n", s.Len())
}
