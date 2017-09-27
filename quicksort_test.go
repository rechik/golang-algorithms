package golang_algorithms

import (
	"testing"
)

var unsortedInts = []int{
	0: 6,
	1: 2,
	2: 12,
	3: 1,
	4: 12,
	5: 11,
	6: 4,
	7: 0,
	8: 5,
	9: 1,
	10: 7,
	11: 9,
	12: 10,
	13: 0,
}

func TestQuickSort(t *testing.T) {
	sorted := QuickSort(unsortedInts)
	if len(unsortedInts) != len(sorted) && !IsSorted(sorted) {
		t.Errorf("Got %d", sorted)
	}
}