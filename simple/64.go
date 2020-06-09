package main

import "fmt"

func main() {
	fmt.Print(sumNums(9))
}
func sumNums(n int) int {
	if n == 1 {
		return 1
	}
	nums := sumNums(n - 1)
	return n + nums
}
