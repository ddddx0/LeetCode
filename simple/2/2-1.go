package main

import "fmt"

func main() {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}

	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	numbers := addTwoNumbers1(l1, l2)
	fmt.Print(numbers)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	node1 := l1
	node2 := l2
	r := new(ListNode)
	res := r
	carry := 0
	for node1 != nil || node2 != nil || carry != 0 {
		r.Next = new(ListNode)
		r = r.Next

		a := 0
		b := 0
		if node1 != nil {
			a = node1.Val
		}
		if node2 != nil {
			b = node2.Val
		}

		r.Val = (a + b + carry) % 10
		carry = (a + b + carry) / 10

		if node1 != nil {
			node1 = node1.Next
		}
		if node2 != nil {
			node2 = node2.Next
		}

	}
	return res.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
