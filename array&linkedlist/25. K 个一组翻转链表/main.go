package main

func main() {

}

//* Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	// pre, tail
	start := &ListNode{
		Val:  0,
		Next: head,
	}
	pre := start

	for head != nil {
		tail := pre
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return start.Next
			}
		}

		nex := tail.Next
		head, tail = reverseList(head, tail)
		pre.Next = head
		tail.Next = nex
		pre = tail
		head = tail.Next

	}

	return start.Next
}

// 调换指定片段的链表
func reverseList(head, tail *ListNode) (*ListNode, *ListNode) {
	prev := tail.Next
	p := head
	for prev != tail {
		nex := p.Next
		p.Next = prev
		prev = p
		p = nex
	}
	return tail, head
}
