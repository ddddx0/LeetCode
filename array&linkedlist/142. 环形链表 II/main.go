package main

func main() {

}

//* Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 使用slow，fast 2个标记，slow每次向后移动一位，fast每次向后移动2位，若slow和fast重合则说明有环
func detectCycle(head *ListNode) *ListNode {
	var slow, fast = head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return slow
		}
	}
	return nil
}
