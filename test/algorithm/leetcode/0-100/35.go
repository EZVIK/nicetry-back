package main

import (
	"fmt"
)

/**
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

你可以假设数组中无重复元素。

示例 1:

输入: [1,3,5,6], 5
输出: 2
示例2:

输入: [1,3,5,6], 2
输出: 1
示例 3:

输入: [1,3,5,6], 7
输出: 4
示例 4:

输入: [1,3,5,6], 0
输出: 0

*/

func searchInsert(nums []int, target int) int {

	i, _ := 0, len(nums)-1
	for {
		if i == len(nums) {
			return i
		}
		//if j == 0 {
		//	return j
		//}

		if target > nums[i] {
			i++
		} else {
			return i
		}
		//
		//if target < nums[j] {
		//	j--
		//} else {
		//	return j
		//}
	}
}

//
func searchInsert2(nums []int, target int) int {
	n := len(nums)
	left, right := 0, n-1
	ans := n
	for left <= right {
		mid := (right-left)>>1 + left
		if target <= nums[mid] {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ans
}

func main() {

	a := []int{1, 3, 5, 7}

	target := 17

	fmt.Println(searchInsert(a, target))

}
