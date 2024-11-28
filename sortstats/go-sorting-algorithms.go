package main

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"
)

func bubbleSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func insertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
	return arr
}

func selectionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		minIdx := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
	return arr
}

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[len(arr)/2]
	left := make([]int, 0)
	middle := make([]int, 0)
	right := make([]int, 0)

	for _, num := range arr {
		switch {
		case num < pivot:
			left = append(left, num)
		case num == pivot:
			middle = append(middle, num)
		case num > pivot:
			right = append(right, num)
		}
	}

	left = quickSort(left)
	right = quickSort(right)

	return append(append(left, middle...), right...)
}

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

func generateRandomSlice(n int) []int {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = rand.Intn(1000000)
	}
	return slice
}

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		slice := generateRandomSlice(1000)
		b.StartTimer()
		bubbleSort(slice)
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		slice := generateRandomSlice(1000)
		b.StartTimer()
		insertionSort(slice)
	}
}

func BenchmarkSelectionSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		slice := generateRandomSlice(1000)
		b.StartTimer()
		selectionSort(slice)
	}
}

func BenchmarkQuickSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		slice := generateRandomSlice(1000)
		b.StartTimer()
		quickSort(slice)
	}
}

func BenchmarkMergeSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		slice := generateRandomSlice(1000)
		b.StartTimer()
		mergeSort(slice)
	}
}

func BenchmarkStdSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		slice := generateRandomSlice(1000)
		b.StartTimer()
		sort.Ints(slice)
	}
}

func timeSort(sortFunc func([]int) []int, arr []int) ([]int, time.Duration) {
	start := time.Now()
	sorted := sortFunc(append([]int{}, arr...))
	duration := time.Since(start)
	return sorted, duration
}

func main() {
	arr1 := []int{64, 34, 25, 12, 22, 11, 90}
	arr2 := []int{5, 2, 8, 12, 1, 6}
	arr3 := []int{1, 2, 3, 4, 5}
	threshold:=1000

	fmt.Println("Bubble Sort:")
	result, duration := timeSort(bubbleSort, arr1)
	fmt.Printf("%v\nTime: %v\n", result, duration)
	result, duration = timeSort(bubbleSort, arr2)
	fmt.Printf("%v\nTime: %v\n", result, duration)
	result, duration = timeSort(bubbleSort, arr3)
	fmt.Printf("%v\nTime: %v\n", result, duration)

	fmt.Println("\nInsertion Sort:")
	result, duration = timeSort(insertionSort, arr1)
	fmt.Printf("%v\nTime: %v\n", result, duration)
	result, duration = timeSort(insertionSort, arr2)
	fmt.Printf("%v\nTime: %v\n", result, duration)
	result, duration = timeSort(insertionSort, arr3)
	fmt.Printf("%v\nTime: %v\n", result, duration)

	fmt.Println("\nSelection Sort:")
	result, duration = timeSort(selectionSort, arr1)
	fmt.Printf("%v\nTime: %v\n", result, duration)
	result, duration = timeSort(selectionSort, arr2)
	fmt.Printf("%v\nTime: %v\n", result, duration)
	result, duration = timeSort(selectionSort, arr3)
	fmt.Printf("%v\nTime: %v\n", result, duration)

	fmt.Println("\nQuick Sort:")
	result, duration = timeSort(quickSort, arr1)
	fmt.Printf("%v\nTime: %v\n", result, duration)
	result, duration = timeSort(quickSort, arr2)
	fmt.Printf("%v\nTime: %v\n", result, duration)
	result, duration = timeSort(quickSort, arr3)
	fmt.Printf("%v\nTime: %v\n", result, duration)

	fmt.Println("\nMerge Sort:")
	result, duration = timeSort(mergeSort, arr1)
	fmt.Printf("%v\nTime: %v\n", result, duration)
	result, duration = timeSort(mergeSort, arr2)
	fmt.Printf("%v\nTime: %v\n", result, duration)
	result, duration = timeSort(mergeSort, arr3)
	fmt.Printf("%v\nTime: %v\n", result, duration)

	// Benchmark with larger array
	fmt.Println("\nBenchmark with random integers, threshold = ",threshold)
	largeArr := make([]int, threshold)
	for i := range largeArr {
		largeArr[i] = i
	}
	for i := range largeArr {
		j := i + rand.Intn(threshold-i)
		largeArr[i], largeArr[j] = largeArr[j], largeArr[i]
	}

	_, duration = timeSort(bubbleSort, largeArr)
	fmt.Printf("Bubble Sort Time: %v\n", duration)
	_, duration = timeSort(insertionSort, largeArr)
	fmt.Printf("Insertion Sort Time: %v\n", duration)
	_, duration = timeSort(selectionSort, largeArr)
	fmt.Printf("Selection Sort Time: %v\n", duration)
	_, duration = timeSort(quickSort, largeArr)
	fmt.Printf("Quick Sort Time: %v\n", duration)
	_, duration = timeSort(mergeSort, largeArr)
	fmt.Printf("Merge Sort Time: %v\n", duration)
}
