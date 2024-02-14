package sorting

// findMinIndexAndValue finds the minimum value and its index within a specified range of an array.
func findMinIndexAndValue(array []int, start, end int) (int, int) {
	minIndex := start
	min := array[start]
	for i := start; i <= end; i++ {
		if array[i] <= min {
			min = array[i]
			minIndex = i
		}
	}
	return minIndex, min
}

// SelectionSort sorts a slice of integers in ascending order using the selection sort algorithm.
func SelectionSort(array []int) {
	length := len(array) - 1
	for i := 0; i < length; i++ {
		minIndex, _ := findMinIndexAndValue(array, i+1, length)
		if array[minIndex] < array[i] {
			array[i], array[minIndex] = array[minIndex], array[i]
		}
	}
}
