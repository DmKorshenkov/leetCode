func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}


	mid := getMid(head)
	right := mid.Next
	mid.Next = nil
//	fmt.Println("2")

	leftSorted := sortList(head)
	rightSorted := sortList(right)


	return merge(leftSorted, rightSorted)
}
func getMid(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func merge(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			curr.Next = l1
			l1 = l1.Next
		} else {
			curr.Next = l2
			l2 = l2.Next
		}
		curr = curr.Next
	}

	if l1 != nil {
		curr.Next = l1
	} else {
		curr.Next = l2
	}

	return dummy.Next
}
