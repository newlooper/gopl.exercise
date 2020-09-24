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
