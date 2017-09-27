package golang_algorithms

import (
	"testing"
	"sort"
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

func BenchmarkQuickSortInts(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QuickSort(unsortedInts)
	}
}

func BenchmarkStandartSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sort.Ints(unsortedInts)
	}
}