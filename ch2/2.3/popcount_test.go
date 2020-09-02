package popcount

import "testing"

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCountExpr(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCountLoop(x uint64) int {
	sum := 0
	for i := 0; i < 8; i++ {
		sum += int(pc[byte(x>>uint(i))])
	}
	return sum
}

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

/*
go test -bench=.

goos: windows
goarch: amd64
pkg: gopl.exercise/2.3
BenchmarkExpr-8         330578056                3.45 ns/op
BenchmarkLoop-8         157481120                7.39 ns/op
PASS
ok      gopl.exercise/2.3       3.604s

*/