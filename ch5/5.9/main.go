package main

import "strings"

var s = "The quick brown $fox jumps over the lazy $dog"

func main() {
	println(expand(s, f))
}

func expand(s string, f func(string) string) string {
	words := strings.Split(s, " ")
	for i, placeholder := range words {
		if strings.HasPrefix(placeholder, "$") {
			words[i] = f(placeholder[1:])
		}
	}
	return strings.Join(words, " ")
}

func f(p string) string {
	return p
}
