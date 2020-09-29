package intset

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func popCountExpr(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// IntersectWith return intersection of s and t.
func (s *BitIntSet) IntersectWith(t *BitIntSet) *BitIntSet {
	u := new(BitIntSet)
	for _, i := range t.Elems() {
		if s.Has(i) {
			u.Add(i)
		}
	}
	return u
}

// DifferenceWIth return difference of s and t.
func (s *BitIntSet) DifferenceWith(t *BitIntSet) *BitIntSet {
	d := new(BitIntSet)
	for _, i := range s.Elems() {
		if !t.Has(i) {
			d.Add(i)
		}
	}
	return d
}

// SymmetricDifference return symmetric difference of s and t.
func (s *BitIntSet) SymmetricDifference(t *BitIntSet) *BitIntSet {
	m := s.copy()
	m.UnionWith(t)
	return m.DifferenceWith(s.IntersectWith(t))
}

func (s *BitIntSet) copy() *BitIntSet {
	clone := new(BitIntSet)
	clone.words = make([]uint, len(s.words))
	copy(clone.words, s.words)
	clone.length = s.length
	return clone
}
