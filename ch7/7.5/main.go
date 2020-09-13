package main

import (
	"fmt"
	"io"
	"os"
)

// 典型的装饰器 Decorator
type LimitedReader struct {
	r io.Reader // has-a io.Reader
	n int
}

// is-a io.Reader
func (l *LimitedReader) Read(p []byte) (n int, err error) {
	if l.n <= 0 {
		return 0, io.EOF
	}

	if len(p) > l.n {
		p = p[:l.n]
	}
	n, err = l.r.Read(p)
	l.n -= n
	return
}

// 工厂
func SimpleLimitedReader(r io.Reader, n int) io.Reader {
	return &LimitedReader{r, n}
}

func main() {
	var b = make([]byte, 1024)

	lr := SimpleLimitedReader(os.Stdin, 5)
	n, err := lr.Read(b)

	fmt.Printf("%v\t%v", n, err)
}

/*
go run .

abc123()_
5       <nil>

 */
