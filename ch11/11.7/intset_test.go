package intset

import (
	"math/rand"
	"testing"
)

func newBitIntSet() *BitIntSet {
	return &BitIntSet{}
}

const max = 10000

func benchAdd(b *testing.B, set IntSet, n int) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < n; j++ {
			set.Add(rand.Intn(max))
		}
		set.Clear()
	}
}

func benchAddAll(b *testing.B, set IntSet, batchSize int) {
	ints := randomInts(batchSize)
	for i := 0; i < b.N; i++ {
		set.AddAll(ints...)
		set.Clear()
	}
}

func benchUnionWith(bm *testing.B, a, b IntSet, n int) {
	randomAdd(a, n)
	randomAdd(b, n)
	for i := 0; i < bm.N; i++ {
		a.UnionWith(b)
	}
}

func randomAdd(set IntSet, n int) {
	for i := 0; i < n; i++ {
		set.Add(rand.Intn(max))
	}
}

func randomInts(n int) []int {
	ints := make([]int, n)
	for i := 0; i < n; i++ {
		ints[i] = rand.Intn(max)
	}
	return ints
}

///////////////////////////////
// BitIntSet.Add
func BenchmarkBitIntSetAdd10(b *testing.B) {
	benchAdd(b, newBitIntSet(), 10)
}

func BenchmarkBitIntSetAdd100(b *testing.B) {
	benchAdd(b, newBitIntSet(), 100)
}

func BenchmarkBitIntSetAdd1000(b *testing.B) {
	benchAdd(b, newBitIntSet(), 1000)
}

///////////////////////////////
// MapIntSet.Add
func BenchmarkMapIntSetAdd10(b *testing.B) {
	benchAdd(b, NewMapSet(), 10)
}

func BenchmarkMapIntSetAdd100(b *testing.B) {
	benchAdd(b, NewMapSet(), 100)
}

func BenchmarkMapIntSetAdd1000(b *testing.B) {
	benchAdd(b, NewMapSet(), 1000)
}

///////////////////////////////
// BitIntSet.AddAll
func BenchmarkBitIntSetAddAll10(b *testing.B) {
	benchAddAll(b, newBitIntSet(), 10)
}

func BenchmarkBitIntSetAddAll100(b *testing.B) {
	benchAddAll(b, newBitIntSet(), 100)
}

func BenchmarkBitIntSetAddAll1000(b *testing.B) {
	benchAddAll(b, newBitIntSet(), 1000)
}

///////////////////////////////
// MapIntSet.AddAll
func BenchmarkMapIntSetAddAll10(b *testing.B) {
	benchAddAll(b, NewMapSet(), 10)
}

func BenchmarkMapIntSetAddAll100(b *testing.B) {
	benchAddAll(b, NewMapSet(), 100)
}

func BenchmarkMapIntSetAddAll1000(b *testing.B) {
	benchAddAll(b, NewMapSet(), 1000)
}

///////////////////////////////
// BitIntSet.UnionWith
func BenchmarkBitIntSetUnionWith10(b *testing.B) {
	benchUnionWith(b, newBitIntSet(), newBitIntSet(), 10)
}

func BenchmarkBitIntSetUnionWith100(b *testing.B) {
	benchUnionWith(b, newBitIntSet(), newBitIntSet(), 100)
}

func BenchmarkBitIntSetUnionWith1000(b *testing.B) {
	benchUnionWith(b, newBitIntSet(), newBitIntSet(), 1000)
}

///////////////////////////////
// MapIntSet.UnionWith
func BenchmarkMapIntSetUnionWith10(b *testing.B) {
	benchUnionWith(b, NewMapSet(), NewMapSet(), 10)
}

func BenchmarkMapIntSetUnionWith100(b *testing.B) {
	benchUnionWith(b, NewMapSet(), NewMapSet(), 100)
}

func BenchmarkMapIntSetUnionWith1000(b *testing.B) {
	benchUnionWith(b, NewMapSet(), NewMapSet(), 1000)
}

/*
goos: linux
goarch: amd64
pkg: gopl.exercise/ch11/11.7
BenchmarkBitIntSetAdd10-8                 889567              1234 ns/op
BenchmarkBitIntSetAdd100-8                372495              3102 ns/op
BenchmarkBitIntSetAdd1000-8                58004             20786 ns/op
BenchmarkMapIntSetAdd10-8                1883686               630 ns/op
BenchmarkMapIntSetAdd100-8                167761              7103 ns/op
BenchmarkMapIntSetAdd1000-8                14985             79730 ns/op
BenchmarkBitIntSetAddAll10-8             1000000              1037 ns/op
BenchmarkBitIntSetAddAll100-8             802938              1401 ns/op
BenchmarkBitIntSetAddAll1000-8            268288              4416 ns/op
BenchmarkMapIntSetAddAll10-8             2747330               443 ns/op
BenchmarkMapIntSetAddAll100-8             223131              5243 ns/op
BenchmarkMapIntSetAddAll1000-8             19538             61244 ns/op
BenchmarkBitIntSetUnionWith10-8          2414660               456 ns/op
BenchmarkBitIntSetUnionWith100-8         2416165               502 ns/op
BenchmarkBitIntSetUnionWith1000-8        2441613               503 ns/op
BenchmarkMapIntSetUnionWith10-8          2074341               549 ns/op
BenchmarkMapIntSetUnionWith100-8          141489              8760 ns/op
BenchmarkMapIntSetUnionWith1000-8          10000            109465 ns/op
PASS
ok      gopl.exercise/ch11/11.7 26.157s
*/
