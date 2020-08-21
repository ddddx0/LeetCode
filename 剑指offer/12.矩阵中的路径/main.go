package main

import "fmt"

// 未完成
func main() {
	//board := [][]byte{{"A","B","C","E"},}
	//fmt.Println(exist())
}

func exist(board [][]byte, word string) bool {
	src := []byte(word)

	var row, column int
	for i, r := range board {
		for j, c := range r {
			if c == src[0] {
				row = i
				column = j
			}
		}
	}
	fmt.Println(row)
	fmt.Println(column)

	return false
}
