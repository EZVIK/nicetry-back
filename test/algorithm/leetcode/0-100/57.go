package main

import (
	"fmt"
	"math"
)

func twoSum1(nums []int, target int) []int {
	l, r := 0, len(nums)-1

	for {
		if r-l == 1 {
			return []int{nums[l], nums[r]}
		}
		sum := nums[l] + nums[r]
		if sum > target {
			r--
		} else if sum < target {
			l++
		} else {
			return []int{nums[l], nums[r]}
		}
	}
}

// find target's index in a order array
func BinarySearch(nums []int, target int) int {

	ll := len(nums)

	if ll < 3 {
		for i, k := range nums {
			if k == target {
				return i
			}
		}
		return -1
	}

	start, mid := 0, 0

	if ll%2 == 0 {
		mid = ll/2 - 1
	} else {
		mid = ll / 2
	}

	max := int(math.Log2(float64(ll)))
	times := 0

	for {

		res := nums[mid] - target

		if res > 0 {
			mid = (start + mid) / 2
		} else if res < 0 {
			mid = (mid + ll) / 2
		} else {
			return mid
		}

		if times == max {
			return -1
		}
		times++
	}
}

//
func twoSum3(nums []int, target int) []int {

	l, r := 0, len(nums)-1

	for {
		sum := nums[l] + nums[r]
		diff1 := target - nums[l] // big
		diff2 := target - nums[r] // semall
		if sum > target {

			if ans := BinarySearch(nums, diff1); ans != -1 {
				return []int{nums[l], nums[ans]}
			}

			r--
		} else if sum < target {

			if ans := BinarySearch(nums, diff2); ans != -1 {
				return []int{nums[r], nums[ans]}
			}

			l++
		} else {
			return []int{nums[l], nums[r]}
		}
	}
}

func twoSum22(nums []int, target int) []int {
	dist := make([]bool, nums[len(nums)-1]+1)

	for _, v := range nums {
		find := target - v
		if find >= nums[0] && find <= nums[len(nums)-1] && dist[find] {
			return []int{find, v}
		}
		dist[v] = true
	}
	return []int{}
}

func main() {

	//a := []int{10,26,30,31,47,60,70, 89, 92, 99, 103, 105, 130, 143, 154, 169}
	a := []int{1, 2}

	fmt.Println(twoSum(a, 3))
}
