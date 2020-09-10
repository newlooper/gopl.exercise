package main

import (
	"fmt"
	"gopl.exercise/ch6/6.1/intset"
)

func main() {
	s1 := new(intset.IntSet)
	s1.Add(1)
	s1.Add(9)
	s1.Add(144)
	fmt.Printf("%d\n", s1.Len())

	s2 := new(intset.IntSet)
	s2.Add(79)

	s1.UnionWith(s2)
	fmt.Printf("%d\n", s1.Len())

	s1.Remove(144)
	fmt.Printf("%d\n", s1.Len())

	s3 := s1.Copy()
	s3.Add(123)
	fmt.Printf("%s\n%s", s1, s3)
}
