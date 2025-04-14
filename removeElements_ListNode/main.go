package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{}
	newHead := dummy

	for head != nil {
		if head.Val != val {
			fmt.Println(head.Val)
			dummy.Next = head
			dummy = dummy.Next
		}
		head = head.Next
	}

	dummy.Next = nil
	return newHead.Next
}

func removeElementsInPlace(head *ListNode, val int) *ListNode {

	for head != nil && head.Val == val {
		head = head.Next
	}

	current := head

	for current != nil && current.Next != nil {
		if current.Next.Val == val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}

	return head
}
