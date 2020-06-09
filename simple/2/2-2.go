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

	numbers := addTwoNumbers(l1, l2)
	fmt.Print(numbers)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return addTwoNumbersAndCarry(l1, l2, 0)
}

func addTwoNumbersAndCarry(l1, l2 *ListNode, carry int) *ListNode {
	res := &ListNode{
		Val:  0,
		Next: nil,
	}

	if l1 == nil && l2 == nil && carry == 0 {
		return nil
	}
	if l1 == nil {
		l1 = &ListNode{
			Val:  0,
			Next: nil,
		}
	}
	if l2 == nil {
		l2 = &ListNode{
			Val:  0,
			Next: nil,
		}
	}
	value := l1.Val + l2.Val + carry
	res.Val = value % 10
	if value >= 10 {
		carry = 1
	} else {
		carry = 0
	}

	next := addTwoNumbersAndCarry(l1.Next, l2.Next, carry)

	res.Next = next
	return res
}
