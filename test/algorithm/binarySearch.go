package main

import (
	"fmt"
	"math"
)

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

	for {
		res := nums[mid] - target

		if res > 0 {
			mid = (start + mid) / 2
		} else if res < 0 {
			mid = (mid + ll) / 2
		} else {
			return mid
		}
	}

}

// return -1
func BinarySearch2(nums []int, target int) int {

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

func main() {

	a := []int{1, 2, 3, 4}

	fmt.Println(BinarySearch(a, 4))
}
