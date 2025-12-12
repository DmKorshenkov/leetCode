func removeNthFromEnd(head *ListNode, n int) {
	arr := make([]int, 0, 30)
	newHead := &ListNode{}
	curr := newHead
	//	tmp := newHead
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	fmt.Println(arr)
	for i := range arr {
		if i == len(arr)-n {
			//fmt.Println("\u2714", arr[i])
			continue
		}
	//	fmt.Println(arr[i])
		newHead.Next = &ListNode{Val: arr[i]}
		newHead = newHead.Next
	}

	return curr.Next
}
