package intset

import (
	"bytes"
	"fmt"
)

const bits = 64

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words  []uint64
	length int // individual member for performance
}

// IntersectWith return intersection of s and t.
func (s *IntSet) IntersectWith(t *IntSet) *IntSet {
	u := new(IntSet)
	for _, i := range t.Elems() {
		if s.Has(i) {
			u.Add(i)
		}
	}
	return u
}

// DifferenceWIth return difference of s and t.
func (s *IntSet) DifferenceWith(t *IntSet) *IntSet {
	d := new(IntSet)
	for _, i := range s.Elems() {
		if !t.Has(i) {
			d.Add(i)
		}
	}
	return d
}

// SymmetricDifference return symmetric difference of s and t.
func (s *IntSet) SymmtericDifference(t *IntSet) *IntSet {
	m := s.Copy()
	m.UnionWith(t)
	return m.DifferenceWith(s.IntersectWith(t))
}

// recover bits to it original value
func (s *IntSet) Elems() []int {
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
func (s *IntSet) AddAll(x ...int) {
	for _, i := range x {
		s.Add(i)
	}
}

// return the number of elements
func (s *IntSet) Len() int {
	return s.length
}

// remove x from the set
func (s *IntSet) Remove(x int) {
	if s.Has(x) {
		word, bit := x/bits, uint(x%bits)
		s.words[word] &^= 1 << bit
		s.length--
	}
}

// remove all elements from the set
func (s *IntSet) Clear() {
	s.words = nil
	s.length = 0
}

// return a copy of the set
func (s *IntSet) Copy() *IntSet {
	clone := new(IntSet)
	clone.words = make([]uint64, len(s.words))
	copy(clone.words, s.words)
	clone.length = s.length
	return clone
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/bits, uint(x%bits)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/bits, uint(x%bits)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
	s.length++
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
	s.length += t.length
}

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
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
