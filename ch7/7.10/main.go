package main

import (
	"fmt"
	"sort"
)

func IsPalindrome(s sort.Interface) bool {
	tail := s.Len() - 1
	for i := 0; i < s.Len()/2; i++ {
		if s.Less(i, tail-i) || s.Less(tail-i, i) {
			return false
		}
	}
	return true
}

func main() {
	s := []int{1, 2, 3, 3, 2, 1}
	fmt.Printf("%v\t", s)
	fmt.Println(IsPalindrome(sort.IntSlice(s)))

	s = s[0:2]
	fmt.Printf("%v\t", s)
	fmt.Println(IsPalindrome(sort.IntSlice(s)))
}
