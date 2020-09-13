package main

import (
	"fmt"
	"io"
)

// 实现 io.Reader 接口
type StringReader struct {
	s   string
	pos int
}

// 满足此接的代码行为详见：https://golang.org/pkg/io/#Reader
func (r *StringReader) Read(p []byte) (n int, err error) {
	if len(p) == 0 {
		return 0, nil
	}

	n = copy(p, r.s[r.pos:])
	if r.pos += n; r.pos >= len(r.s) {
		err = io.EOF
	}
	return
}

func SimpleNewReader(s string) *StringReader {
	return &StringReader{s, 0}
}

func main() {
	var b = make([]byte, 1024)

	sr := SimpleNewReader("hello")
	n, err := sr.Read(b)

	fmt.Printf("%v\t%v", n, err)
}
