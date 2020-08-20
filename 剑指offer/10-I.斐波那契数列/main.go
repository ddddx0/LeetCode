package main

import "fmt"

func main() {
	fmt.Println(fib(5))
}

// 动态规划
func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	//dp := make([]int, n+1)
	//dp[0] = 0
	//dp[1] = 1

	f1, f2, res := 0, 1, 0

	for i := 2; i <= n; i++ {
		//dp[i] = (dp[i-1] + dp[i-2]) % 1000000007
		res = (f1 + f2) % 1000000007

		f1 = f2

		f2 = res
	}
	return res
	//return dp[n]
}
