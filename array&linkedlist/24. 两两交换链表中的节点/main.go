package main

import "fmt"

func main() {
	list := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}

	fmt.Println(swapPairs(list))
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	// 新建一个链，用以持有之前的链
	res := &ListNode{0, head}
	pre := res

	for pre.Next != nil && pre.Next.Next != nil {
		//first,second 表示需要调换的第一个和第二个
		first, second := pre.Next, pre.Next.Next

		// 调换位置
		pre.Next = second
		pre.Next.Next = second.Next
		second.Next = pre.Next

		// 向后改变pre位置
		pre = first
	}
	return res.Next
}
