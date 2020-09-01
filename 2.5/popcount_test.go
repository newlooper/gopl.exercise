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

func BenchmarkBitwiseAnd(b *testing.B) {
	benchmark(b, PopCountBitwiseAnd)
}

/*
go test -bench=.

goos: windows
goarch: amd64
pkg: gopl.exercise/2.5
BenchmarkExpr-8                 317461408                3.61 ns/op
BenchmarkLoop-8                 85717346                14.8 ns/op
BenchmarkShift-8                33334351                37.3 ns/op
BenchmarkBitwiseAnd-8           136208728                9.22 ns/op
PASS
ok      gopl.exercise/2.5       6.399s

*/
