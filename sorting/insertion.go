package sorting

// InsertionSort sorts a slice of integers in ascending order using the insertion sort algorithm.
func InsertionSort(array []int) {
	n := len(array)
	if n <= 1 {
		// Nothing to sort if array has 0 or 1 element
		return
	}

	for i := 1; i < n; i++ {
		key := array[i]
		j := i - 1
		for j >= 0 && array[j] > key {
			array[j+1] = array[j]
			j--
		}
		array[j+1] = key
	}
}
