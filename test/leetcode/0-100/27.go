package main

import "fmt"

/**
示例 1：

输入：nums = [3,2,2,3], val = 3
输出：2, nums = [2,2]
解释：函数应该返回新的长度 2, 并且 nums 中的前两个元素均为 2。你不需要考虑数组中超出新长度后面的元素。例如，函数返回的新长度为 2 ，而 nums = [2,2,3,3] 或 nums = [2,2,0,0]，也会被视作正确答案。
示例 2：

输入：nums = [0,1,2,2,3,0,4,2], val = 2
输出：5, nums = [0,1,4,0,3]
解释：函数应该返回新的长度 5, 并且 nums 中的前五个元素为 0, 1, 3, 0, 4。注意这五个元素可为任意顺序。你不需要考虑数组中超出新长度后面的元素。


提示：

0 <= nums.length <= 100
0 <= nums[i] <= 50
0 <= val <= 100

*/
func removeElement(nums []int, val int) int {
	var ll int

	for _, v := range nums {
		if val != v {
			nums[ll] = v
			ll += 1
		}
	}
	return ll
}

func removeDuplicates27(nums []int) int {
	var ll int
	var num = -10001
	for _, v := range nums {
		if num != v {
			num = v
			nums[ll] = v
			ll += 1
		}
	}
	return ll
}

func main() {

	nums := []int{3, 3, 4, 3, 3, 3, 3, 3, 3, 2, 2, 3}
	// 		[)

	k3 := removeElement(nums, 3)

	fmt.Println(k3)
}
