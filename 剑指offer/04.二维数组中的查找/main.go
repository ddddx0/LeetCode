package main

import "fmt"

func main() {
	//matrix := [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}
	matrix := [][]int{{1, 1}}
	//matrix1 := [][]int{{-5}}
	target := 1
	//fmt.Println(findNumberIn2DArray(matrix1, target))
	fmt.Println(findNumberIn2DArray1(matrix, target))
}

// 自己想的方法，m[x][y]
// target小于当前数字则直接查询下一个数组，即x+1
func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[len(matrix)-1]) == 0 || target < matrix[0][0] || target > matrix[len(matrix)-1][len(matrix[len(matrix)-1])-1] {
		return false
	}
	for _, v := range matrix {
		for _, val := range v {
			if target == val {
				return true
			} else if target < val {
				continue
			}
		}
	}
	return false
}

// 标准答案，m[x][y]
// 从左上或者右下开始
// target小于当前数字，则x+1
// target大于当前数字，则y+1
func findNumberIn2DArray1(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[len(matrix)-1]) == 0 || target < matrix[0][0] || target > matrix[len(matrix)-1][len(matrix[len(matrix)-1])-1] {
		return false
	}
	x, y := len(matrix)-1, 0

	for x >= 0 && y <= len(matrix[0])-1 {
		val := matrix[x][y]
		if target == val {
			return true
		} else if target < val {
			x--
		} else if target > val {
			y++
		}
	}

	return false
}
