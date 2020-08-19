package main

import "fmt"

func main() {
	list := &ListNode{
		Val: 10,
		Next: &ListNode{
			Val: 11,
			Next: &ListNode{
				Val: 12,
				Next: &ListNode{
					Val:  13,
					Next: nil,
				},
			},
		},
	}
	fmt.Println(reversePrint(list))
}

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 顺序取出，反转数组
func reversePrint(head *ListNode) []int {
	res := []int{}
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}

	for i, l := 0, len(res); i < (len(res) / 2); {
		i, res[i], res[l-i-1] = i+1, res[l-i-1], res[i]
	}

	return res
}
