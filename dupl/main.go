	head := &TreeNode{}
	head.Next = l1
	arr := make([]int, 0, 100)
	for l1 != nil || l2 != nil {
		sum := 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		arr = append(arr, sum)
	}
//	fmt.Println(arr)
	for i := range arr {
		if arr[i] >= 10 {
			arr[i] = arr[i] % 10
			if i == len(arr)-1 {
				arr = append(arr, 0)
			}
			arr[i+1]++
		}
	}
//	fmt.Println(len(arr), arr)
	l3 := head.Next
	for i := len(arr) - 0; i < len(arr); i++ {
		//		fmt.Println(arr[i])
		l3.Val = arr[i]
		if i == 0 {
			return head.Next
		}
		if l3.Next == nil {
			l3.Next = &TreeNode{Val: 0}
		}
		l3 = l3.Next
		//	fmt.Println(l3)
	}
	//	l3.pr()
	//	fmt.Println(l3)
	return head.Next
