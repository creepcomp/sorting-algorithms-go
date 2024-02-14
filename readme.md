# Sorting Algorithms in Go

This repository contains implementations of popular sorting algorithms in Go. These sorting algorithms are commonly used to efficiently rearrange elements in a collection in ascending order.

## Implemented Algorithms

1. **Insertion Sort**: This algorithm iterates over the array, selecting each element and shifting it to its correct position in the sorted subarray.
   
2. **Selection Sort**: This algorithm divides the array into a sorted and an unsorted region. It repeatedly selects the smallest (or largest) element from the unsorted region and moves it to the end of the sorted region.

3. **Quicksort**: This algorithm divides the array into two partitions around a pivot element. Elements smaller than the pivot are placed to its left, and elements greater than the pivot are placed to its right. The partitions are then recursively sorted.

## Usage

To use these sorting algorithms in your Go projects, import the `sorting` package and call the corresponding functions:

```go
package main

import (
	"fmt"
	"github.com/yourmodulepath/sorting-algorithms-go/sorting"
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
