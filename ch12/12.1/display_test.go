package display

import "testing"

func TestMapKeys(t *testing.T) {
	sm := map[struct{ x int }]int{
		{1}: 2,
		{2}: 3,
	}
	Display("sm", sm)
	// Output:
	// Display sm (map[struct { x int }]int):
	// sm[{x: 2}] = 3
	// sm[{x: 1}] = 2

	am := map[[3]int]int{
		{1, 2, 3}: 3,
		{2, 3, 4}: 4,
	}
	Display("am", am)
	// Output:
	// Display am (map[[3]int]int):
	// am[{1, 2, 3}] = 3
	// am[{2, 3, 4}] = 4
}
