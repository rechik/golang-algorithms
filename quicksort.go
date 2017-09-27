package golang_algorithms

func QuickSort(array []int) []int {
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

		sorted := QuickSort(less)
		sorted = append(sorted, pivot)
		sorted = append(sorted, QuickSort(greater)...)

		return sorted
	}
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