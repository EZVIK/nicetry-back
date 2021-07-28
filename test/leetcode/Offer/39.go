package main

import (
	"fmt"
)

// [1, 2, 3, 2, 2, 2, 5, 4, 2]
// [1, -1,-1,1, 1, 1, -1,-1,1]
// Boyer-Moore 投票算法
func majorityElement(nums []int) int {
	count := 0
	candidate := -1

	for _, v := range nums {

		if count == 0 {
			candidate = v
		}
		if v == candidate {
			count += 1
		} else {
			count += -1
		}
		fmt.Println(count)
	}
	return candidate
}

func main() {

	nums := []int{1, 2, 3, 2, 2, 2, 5, 4, 2}

	fmt.Println("RESULT:", majorityElement(nums))
}
