package main

import (
	"fmt"
)

func quickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	pivotIndex := partition(arr)
	
	quickSort(arr[:pivotIndex])
	quickSort(arr[pivotIndex+1:])
}

func partition(arr []int) int {
	pivot := arr[len(arr)-1]
	i := -1

	for j := 0; j < len(arr)-1; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1],arr[len(arr)-1] = arr[len(arr)-1],arr[i+1]

	return i + 1
}

func main() {
	arr := []int{64, 25, 12, 22, 11, 90}
	fmt.Println("unsorted array:", arr)

	quickSort(arr)

	fmt.Println("sorted array:", arr)
}
