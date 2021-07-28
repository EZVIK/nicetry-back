package main

import (
	"fmt"
	"time"
)

/**
 * @Description:

	F(0) = 0,   F(1) = 1
	F(N) = F(N - 1) + F(N - 2), 其中 N > 1.

 * @param n
 * @return int
*/
//func fib(n int) int {
//	fmt.Println(runtime.StartTrace())
//	if n == 0 {
//		return 0
//	} else if n == 1 {
//		return 1
//	} else {
//		return fib(n-1) + fib(n-2)
//	}
//}
// F0 0
// F1 1
// F2 1
// F3 2
// F4 3
// F5 5
// F6 8
/*
class Solution {
    public int fib(int n) {
        if(n == 0) return 0;
        int[] dp = new int[n + 1];
        dp[0] = 0;
        dp[1] = 1;
        for(int i = 2; i <= n; i++){
            dp[i] = dp[i-1] + dp[i-2];
            dp[i] %= 1000000007;
        }
        return dp[n];
    }
}
*/

func fib(n int) int {
	if n == 0 {
		return 0
	}

	dp := make([]int, n+1)

	dp[0] = 0
	dp[1] = 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
		dp[i] %= 1000000007
	}

	return dp[n]
}

func fib1(n int) int {

	if n == 0 {
		return 0
	}

	arr := make([]int, n+1)
	arr[0] = 0
	arr[1] = 1
	for i := 2; i <= n; i++ {
		arr[i] = arr[i-1] + arr[i-2]
		arr[i] %= 1000000007
	}

	return arr[n]
}

func main() {
	start := time.Now()
	fmt.Println(fib1(45))
	fmt.Println(time.Since(start))
}
