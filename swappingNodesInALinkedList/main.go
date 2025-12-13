func swapNodes(head *ListNode, k int) *ListNode {
  	curr := &ListNode{}
	newHead := curr
	arr := make([]*ListNode, 0, 100000)
    n := k
	for head != nil {
		arr = append(arr, head)
		head = head.Next
	}
	if len(arr) == 1 {
		return arr[0]
	}
	//!!!!!!
	arr[n-1], arr[len(arr)-n] = arr[len(arr)-n], arr[n-1]
	//!!!!!!!!!!!!!!!!!!!! MU MENYAEM MASSIV LISTOV GOOVNA
	arr[n-1].Next, arr[len(arr)-n].Next = arr[len(arr)-n].Next, arr[n-1].Next
	for i := range arr {
		curr.Next = &ListNode{Val: arr[i].Val}
		curr = curr.Next
	}

	return newHead.Next  
}
