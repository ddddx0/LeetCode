package main

import "fmt"

func main() {
	ii := []int{2, 2, 2, 0, 1}
	fmt.Println(minArray(ii))
}

// 二分法查询最小旋转数组
func minArray(numbers []int) int {
	left, right := 0, len(numbers)-1

	if right == 0 {
		return numbers[0]
	}

	for left < right {
		mid := (left + right) / 2
		if numbers[mid] > numbers[right] {
			// 当中间值大于最右值时，说明左侧区间为顺序区间且mid不可能为最小值，left右移到mid+1（即不需要包括mid）
			left = mid + 1
		} else if numbers[mid] < numbers[right] {
			// 当中间值小于最右值，说明右侧为顺序区间且mid有可能为最小值，则right左移到mid处
			right = mid
		} else if numbers[mid] == numbers[right] {
			// 在有重复数的情况下，2值相等并不能判定为找到最终结果，需左移right值后观察情况，
			// 若左移后出现不相等情况，则走上面2种情况，
			// 若左移后相等，则继续，直到最后left和right重合，退出循环，则说明全数组都是重复数
			right--
		}
	}
	// 跳出循环后，则left和right指向的同一位置，left和right都能代表最小值
	return numbers[left]
}
