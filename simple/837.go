package main

import (
	"fmt"
)

func main() {
	game := new21Game(10, 1, 10)
	fmt.Println(len(game))
	fmt.Print(game)
}

func new21Game(N int, K int, W int) []float64 {
	rate := make([]float64, K+W)

	for i := K; i <= N && i <= K+W; i++ {
		rate[i] = 1.0
	}

	for i := K - 1; 0 <= i; i-- {
		for j := 1; j <= W; j++ {
			rate[i] += rate[i+j] / float64(W)
		}
	}
	return rate
}

func new21Game1(N int, K int, W int) float64 {

	dp := make([]float64, K+W)
	for i := K; i <= N && i < K+W; i++ {
		dp[i] = 1.0
	}
	for i := K - 1; i >= 0; i-- {
		for j := 1; j <= W; j++ {
			dp[i] += dp[i+j] / float64(W)
		}
	}
	return dp[0]
}
