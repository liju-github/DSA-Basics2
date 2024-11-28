package main

import (
	"fmt"
	"strings"
)

func main() {
	arr1 := []int{7, 2, 3, 12, 5, 6, 11, 8, 9}
	fmt.Println("bubble sorted array of numbers in ascending := ", SelectionSortNumbers(arr1))
	arr2 := []string{"banana","apple","strawberry"}
	fmt.Println("bubble sorted array of alphabets in ascending := ",SelectionSortAlphabets(arr2))
}

func SelectionSortNumbers(arr []int) []int {
	length := len(arr)

	for i := 0; i < length-1; i++ {
		minIndex := i
		for j := i+1; j < length; j++ {
			if arr[j] > arr[minIndex] {
				minIndex = j
			}
		}

		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}

	return arr

}

func SelectionSortAlphabets(arr []string)[]string {
	length := len(arr)

	for i := 0; i < length-1; i++ {
		minIndex := i
		for j := i+1; j < length; j++ {
			if strings.Compare(arr[minIndex],arr[j]) > 0  {
				minIndex = j
			}
		}

		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}

	return arr
}
