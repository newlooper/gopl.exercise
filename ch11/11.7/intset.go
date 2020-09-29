package intset

////////////////////////////////////
// IntSet 的实现者应该拥有的方法
type IntSet interface {
	Has(x int) bool
	Add(x int)
	AddAll(nums ...int)
	UnionWith(t IntSet)
	Len() int
	Remove(x int)
	Clear()
	Copy() IntSet
	String() string
	Elems() []int
}
