package main

import "fmt"

func main() {
	s := []string{"1", "2", "2", "3", "3", "3", "4", "4", "4", "4", "5", "5", "5", "5", "5"}
	fmt.Println(dropDup(s))
}

func dropDup(s []string) []string {
	if len(s) == 0 {
		return s
	}

	i := 0
	for j := 1; j < len(s); j++ {
		if s[i] != s[j] {
			i++
			s[i] = s[j]
		}
	}
	return s[:i+1]
}
