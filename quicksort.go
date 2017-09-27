package golang_algorithms

import (
	//"math/rand"
)

func QuickSortBlunt(array []int) []int {
	if len(array) < 2 {
		return array
	} else {
		pivot := array[0]

		var less, greater []int

		for _, v := range array[1:] {
			if (v <= pivot) {
				less = append(less, v)
			} else if (v > pivot) {
				greater = append(greater, v)
			}
		}

		sorted := QuickSortBlunt(less)
		sorted = append(sorted, pivot)
		sorted = append(sorted, QuickSortBlunt(greater)...)

		return sorted
	}
}

func QuickSortOptimize(array []int) {
	if len(array) < 2 {
		return
	}

	// Set boundaries of used elements
	left, right := 0, len(array) - 1

	// Used for divide arrays values on two arrays
	// Can by any index between [left, right]
	pivot := 0

	// Move pivot value to the end
	array[pivot], array[right] = array[right], array[pivot]

	//  Move value that less than the pivot value to the beginning
	// and move value that greater than pivot value to the end
	for k := range array {
		if array[k] < array[right] {
			array[k], array[left] = array[left], array[k]
			left++
		}
	}

	// Move pivot value between left and right arrays,
	// after the last smaller element
	array[left], array[right] = array[right], array[left]

	// Repeate actions fo left and right arrays
	QuickSortOptimize(array[:left])
	QuickSortOptimize(array[left + 1:])
}

func IsSorted(array []int) bool {
	n := len(array)
	for i := n - 1; i > 0; i-- {
		if i < i-1 {
			return false
		}
	}
	return true
}