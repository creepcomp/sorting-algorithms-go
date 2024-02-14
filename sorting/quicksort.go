package sorting

// partition partitions the array into two parts, where elements less than or equal to the pivot are on the left,
// and elements greater than the pivot are on the right.
func partition(array []int, low, high int) int {
	// Choose the rightmost element as the pivot
	pivot := array[high]

	// Index of smaller element
	i := low - 1

	// Iterate through the array
	for j := low; j < high; j++ {
		// If the current element is smaller than or equal to the pivot
		if array[j] <= pivot {
			i++
			// Swap array[i] and array[j]
			array[i], array[j] = array[j], array[i]
		}
	}

	// Swap array[i+1] and array[high] (pivot)
	array[i+1], array[high] = array[high], array[i+1]

	// Return the index of the pivot element
	return i + 1
}

// QuickSort sorts a slice of integers in ascending order using the quicksort algorithm.
func QuickSort(array []int, low, high int) {
	if low < high {
		// Partition the array, and get the index of the pivot element
		pivot := partition(array, low, high)

		// Recursively sort the subarrays
		QuickSort(array, low, pivot-1)
		QuickSort(array, pivot+1, high)
	}
}
