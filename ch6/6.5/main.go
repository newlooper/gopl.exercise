package main

import (
	"fmt"
	"gopl.exercise/ch6/6.5/intset"
)

func main() {
	s1 := new(intset.IntSet)
	s1.AddAll(1, 2, 3)

	s2 := new(intset.IntSet)
	s2.AddAll(1, 3, 5, 7)

	s3 := s1.IntersectWith(s2)
	fmt.Printf("%s\t%dnums\n", s3, s3.Len())

	s4 := s1.DifferenceWith(s2)
	fmt.Printf("%s\t%dnums\n", s4, s4.Len())

	s5 := s1.SymmetricDifference(s2)
	fmt.Printf("%s\t%dnums\n", s5, s5.Len())

	s1.UnionWith(s2)
	fmt.Printf("%s\t%dnums\n", s1, s1.Len())
}
