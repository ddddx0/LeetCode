package main

import "container/list"

func main() {

}

type CQueue struct {
	in, out *list.List
}

func Constructor() CQueue {
	return CQueue{
		in:  list.New(),
		out: list.New(),
	}
}

func (this *CQueue) AppendTail(value int) {
	this.in.PushFront(value)
}

func (this *CQueue) DeleteHead() int {
	if this.out.Len() == 0 {
		for this.in.Len() > 0 {
			this.out.PushFront(this.in.Remove(this.in.Front()))
		}
	}

	if this.out.Len() != 0 {
		return this.out.Remove(this.out.Front()).(int)
	}

	return -1
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
