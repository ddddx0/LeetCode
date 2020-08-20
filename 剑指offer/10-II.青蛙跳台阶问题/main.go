package main

import "fmt"

func main() {
	fmt.Println(numWays(0))
}

// 反向动态规划
func numWays(n int) int {
	if n == 0 {
		return 1
	}
	if n == 2 || n == 1 {
		return n
	}
	dp := make([]int, n+1)
	dp[n] = 0
	dp[n-1] = 1
	dp[n-2] = 2
	for i := n; i > 2; i-- {
		dp[i-3] = (dp[i-1] + dp[i-2]) % 1000000007
	}

	return dp[0]
}
