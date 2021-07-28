package main

import (
	"math"
)

/**
	把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。
    输入一个递增排序的数组的一个旋转，输出旋转数组的最小元素。
    例如，数组 [3,4,5,1,2] 为 [1,2,3,4,5] 的一个旋转，该数组的最小值为1。
			   5,1,2,3,4
*/

// 37912
// [3,4,5,1,2] 为 [1,2,3,4,5]
func minArray(numbers []int) int {
	min := math.MaxInt64

	// spin index = x
	first, last := numbers[0], numbers[len(numbers)-1]
	if first > last {

		for i := len(numbers) - 1; i >= 0; i-- {
			if numbers[i] < min {
				min = numbers[i]
			}
		}

		return min

	} else {

		for i := 0; i >= 0; i-- {
			if numbers[i] < min {
				min = numbers[i]
			}
		}

		return min
	}
}

func main() {

}
