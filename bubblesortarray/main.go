package main

import (
	"fmt"
	"strings"
)

func main(){
	arr1 := []int{7,2,3,12,5,6,11,8,9}
	fmt.Println("bubble sorted array of numbers in ascending := ",BubbleSortNumbers(arr1))
	arr2 := []string{"banana","apple","strawberry"}
	fmt.Println("bubble sorted array of alphabets in ascending := ",BubbleSortAlphabets(arr2))
}

func BubbleSortNumbers(arr []int) []int {
	length := len(arr)
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j],arr[j+1] = arr[j+1],arr[j]
			}
		}
	}
	return arr
}


func BubbleSortAlphabets(arr []string) []string{
 
	length := len(arr)

	for i:=0;i<length-1;i++{
		for j:=0;j<length-i-1;j++{ //0 if a == b, -1 if a < b, and +1 if a > b.
			if strings.Compare(arr[j],arr[j+1]) > 0{
                arr[j],arr[j+1]=arr[j+1],arr[j]
			}
		}
	}

	return arr
}