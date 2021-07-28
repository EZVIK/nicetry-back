package main

import (
	"fmt"
	"math"
	"strconv"
)

//给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
//
//回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。例如，121 是回文，而 123 不是。

func reverse1(x int) int {
	ans := 0
	xx := int(math.Abs(float64(x)))
	l := len(strconv.Itoa(xx))
	length := int(math.Pow10(l - 1))

	for {
		if x == 0 {
			break
		}
		ans += (x % 10) * length
		x /= 10
		length /= 10
		if ans > 2147483648-1 || ans < -2147483648 {
			return 0
		}
	}
	return ans
}

func isPalindrome(x int) bool {

	if x < 0 {
		return false
	}

	xx := int(math.Abs(float64(x)))
	l := len(strconv.Itoa(xx))
	if l == 1 {
		return true
	}

	length := int(math.Pow10(l - 1))
	blength := 1
	loopTimes := 0
	i := 0
	if l%2 == 0 {
		loopTimes = l / 2
	} else {
		loopTimes = (l - 1) / 2
	}

	for {
		start := (xx / length) % 10
		end := (xx / blength) % 10

		if start == end {
			length /= 10
			blength *= 10
			i++
		} else {
			return false
		}

		if i == loopTimes {
			return true
		}
	}
}

func main() {

	// 6 100000
	a := 0

	fmt.Println(isPalindrome(a))

}
