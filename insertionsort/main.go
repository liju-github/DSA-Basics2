package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(InsertionSortNumbers([]int{8, 2, 5, 1, 9, 10, 2}))
	fmt.Println(InsertionSortAlphabets([]string{"orange", "banana", "apple", "strawberry"}))

}

func InsertionSortNumbers(arr []int) []int {

	for i := 1; i < len(arr); i++ {
		temp := arr[i]

		j := i - 1
		for j > (-1) && arr[j] > temp {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = temp
	}

	return arr
}

func InsertionSortAlphabets(arr []string) []string {
	for i := 1; i < len(arr); i++ {
		temp := arr[i]
		j := i - 1

		for j > (-1) && strings.Compare(arr[j], temp) >= 1 {
			arr[j+1] = arr[j]
			j -= 1
		}

		arr[j+1] = temp
	}
	return arr
}
