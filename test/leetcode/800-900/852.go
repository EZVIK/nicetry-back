package main

import (
	"fmt"
	"sort"
)

func peakIndexInMountainArray(arr []int) int {
	for i := 1; ; i++ {
		if arr[i] > arr[i+1] {
			return i
		}
	}
}

/*
					5
				4 		3
			3 				2
		2 						1
	1
0


					5
				4 		3
			 3				2
		 						1


0  1  2  3  4  5  6  7  8
*/

// binary Search
func binarySearchPeek(arr []int) int {
	start := 0
	end := len(arr)
	mid := (start + end) / 2

	for {
		// left
		if arr[mid] > arr[mid+1] {
			end = mid
			mid = (start + end) / 2
		}

		// right
		if arr[mid] < arr[mid+1] {
			start = mid
			mid = (start + end) / 2
		}

		// peek
		if arr[mid] > arr[mid+1] && arr[mid] > arr[mid-1] {
			return mid
		}
	}
}

func peakIndexInMountainArray2(arr []int) int {
	return sort.Search(len(arr)-1, func(i int) bool { return arr[i] > arr[i+1] })
}

func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 3, 2, 1}
	//arr := []int{24,69,99,79,78,100,67,36,26,19}

	fmt.Println(binarySearchPeek(arr))
}
