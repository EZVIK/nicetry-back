package main

import "fmt"

//有1、2、3、4个数字，可以组成多少个互不相同并且无重复数字的三位数？都是多少？

func SearchCombinationNumber(length, bit int) int {
	// length! / (length - bit)! * bit!

	if bit > length && (length < 0 || bit < 0) {
		return 0
	}

	return factorial(length) / factorial(length-bit) * factorial(bit)
}

func factorial(n int) int {
	ans := 1

	for i := 1; i <= n; i++ {
		ans *= i
	}
	return ans
}

func main() {
	//nums := []int{1, 2, 3, 4}

	fmt.Println(SearchCombinationNumber(3, 2))
	// 12，13，14, 21, 31, 41
	// 23, 24, 32, 42
	// 34, 43
}

// 40320
