package golang_algorithms

// Search and return index of given item
// If no result then return -1
func binary_search_ints(list []int, item int) int {
	var low = 0
	var high = len(list) - 1
	var mid int
	var guess int

	for low <= high {
		mid = int(uint(low+high) >> 1)

		guess = list[mid]

		if guess == item {
			return mid
		}

		if guess > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}
