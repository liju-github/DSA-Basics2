package main

import "fmt"



func MergeSort(arr []int)[]int  {
	
	if len(arr) <= 1{
		return arr
	}

	mid := len(arr)/2
	left := MergeSort(arr[:mid])
	right :=  MergeSort(arr[mid:])

	return merge(left,right)
	
}


func merge(left , right []int) []int  {
	result := []int{}
	i,j := 0,0

	for i<len(left) && j<len(right){
		if left[i]<=right[j]{
			result = append(result, left[i])
		    i++
		}else{
			result = append(result, right[j])
		    j++
		}
	}

	for i<len(left){
		result = append(result, left[i])
		i++
	}

	for j<len(right){
		result = append(result, right[j])
		j++
	}

	return result
}


func main()  {
	arr := []int{5,1,8,3,8,2,10,22,0}

	fmt.Println("unsorted array : ",arr)
	fmt.Println("sorted array : ",MergeSort(arr))
}