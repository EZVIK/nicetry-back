package main

import (
	"fmt"
)

// [1, 2, 3, 2, 2, 2, 5, 4, 2]
// [1, -1,-1,1, 1, 1, -1,-1,1]
// Boyer-Moore 投票算法
func majorityElement(nums []int) int {
	count := 0
	var candidate int

	for _, v := range nums {

		if count == 0 {
			candidate = v
		}

		if v == candidate {
			count += 1
		} else {
			count += -1
		}
	}
	return candidate
}

func majorityElement1(nums []int) int {
	var vote, sum = 0, 0
	for _, x := range nums {
		if vote == 0 {
			sum = x
		}
		if sum == x {
			vote++
		} else {
			vote--
		}
	}
	return sum
}
func main() {

	nums := []int{1, 2, 3, 2, 2, 2, 5, 4, 2}

	fmt.Println("RESULT:", majorityElement(nums))
}
