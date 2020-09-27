package intset

import (
	"testing"
)

func newIntSets() []IntSet {
	return []IntSet{
		&BitIntSet{},
		NewMapSet(),
	}
}

func TestInitially(t *testing.T) {
	for _, s := range newIntSets() {
		if s.Len() != 0 {
			t.Errorf("%T.Len(): got %d, want 0", s, s.Len())
		}
	}
}

func TestLenAfterAddingElements(t *testing.T) {
	for _, s := range newIntSets() {
		s.Add(0)
		s.Add(2000)
		if s.Len() != 2 {
			t.Errorf("%T.Len(): got %d, want 2", s, s.Len())
		}
	}
}

func TestRemove(t *testing.T) {
	for _, s := range newIntSets() {
		s.Add(0)
		s.Remove(0)
		if s.Has(0) {
			t.Errorf("%T: want zero removed, got %s", s, s)
		}
	}
}

func TestClear(t *testing.T) {
	for _, s := range newIntSets() {
		s.Add(0)
		s.Add(1000)
		s.Clear()
		if s.Has(0) || s.Has(1000) {
			t.Errorf("%T: want empty set, got %s", s, s)
		}
	}
}

func TestCopy(t *testing.T) {
	for _, orig := range newIntSets() {
		orig.Add(1)
		copy := orig.Copy()
		copy.Add(2)
		if !copy.Has(1) || orig.Has(2) {
			t.Errorf("%T: want %s, got %s", orig, orig, copy)
		}
	}
}

func TestAddAll(t *testing.T) {
	for _, s := range newIntSets() {
		s.AddAll(0, 2, 4)
		if !s.Has(0) || !s.Has(2) || !s.Has(4) {
			t.Errorf("%T: want {2 4}, got %s", s, s)
		}
	}
}

func TestUnionWith(t *testing.T) {
	/////////////////////////
	// BitIntSet
	b1 := new(BitIntSet)
	b1.AddAll(1, 2, 3)

	b2 := new(BitIntSet)
	b2.AddAll(1, 3, 5, 7)

	b1.UnionWith(b2)
	if b1.Len() != 5 {
		t.Errorf("%v: want 5, got %d", b1, b1.Len())
	}

	/////////////////////////
	// MapIntSet
	m1 := NewMapSet()
	m1.AddAll(1, 2, 3)

	m2 := NewMapSet()
	m2.AddAll(1, 3, 5, 7)

	m1.UnionWith(m2)
	if m1.Len() != 5 {
		t.Errorf("%v: want 5, got %d", m1, m1.Len())
	}
}
