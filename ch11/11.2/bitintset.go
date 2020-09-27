package intset

import (
	"bytes"
	"fmt"
)

const bits = 32 << (^uint(0) >> 63)

// An BitIntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type BitIntSet struct { // IntSet 接口的实现者之一：位向量实现
	words  []uint
	length int // individual member for performance
}

// recover bits to it original value
func (s *BitIntSet) Elems() []int {
	if s.Len() == 0 {
		return nil
	}
	ints := make([]int, 0, s.Len())
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < bits; j++ {
			if word&(1<<uint(j)) != 0 {
				ints = append(ints, bits*i+j)
			}
		}
	}
	return ints
}

// batch add
func (s *BitIntSet) AddAll(x ...int) {
	for _, i := range x {
		s.Add(i)
	}
}

// return the number of elements
func (s *BitIntSet) Len() int {
	return s.length
}

// remove x from the set
func (s *BitIntSet) Remove(x int) {
	if s.Has(x) {
		word, bit := x/bits, uint(x%bits)
		s.words[word] &^= 1 << bit
		s.length--
	}
}

// remove all elements from the set
func (s *BitIntSet) Clear() {
	s.words = nil
	s.length = 0
}

// return a copy of the set
func (s *BitIntSet) Copy() IntSet {
	clone := new(BitIntSet)
	clone.words = make([]uint, len(s.words))
	copy(clone.words, s.words)
	clone.length = s.length
	return clone
}

// Has reports whether the set contains the non-negative value x.
func (s *BitIntSet) Has(x int) bool {
	word, bit := x/bits, uint(x%bits)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *BitIntSet) Add(x int) {
	word, bit := x/bits, uint(x%bits)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
	s.length++
}

// UnionWith sets s to the union of s and t.
func (s *BitIntSet) UnionWith(t IntSet) {
	if t, ok := t.(*BitIntSet); ok { // 断言，放大成员集
		for i, tword := range t.words {
			if i < len(s.words) {
				s.words[i] |= tword
			} else {
				s.words = append(s.words, tword)
			}
		}
	} else {
		for _, i := range t.Elems() {
			s.Add(i)
		}
	}
	s.length = 0
	for _, word := range s.words {
		s.length += popCountExpr(uint64(word))
	}
}

// String returns the set as a string of the form "{1 2 3}".
func (s *BitIntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < bits; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", bits*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
