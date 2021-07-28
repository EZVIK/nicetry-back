package main

import (
	"fmt"
	"math"
	"strconv"
)

// 120 => 21

// 123 => 321
//-231 <= x <= 231 - 1
func reverse(x int) int {

	max := int(math.Pow(2, 31))
	if x > max-1 || x < -max {
		return 0
	}

	ans := int(math.Abs(float64(x)))
	l := len(strconv.Itoa(ans))
	length := int(math.Pow10(l - 1))

	if x < 0 {
		ans = -rrr(ans, length)
	} else {
		ans = rrr(ans, length)
	}

	if ans > max-1 || ans < -max {
		return 0
	} else {
		return ans
	}

}

func rr(x, times int) int {

	l := len(strconv.Itoa(x))

	if l == 1 {
		return x * times
	} else {
		xc := x % int(math.Pow10(l-1))

		x1 := x / int(math.Pow10(l-1)) * times

		p := rr(xc, times*10)

		return p + x1
	}
}

// 1001
func rrr(x, times int) int {
	l := len(strconv.Itoa(x))
	if l == 1 {
		return x * 1
	} else {
		last := (x % 10) * times
		xc := x / 10
		return last + rrr(xc, times/10)
	}
}

func r(x int) int {
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

// -231 <= x <= 231 - 1
func main() {
	a := -10200
	fmt.Println(r(a))

}
