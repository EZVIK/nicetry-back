package main

import (
	"fmt"
)

/**
在一个 n * m 的二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。请完成一个高效的函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。


[
  [1,   4,  7, 11, 15],
  [2,   5,  8, 12, 19],
  [3,   6,  9, 16, 22],
  [10, 13, 14, 17, 24],
  [18, 21, 23, 26, 30]
]

*/

func findNumberIn2DArray1(matrix [][]int, target int) bool {

	if len(matrix) == 0 {
		return false
	}

	for _, arr := range matrix {

		l := len(arr) - 1

		if len(arr) == 0 {
			return false
		}

		if target < arr[0] || target > arr[l] {
			continue
		}

		front := arr[0] - target
		tail := arr[l] - target

		if front == 0 || tail == 0 {
			return true
		}

		for _, subArr := range arr {
			if subArr == target {
				return true
			}
		}

	}

	return false

}

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	if len(matrix[0]) == 0 {
		return false
	} else if len(matrix[0]) == 1 && matrix[0][0] == target {
		return true
	}

	deep := len(matrix) - 1
	length := len(matrix[0])

	for i, j := deep, 0; i >= 0 && j < length; {

		if matrix[i][j] > target {
			i--
		} else if matrix[i][j] < target {
			j++
		} else {
			return true
		}
	}
	return false
}

func main() {
	// [[1,4,7,11,15],
	// [2,5,8,12,19],
	// [3,6,9,16,22],
	// [10,13,14,17,24],
	// [18,21,23,26,30]]

	arr := make([][]int, 1)
	arr[0] = []int{3, 5}

	//arr[0] = []int{1,   4,  7, 11, 15}
	//arr[1] = []int{2,   5,  8, 12, 19}
	//arr[2] = []int{3,   6,  9, 16, 22}
	//arr[3] = []int{10, 13, 14, 17, 24}
	//arr[4] = []int{18, 21, 23, 26, 30}

	//arr[0] = []int{1,2,3,4,5}
	//arr[1] = []int{6,7,8,9,10}
	//arr[2] = []int{11,12,13,14,15}
	//arr[3] = []int{16,17,18,19,20}
	//arr[4] = []int{21,22,23,24,25}

	fmt.Println(findNumberIn2DArray(arr, 3))
}
