package main

func main() {

}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var result *ListNode

	if l1 == nil && l2 == nil {
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

	i := sum(l1.Val, l2.Val)

	result = &ListNode{
		Val:  i,
		Next: addTwoNumbers(l1.Next, l2.Next),
	}

	return result
}

func twoNumbers(l1 *ListNode, l2 *ListNode, minus int) *ListNode {
	var result *ListNode

	if l1 == nil && l2 == nil {
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

	i := sum(l1.Val, l2.Val)
	i = i + minus

	if i > 10 {

	}
	result = &ListNode{
		Val:  i % 10,
		Next: twoNumbers(l1.Next, l2.Next, 1),
	}

	return result
}

func sum(a, b int) int {
	return a + b
}

type ListNode struct {
	Val  int
	Next *ListNode
}
