package treesort

import (
	"sort"
	"testing"
)

func TestSort(t *testing.T) {
	tests := []([]int){
		[]int{5, 4, 3, 2, 1},
		[]int{8989, 65, 1029},
	}
	for _, test := range tests {
		result := make([]int, len(test))
		want := make([]int, len(test))
		copy(want, test)
		copy(result, test)
		Sort(result)
		sort.Sort(sort.IntSlice(want))
		if !equal(want, result) {
			t.Errorf("Sort(%v) = %v, want %v", test, result, want)
		}
	}
}

func equal(x, y []int) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}
