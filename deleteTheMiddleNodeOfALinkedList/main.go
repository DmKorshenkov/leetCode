func deleteMiddle(head *ListNode) *ListNode {
	res := head
	slow := head
	fast := head
	//	step := 0

	if head.Next == nil {
		return nil
	}
	lastElem := slow
	nextElem := slow.Next
	for fast != nil && fast.Next != nil {
		lastElem = slow
		slow = slow.Next
		nextElem = slow.Next
		fast = fast.Next.Next
		//		step++
	}

	lastElem.Next = nextElem
	//	fmt.Println(ln, step)
	return res
}
