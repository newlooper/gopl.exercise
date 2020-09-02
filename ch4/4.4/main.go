package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(rotateLeft(s, 7))
}

func rotateLeft(s []int, n int) []int {
	n = n % len(s)
	if n > 0 {
		temp := append(s[:0:0], s[0:n]...)
		copy(s, s[n:])
		copy(s[len(s)-n:], temp)
	}
	return s
}
