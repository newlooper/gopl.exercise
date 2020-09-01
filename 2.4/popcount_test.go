package popcount

import "testing"

func benchmark(b *testing.B, f func(uint64) int) {
	for i := 0; i < b.N; i++ {
		f(uint64(i))
	}
}

func BenchmarkExpr(b *testing.B) {
	benchmark(b, PopCountExpr)
}

func BenchmarkLoop(b *testing.B) {
	benchmark(b, PopCountLoop)
}

func BenchmarkShift(b *testing.B) {
	benchmark(b, PopCountShift)
}


/*
go test -bench=.

goos: windows
goarch: amd64
pkg: gopl.exercise/2.4
BenchmarkExpr-8         331490704                3.57 ns/op
BenchmarkLoop-8         85721020                14.2 ns/op
BenchmarkShift-8        36364516                37.0 ns/op
PASS
ok      gopl.exercise/2.4       6.403s

*/
