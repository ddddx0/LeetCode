package main

func main() {

}

//* Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 1,使用slow，fast 2个标记，slow每次向后移动一位，fast每次向后移动2位，若slow和fast重合则说明有环,并记录重合位置ptr
// 2,从以head和ptr为起点往后位移，相遇位置即为环的起始点（数学公式计算得出）
func detectCycle(head *ListNode) *ListNode {
	var slow, fast = head, head
	var ptr *ListNode = nil

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			ptr = slow
			break
		}
	}

	for ptr != nil {
		if head == ptr {
			return head
		}
		head, ptr = head.Next, ptr.Next
	}

	return ptr
}
