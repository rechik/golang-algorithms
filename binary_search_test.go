package golang_algorithms

import (
	"sort"
	"testing"
)

var data = []int{0: -10, 1: -5, 2: 0, 3: 1, 4: 2, 5: 3, 6: 5, 7: 7, 8: 11, 9: 100, 10: 110, 11: 150, 12: 1000, 13: 10000}

var tests = []struct {
	name string
	n    int
	i    int
}{
	{"-10 0", -10, 0},
	{"-5 1", -5, 1},
	{"0 2", 0, 2},
	{"1 3", 1, 3},
	{"100 9", 100, 9},
	{"666 -1", 666, -1},
}

func TestSearch(t *testing.T) {
	for _, e := range tests {
		i := BinarySearchInts(data, e.n)
		if i != e.i {
			t.Errorf("%s: expected index %d; got %d", e.name, e.i, i)
		}
	}
}

func BenchmarkBinarySearchInts(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BinarySearchInts(data, 666)
	}
}

func BenchmarkStandartSearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sort.SearchInts(data, 666)
	}
}
