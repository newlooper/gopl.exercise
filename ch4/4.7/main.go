package main

import (
	"fmt"
	"unicode/utf8"
)

var turns = 0

func main() {
	s := []byte("一a 二b 三c")
	fmt.Printf("%d\t|% x|\t%q\n", turns, s, s)
	reverseUTF8(s)
	fmt.Printf("%d\t|% x|\t%q\n\n", turns, s, s)
}

func reverseUTF8(s []byte) {
	for i := 0; i < len(s); {
		_, size := utf8.DecodeRune(s[i:])
		if size > 1 { // 非 ascii 字符占用大于一个字节
			reverse(s[i : i+size]) // UTF-8 编码局部翻转
			turns++
			fmt.Printf("%d\t|% x|\t%q\n", turns, s, s)
		}
		i += size
	}
	reverse(s) // 最后整体翻转
	turns++
}

func reverse(s []byte) {
	tail := len(s) - 1
	for i := 0; i < len(s)/2; i++ {
		s[i], s[tail-i] = s[tail-i], s[i]
	}
}

/*
go run .

0       |e4 b8 80 61 20 e4 ba 8c 62 20 e4 b8 89 63|     "一a 二b 三c"
1       |80 b8 e4 61 20 e4 ba 8c 62 20 e4 b8 89 63|     "\x80\xb8\xe4a 二b 三c"
2       |80 b8 e4 61 20 8c ba e4 62 20 e4 b8 89 63|     "\x80\xb8\xe4a \x8c\xba\xe4b 三c"
3       |80 b8 e4 61 20 8c ba e4 62 20 89 b8 e4 63|     "\x80\xb8\xe4a \x8c\xba\xe4b \x89\xb8\xe4c"
4       |63 e4 b8 89 20 62 e4 ba 8c 20 61 e4 b8 80|     "c三 b二 a一"

*/
