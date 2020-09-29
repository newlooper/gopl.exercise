package popcount

import "testing"

//////////////////////////////////////////////////
// horizontal comparison
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

//////////////////////////////////////////////////
// vertical comparison

// benchDiffTimes
func benchDiffTimes(b *testing.B, n int, f func(uint64) int) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < n; j++ {
			f(uint64(j))
		}
	}
}

func BenchmarkPopCountExpr10(b *testing.B) {
	benchDiffTimes(b, 10, PopCountExpr)
}

func BenchmarkPopCountExpr100(b *testing.B) {
	benchDiffTimes(b, 100, PopCountExpr)
}

func BenchmarkPopCountExpr1000(b *testing.B) {
	benchDiffTimes(b, 1000, PopCountExpr)
}

func BenchmarkPopCountLoop10(b *testing.B) {
	benchDiffTimes(b, 10, PopCountLoop)
}

func BenchmarkPopCountLoop100(b *testing.B) {
	benchDiffTimes(b, 100, PopCountLoop)
}

func BenchmarkPopCountLoop1000(b *testing.B) {
	benchDiffTimes(b, 1000, PopCountLoop)
}
func BenchmarkPopCountShift10(b *testing.B) {
	benchDiffTimes(b, 10, PopCountShift)
}

func BenchmarkPopCountShift100(b *testing.B) {
	benchDiffTimes(b, 100, PopCountShift)
}

func BenchmarkPopCountShift1000(b *testing.B) {
	benchDiffTimes(b, 1000, PopCountShift)
}

func BenchmarkPopCountBitwiseAnd10(b *testing.B) {
	benchDiffTimes(b, 10, PopCountBitwiseAnd)
}

func BenchmarkPopCountBitwiseAnd100(b *testing.B) {
	benchDiffTimes(b, 100, PopCountBitwiseAnd)
}

func BenchmarkPopCountBitwiseAnd1000(b *testing.B) {
	benchDiffTimes(b, 1000, PopCountBitwiseAnd)
}

/*
go test -bench=.
goos: linux
goarch: amd64
pkg: gopl.exercise/ch11/11.6
BenchmarkExpr-8                         397002367                2.96 ns/op
BenchmarkLoop-8                         93851917                12.7 ns/op
BenchmarkShift-8                        34220011                34.2 ns/op
BenchmarkBitwiseAnd-8                   100000000               10.1 ns/op
BenchmarkPopCountExpr10-8               33610244                35.2 ns/op
BenchmarkPopCountExpr100-8               3631207               330 ns/op
BenchmarkPopCountExpr1000-8               358345              3352 ns/op
BenchmarkPopCountLoop10-8               11749201               103 ns/op
BenchmarkPopCountLoop100-8               1245673               962 ns/op
BenchmarkPopCountLoop1000-8               122538              9633 ns/op
BenchmarkPopCountShift10-8               3431937               353 ns/op
BenchmarkPopCountShift100-8               353660              3416 ns/op
BenchmarkPopCountShift1000-8               33517             34133 ns/op
BenchmarkPopCountBitwiseAnd10-8         49750624                24.1 ns/op
BenchmarkPopCountBitwiseAnd100-8         2323900               489 ns/op
BenchmarkPopCountBitwiseAnd1000-8         174968              5752 ns/op
PASS
ok      gopl.exercise/ch11/11.6 21.961s
*/
