package main

import (
	"fmt"

	"github.com/creepcomp/sorting-algorithms-go/sorting"
)

func main() {
	arr := []int{5, 2, 4, 6, 1, 3}
	fmt.Println("Original array:", arr)

	// Example usage of insertion sort
	insertionArr := make([]int, len(arr))
	copy(insertionArr, arr)
	sorting.InsertionSort(insertionArr)
	fmt.Println("Insertion Sort Result:", insertionArr)

	// Example usage of selection sort
	selectionArr := make([]int, len(arr))
	copy(selectionArr, arr)
	sorting.SelectionSort(selectionArr)
	fmt.Println("Selection Sort Result:", selectionArr)

	// Example usage of quicksort
	quicksortArr := make([]int, len(arr))
	copy(quicksortArr, arr)
	sorting.QuickSort(quicksortArr, 0, len(quicksortArr)-1)
	fmt.Println("Quicksort Result:", quicksortArr)
}
