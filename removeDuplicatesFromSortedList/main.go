package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (list *ListNode) Append(val int) {
	head := list
	for head.Next != nil {
		head = head.Next
	}
	head.Next = &ListNode{Val: val}
	list = head
}

func (list *ListNode) print() {
	head := list

	for head != nil {
		fmt.Println(head)
		head = head.Next
	}
}

func main() {
	//	list := New()
	list := new(ListNode)
	list.Append(1)
	list.Append(4)
	list.Append(4)
	deleteDuplicates(list).print()

}

func deleteDuplicates(head *ListNode) *ListNode {
	//	current := head
	ptr := head

	for head != nil && head.Next != nil {
		if head.Next.Val == head.Val {
			head.Next = head.Next.Next
			continue
		}
		head = head.Next
	}
	return ptr
}
