package main

import "fmt"

func main() {
	num := []int{2, 3, 1, 0, 2, 5, 3}
	fmt.Println(findRepeatNumber(num))
	fmt.Println(findRepeatNumber1(num))
}

// 暴力轮询
func findRepeatNumber(nums []int) int {
	m := map[int]bool{}
	for _, idx := range nums {
		if m[idx] == true {
			return idx
		}
		m[idx] = true
	}
	return -1
}

// 数数交换
func findRepeatNumber1(nums []int) int {
	for idx, value := range nums {
		if value == idx {
			continue
		} else if nums[value] == value {
			return value
		} else {
			nums[value] = value
		}
	}
	return -1
}
