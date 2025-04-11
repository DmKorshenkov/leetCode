package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (h *ListNode) set(val int) {

	for h.Next != nil {
		h = h.Next
	}
	h.Next = &ListNode{Val: val, Next: nil}
}

func main() {

	ls := &ListNode{Val: 1}

	ls.set(3)
	ls.set(1)
	ls.set(4)

	mid, newHead := getMid(ls)
	right := mid.Next
	mid = nil

	for newHead != nil {
		fmt.Println(newHead)
		newHead = newHead.Next
	}
	fmt.Println()
	for right != nil {
		fmt.Println(right)
		right = right.Next
	}
}

func getMid(head *ListNode) (*ListNode, *ListNode) {
	var tmp *ListNode
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		tmp = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow, tmp
}
