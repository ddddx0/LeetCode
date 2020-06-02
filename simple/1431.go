package main

import "fmt"

func main() {

	candies := []int{2, 3, 5, 1, 3}
	extraCandies := 3
	withCandies := kidsWithCandies(candies, extraCandies)
	fmt.Print(withCandies)
}

func kidsWithCandies(candies []int, extraCandies int) []bool {
	result := make([]bool, len(candies))
	var max = 0

	for _, value := range candies {
		if value > max {
			max = value
		}
	}

	for index, value := range candies {
		result[index] = value+extraCandies >= max
	}

	return result
}
