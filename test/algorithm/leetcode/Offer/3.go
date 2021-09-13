package main

import "fmt"

/**

找出数组中重复的数字。


在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。

数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。

示例 1：

输入：
[2, 3, 1, 0, 2, 5, 3]
输出：2 或 3

*/

func findRepeatNumber(nums []int) int {

	n := make([]int, len(nums))

	for _, k := range nums {
		n[k]++
		if n[k] > 1 {
			return k
		}
	}
	return 0
}

func main() {
	//{0, 0, 0, 0, 0, 0, 0}
	arr := []int{4, 3, 1, 0, 2, 5, 3}

	fmt.Println(findRepeatNumber(arr))
}
