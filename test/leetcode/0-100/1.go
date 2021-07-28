package main

import (
	"fmt"
)

//定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target 的那两个整数，并返回它们的数组下标。
//
//你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
//
//你可以按任意顺序返回答案。
/*

	输入：nums = [2,7,11,15], target = 9

	输出：[0,1]
	解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。

*/

//func twoSum(nums []int, target int) []int {
//
//	for i:=0;i< len(nums);i++{
//		for j:=i+1;j<len(nums);j++{
//			if nums[i] + nums[j] == target {
//				return []int{i, j}
//			}
//		}
//	}
//
//	return []int{}
//}

func twoSum(nums []int, target int) []int {

	resMap := make(map[int]int, len(nums))

	for i, k := range nums {

		if resMap[k] != 0 {
			return []int{resMap[k] - 1, i}
		}

		gap1 := target - k

		resMap[gap1] = i + 1
	}

	return []int{}
}

func main() {
	arr := []int{3, 3, 4, 5}
	//arr := []int{1,3,4,2}
	target := 8

	fmt.Println(twoSum(arr, target))
}
