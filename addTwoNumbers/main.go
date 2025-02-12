package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	// fmt.Println(addTwoNumbers(&ListNode{Val: 98}, &ListNode{Val: 1}))
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	n := int(0)
	list := &ListNode{}
	ptr := list
	for l1 != nil || l2 != nil {
		if l1 != nil {
			n += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n += l2.Val
			l2 = l2.Next
		}

		list.Next = &ListNode{Val: n % 10}
		list = list.Next
		n = n / 10
	}
	if n == 1 {
		list.Next = &ListNode{Val: n}
	}

	return ptr.Next
}
